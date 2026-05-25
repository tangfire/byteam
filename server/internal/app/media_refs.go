package app

import (
	"os"
	"path/filepath"
	"strings"

	"gorm.io/gorm"
)

func (s *Server) mediaURLInUse(url string) bool {
	if mediaURLReferenceCount(s.db, url) > 0 {
		return true
	}
	return s.staticAssetInUse(url)
}

func mediaURLReferenceCount(db *gorm.DB, url string) int64 {
	var count int64
	var total int64
	db.Model(&Person{}).Where("avatar_url = ?", url).Count(&count)
	total += count
	db.Model(&Publication{}).Where("image_url = ?", url).Count(&count)
	total += count
	db.Model(&PublicationLink{}).Where("url = ?", url).Count(&count)
	total += count
	db.Model(&SitePage{}).Where("JSON_SEARCH(content, 'one', ?) IS NOT NULL", url).Count(&count)
	total += count
	return total
}

func (s *Server) staticAssetInUse(url string) bool {
	return s.staticAssetSource(url) != ""
}

func (s *Server) staticAssetSource(url string) string {
	url = strings.TrimSpace(url)
	if url == "" || !strings.HasPrefix(url, "/") || strings.HasPrefix(url, "/uploads/") {
		return ""
	}
	root := strings.TrimSpace(s.cfg.ProjectRoot)
	if root == "" {
		return ""
	}
	for _, dir := range []string{"src", "public"} {
		source, err := findTextSourceInDir(filepath.Join(root, dir), url)
		if err == nil && source != "" {
			return formatStaticSource(root, source)
		}
	}
	return ""
}

func textExistsInDir(root string, needle string) (bool, error) {
	source, err := findTextSourceInDir(root, needle)
	return source != "", err
}

func findTextSourceInDir(root string, needle string) (string, error) {
	info, err := os.Stat(root)
	if err != nil || !info.IsDir() {
		return "", err
	}
	var foundPath string
	err = filepath.WalkDir(root, func(path string, entry os.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if entry.IsDir() {
			name := entry.Name()
			if name == "node_modules" || name == "dist" || name == ".git" {
				return filepath.SkipDir
			}
			return nil
		}
		ext := strings.ToLower(filepath.Ext(entry.Name()))
		switch ext {
		case ".vue", ".ts", ".js", ".json", ".html", ".css", ".md":
		default:
			return nil
		}
		raw, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		if strings.Contains(string(raw), needle) {
			foundPath = path
			return errFoundNeedle{}
		}
		return nil
	})
	if err != nil {
		if _, ok := err.(errFoundNeedle); ok {
			return foundPath, nil
		}
		return "", err
	}
	return "", nil
}

func formatStaticSource(root string, path string) string {
	relative, err := filepath.Rel(root, path)
	if err != nil {
		relative = path
	}
	relative = filepath.ToSlash(relative)
	base := strings.TrimSuffix(filepath.Base(relative), filepath.Ext(relative))
	base = strings.TrimPrefix(base, "VideoPlayer")
	base = strings.TrimSuffix(base, "View")
	base = strings.ReplaceAll(base, "_", " ")
	base = strings.ReplaceAll(base, "-", " ")
	base = strings.TrimSpace(base)
	if base == "" {
		return "源码引用"
	}
	return "源码引用 - " + base
}

type errFoundNeedle struct{}

func (errFoundNeedle) Error() string {
	return "found"
}
