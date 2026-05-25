package app

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type MaintenanceFileStatus struct {
	Exists    bool   `json:"exists"`
	Path      string `json:"path"`
	Size      int64  `json:"size"`
	UpdatedAt string `json:"updatedAt,omitempty"`
}

type MaintenanceDirStatus struct {
	Exists    bool   `json:"exists"`
	Path      string `json:"path"`
	FileCount int    `json:"fileCount"`
	TotalSize int64  `json:"totalSize"`
}

type MaintenanceBackupStatus struct {
	Path       string   `json:"path"`
	Count      int      `json:"count"`
	KeepCount  int64    `json:"keepCount"`
	Newest     string   `json:"newest,omitempty"`
	NewestPath string   `json:"newestPath,omitempty"`
	Dirs       []string `json:"dirs"`
}

type MaintenanceGitStatus struct {
	Available bool     `json:"available"`
	Branch    string   `json:"branch"`
	Changes   []string `json:"changes"`
	Error     string   `json:"error,omitempty"`
}

type MaintenanceStatus struct {
	ContentSnapshot   MaintenanceFileStatus   `json:"contentSnapshot"`
	WeeklyCheckpoint  MaintenanceBackupStatus `json:"weeklyCheckpoints"`
	MonthlyCheckpoint MaintenanceBackupStatus `json:"monthlyCheckpoints"`
	Uploads           MaintenanceDirStatus    `json:"uploads"`
	Backups           MaintenanceBackupStatus `json:"backups"`
	Git               MaintenanceGitStatus    `json:"git"`
}

type MaintenanceCommandResult struct {
	OK         bool   `json:"ok"`
	Action     string `json:"action"`
	Output     string `json:"output"`
	StartedAt  string `json:"startedAt"`
	FinishedAt string `json:"finishedAt"`
	DurationMs int64  `json:"durationMs"`
}

func (s *Server) maintenanceStatus(c *gin.Context) {
	root, err := s.projectRoot()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	status := MaintenanceStatus{
		ContentSnapshot:   fileStatus(filepath.Join(root, "storage", "content", "content.json")),
		WeeklyCheckpoint:  checkpointStatus(filepath.Join(root, "storage", "content", "checkpoints", "weekly")),
		MonthlyCheckpoint: checkpointStatus(filepath.Join(root, "storage", "content", "checkpoints", "monthly")),
		Uploads:           dirStatus(filepath.Join(root, "storage", "uploads")),
		Backups:           backupStatus(filepath.Join(root, "storage", "backups")),
		Git:               gitStatus(root),
	}
	c.JSON(http.StatusOK, status)
}

func checkpointStatus(path string) MaintenanceBackupStatus {
	status := MaintenanceBackupStatus{
		Path:      path,
		KeepCount: 0,
		Dirs:      []string{},
	}
	entries, err := os.ReadDir(path)
	if err != nil {
		return status
	}
	for _, entry := range entries {
		if entry.IsDir() || filepath.Ext(entry.Name()) != ".json" {
			continue
		}
		status.Dirs = append(status.Dirs, entry.Name())
	}
	sort.Strings(status.Dirs)
	status.Count = len(status.Dirs)
	if status.Count > 0 {
		status.Newest = status.Dirs[status.Count-1]
		status.NewestPath = filepath.Join(path, status.Newest)
	}
	return status
}

func (s *Server) runBackup(c *gin.Context) {
	s.runContentSnapshotBackup(c)
}

func (s *Server) runGitSync(c *gin.Context) {
	root, err := s.projectRoot()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if _, _, err := s.createContentSnapshotBackup(root); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	s.runMaintenanceCommand(c, "git-sync", s.cfg.GitSyncCommand, []string{"GIT_SYNC_RUN_BACKUP=false"})
}

func (s *Server) runContentSnapshotBackup(c *gin.Context) {
	root, err := s.projectRoot()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	start := time.Now()
	snapshotPath, backupDir, err := s.createContentSnapshotBackup(root)
	finished := time.Now()
	result := MaintenanceCommandResult{
		OK:         err == nil,
		Action:     "backup",
		StartedAt:  start.Format(time.RFC3339),
		FinishedAt: finished.Format(time.RFC3339),
		DurationMs: finished.Sub(start).Milliseconds(),
	}
	if err != nil {
		result.Output = err.Error()
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "result": result})
		return
	}
	result.Output = strings.Join([]string{
		"Refreshed git-visible CMS snapshot:",
		snapshotPath,
		"",
		"Created local recovery copy:",
		backupDir,
		"",
		"Note: the admin button exports CMS content directly through Go, so it does not depend on mysql/mysqldump command-line clients.",
	}, "\n")
	c.JSON(http.StatusOK, result)
}

func (s *Server) createContentSnapshotBackup(root string) (string, string, error) {
	snapshotPath, err := s.refreshContentSnapshot(root)
	if err != nil {
		return "", "", err
	}
	backupDir := filepath.Join(root, "storage", "backups", time.Now().Format("20060102-150405"))
	if err := os.MkdirAll(backupDir, 0o755); err != nil {
		return "", "", err
	}
	if err := copyFile(snapshotPath, filepath.Join(backupDir, "content.json")); err != nil {
		return "", "", err
	}
	readme := fmt.Sprintf(`BYML CMS content backup created at %s

Files:
- content.json: JSON export of CMS content tables, including draft and soft-deleted rows.

This admin-triggered backup refreshes the git-visible recovery snapshot without using mysql/mysqldump CLI clients.
The scheduled backup service can still create SQL dumps when its container has a compatible MySQL client.
`, time.Now().Format(time.RFC3339))
	if err := os.WriteFile(filepath.Join(backupDir, "README.txt"), []byte(readme), 0o644); err != nil {
		return "", "", err
	}
	pruneLocalBackups(filepath.Join(root, "storage", "backups"), parseInt64Env("BACKUP_KEEP_COUNT", 28))
	return snapshotPath, backupDir, nil
}

func (s *Server) refreshContentSnapshot(root string) (string, error) {
	_, snapshotPath, err := s.exportContentSnapshot(root)
	return snapshotPath, err
}

func (s *Server) runMaintenanceCommand(c *gin.Context, action string, configuredCommand string, extraEnv []string) {
	root, err := s.projectRoot()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	commandPath := configuredCommand
	if strings.TrimSpace(commandPath) == "" {
		commandPath = filepath.Join(root, "scripts", commandFile(action))
	}
	if !filepath.IsAbs(commandPath) {
		commandPath = filepath.Join(root, commandPath)
	}
	commandPath = filepath.Clean(commandPath)
	if info, err := os.Stat(commandPath); err != nil || info.IsDir() {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("maintenance command %q is not available", commandPath)})
		return
	}

	start := time.Now()
	ctx, cancel := context.WithTimeout(c.Request.Context(), s.cfg.MaintenanceCommandTimeout)
	defer cancel()

	cmd := exec.CommandContext(ctx, commandPath)
	cmd.Dir = root
	cmd.Env = append(os.Environ(), extraEnv...)
	output, runErr := cmd.CombinedOutput()
	finished := time.Now()
	result := MaintenanceCommandResult{
		OK:         runErr == nil,
		Action:     action,
		Output:     trimCommandOutput(string(output), 6000),
		StartedAt:  start.Format(time.RFC3339),
		FinishedAt: finished.Format(time.RFC3339),
		DurationMs: finished.Sub(start).Milliseconds(),
	}

	if errors.Is(ctx.Err(), context.DeadlineExceeded) {
		result.OK = false
		c.JSON(http.StatusGatewayTimeout, gin.H{"error": "maintenance command timed out", "result": result})
		return
	}
	if runErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": runErr.Error(), "result": result})
		return
	}
	c.JSON(http.StatusOK, result)
}

func (s *Server) projectRoot() (string, error) {
	root := strings.TrimSpace(s.cfg.ProjectRoot)
	if root == "" {
		root = ".."
	}
	return filepath.Abs(root)
}

func commandFile(action string) string {
	if action == "git-sync" {
		return "git-sync-backup.sh"
	}
	return "backup.sh"
}

func fileStatus(path string) MaintenanceFileStatus {
	status := MaintenanceFileStatus{Path: path}
	info, err := os.Stat(path)
	if err != nil || info.IsDir() {
		return status
	}
	status.Exists = true
	status.Size = info.Size()
	status.UpdatedAt = info.ModTime().Format(time.RFC3339)
	return status
}

func dirStatus(path string) MaintenanceDirStatus {
	status := MaintenanceDirStatus{Path: path}
	info, err := os.Stat(path)
	if err != nil || !info.IsDir() {
		return status
	}
	status.Exists = true
	_ = filepath.WalkDir(path, func(entryPath string, entry os.DirEntry, walkErr error) error {
		if walkErr != nil || entry.IsDir() {
			return nil
		}
		info, err := entry.Info()
		if err != nil {
			return nil
		}
		status.FileCount++
		status.TotalSize += info.Size()
		_ = entryPath
		return nil
	})
	return status
}

func backupStatus(path string) MaintenanceBackupStatus {
	status := MaintenanceBackupStatus{
		Path:      path,
		KeepCount: parseInt64Env("BACKUP_KEEP_COUNT", 28),
		Dirs:      []string{},
	}
	entries, err := os.ReadDir(path)
	if err != nil {
		return status
	}

	var newestModTime time.Time
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		if _, err := time.ParseInLocation("20060102-150405", entry.Name(), time.Local); err != nil {
			continue
		}
		status.Dirs = append(status.Dirs, entry.Name())
		info, err := entry.Info()
		if err == nil && (newestModTime.IsZero() || info.ModTime().After(newestModTime)) {
			newestModTime = info.ModTime()
			status.Newest = entry.Name()
			status.NewestPath = filepath.Join(path, entry.Name())
		}
	}
	sort.Strings(status.Dirs)
	status.Count = len(status.Dirs)
	return status
}

func pruneLocalBackups(path string, keepCount int64) {
	if keepCount <= 0 {
		return
	}
	entries, err := os.ReadDir(path)
	if err != nil {
		return
	}
	dirs := make([]backupDir, 0, len(entries))
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		if _, err := time.ParseInLocation("20060102-150405", entry.Name(), time.Local); err != nil {
			continue
		}
		info, err := entry.Info()
		if err != nil {
			continue
		}
		dirs = append(dirs, backupDir{name: entry.Name(), modTime: info.ModTime()})
	}
	sort.Slice(dirs, func(i, j int) bool {
		return dirs[i].modTime.Before(dirs[j].modTime)
	})
	removeCount := int64(len(dirs)) - keepCount
	if removeCount <= 0 {
		return
	}
	for _, dir := range dirs[:removeCount] {
		_ = os.RemoveAll(filepath.Join(path, dir.name))
	}
}

type backupDir struct {
	name    string
	modTime time.Time
}

func copyFile(src string, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	if _, err := io.Copy(out, in); err != nil {
		_ = out.Close()
		return err
	}
	return out.Close()
}

func parseInt64Env(key string, fallback int64) int64 {
	value := strings.TrimSpace(os.Getenv(key))
	if value == "" {
		return fallback
	}
	var parsed int64
	if _, err := fmt.Sscanf(value, "%d", &parsed); err != nil {
		return fallback
	}
	return parsed
}

func gitStatus(root string) MaintenanceGitStatus {
	status := MaintenanceGitStatus{Changes: []string{}}
	branchCmd := exec.Command("git", "-c", "safe.directory="+root, "branch", "--show-current")
	branchCmd.Dir = root
	branchOut, err := branchCmd.Output()
	if err != nil {
		status.Error = err.Error()
		return status
	}
	status.Available = true
	status.Branch = strings.TrimSpace(string(branchOut))

	statusCmd := exec.Command("git", "-c", "safe.directory="+root, "status", "--short", "--", "storage/content/content.json", "storage/uploads")
	statusCmd.Dir = root
	statusOut, err := statusCmd.Output()
	if err != nil {
		status.Error = err.Error()
		return status
	}
	for _, line := range strings.Split(strings.TrimSpace(string(statusOut)), "\n") {
		line = strings.TrimSpace(line)
		if line != "" {
			status.Changes = append(status.Changes, line)
		}
	}
	return status
}

func trimCommandOutput(output string, limit int) string {
	output = strings.TrimSpace(output)
	if len(output) <= limit {
		return output
	}
	return "... output truncated ...\n" + output[len(output)-limit:]
}
