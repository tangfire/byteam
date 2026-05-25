package app

import (
	"errors"
	"fmt"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var allowedExtensions = map[string]string{
	".jpg":  "image",
	".jpeg": "image",
	".png":  "image",
	".gif":  "image",
	".webp": "image",
	".svg":  "image",
	".pdf":  "document",
	".ppt":  "document",
	".pptx": "document",
	".zip":  "archive",
	".mp4":  "video",
}

func (s *Server) listMedia(c *gin.Context) {
	var items []MediaAsset
	db, page, pageSize := applyListQuery(c, s.db.Model(&MediaAsset{}), "file_name", "original_name", "mime_type")
	if kind := c.Query("kind"); kind != "" {
		db = db.Where("kind = ?", kind)
	}
	paged(c, db.Order("created_at DESC, id DESC"), &items, page, pageSize)
}

func (s *Server) importPublicMedia(c *gin.Context) {
	result, err := s.scanPublicMedia()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

type MediaImportResult struct {
	Scanned int `json:"scanned"`
	Created int `json:"created"`
	Updated int `json:"updated"`
	Skipped int `json:"skipped"`
}

func (s *Server) scanPublicMedia() (MediaImportResult, error) {
	result := MediaImportResult{}
	publicDir := filepath.Clean(s.cfg.PublicDir)
	info, err := os.Stat(publicDir)
	if err != nil {
		return result, fmt.Errorf("public dir %q is not available: %w", publicDir, err)
	}
	if !info.IsDir() {
		return result, fmt.Errorf("public dir %q is not a directory", publicDir)
	}

	err = filepath.WalkDir(publicDir, func(path string, entry os.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if entry.IsDir() {
			return nil
		}

		ext := strings.ToLower(filepath.Ext(entry.Name()))
		kind, ok := allowedExtensions[ext]
		if !ok {
			result.Skipped++
			return nil
		}
		result.Scanned++

		relative, err := filepath.Rel(publicDir, path)
		if err != nil {
			return err
		}
		relative = filepath.ToSlash(relative)
		url := "/" + relative
		fileInfo, err := entry.Info()
		if err != nil {
			return err
		}

		mimeType := mediaMimeType(ext)
		if mimeType == "" {
			mimeType = "application/octet-stream"
		}

		asset := MediaAsset{
			FileName:     entry.Name(),
			OriginalName: entry.Name(),
			URL:          url,
			Path:         path,
			MimeType:     mimeType,
			Size:         fileInfo.Size(),
			Kind:         kind,
		}

		var existing MediaAsset
		err = s.db.Unscoped().Where("url = ?", url).First(&existing).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			if err := s.db.Create(&asset).Error; err != nil {
				return err
			}
			result.Created++
			return nil
		}
		if err != nil {
			return err
		}
		if err := s.db.Unscoped().Model(&existing).Updates(map[string]any{
			"file_name":     asset.FileName,
			"original_name": asset.OriginalName,
			"path":          asset.Path,
			"mime_type":     asset.MimeType,
			"size":          asset.Size,
			"kind":          asset.Kind,
			"deleted_at":    nil,
		}).Error; err != nil {
			return err
		}
		result.Updated++
		return nil
	})

	if err != nil {
		return result, err
	}
	return result, nil
}

func (s *Server) uploadMedia(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "file is required"})
		return
	}
	if file.Size > s.cfg.MaxUploadSize {
		c.JSON(http.StatusBadRequest, gin.H{"error": "file exceeds upload limit"})
		return
	}

	ext := strings.ToLower(filepath.Ext(file.Filename))
	kind, ok := allowedExtensions[ext]
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "unsupported file type"})
		return
	}

	now := time.Now()
	relativeDir := filepath.Join(fmt.Sprintf("%04d", now.Year()), fmt.Sprintf("%02d", now.Month()))
	targetDir := filepath.Join(s.cfg.UploadDir, relativeDir)
	if err := os.MkdirAll(targetDir, 0o755); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	safeName := sanitizeFileName(strings.TrimSuffix(file.Filename, ext))
	fileName := fmt.Sprintf("%d-%s%s", now.UnixNano(), safeName, ext)
	targetPath := filepath.Join(targetDir, fileName)
	if err := c.SaveUploadedFile(file, targetPath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	url := "/uploads/" + filepath.ToSlash(filepath.Join(relativeDir, fileName))
	if s.cfg.PublicBaseURL != "" {
		url = strings.TrimRight(s.cfg.PublicBaseURL, "/") + url
	}

	mimeType := file.Header.Get("Content-Type")
	if mimeType == "" || mimeType == "application/octet-stream" {
		mimeType = mediaMimeType(ext)
	}

	asset := MediaAsset{
		FileName:     fileName,
		OriginalName: file.Filename,
		URL:          url,
		Path:         targetPath,
		MimeType:     mimeType,
		Size:         file.Size,
		Kind:         kind,
	}
	if err := s.db.Create(&asset).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, asset)
}

func (s *Server) deleteMedia(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		return
	}
	var asset MediaAsset
	if err := s.db.First(&asset, id).Error; err != nil {
		notFoundOrError(c, err)
		return
	}
	if mediaURLInUse(s.db, asset.URL) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "media is still referenced by content"})
		return
	}
	if err := s.db.Delete(&asset).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"deleted": true})
}

func sanitizeFileName(name string) string {
	name = strings.ToLower(name)
	var b strings.Builder
	for _, r := range name {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') || r == '-' || r == '_' {
			b.WriteRune(r)
		} else if r == ' ' || r == '.' {
			b.WriteRune('-')
		}
	}
	out := strings.Trim(b.String(), "-_")
	if out == "" {
		return "file"
	}
	if len(out) > 80 {
		return out[:80]
	}
	return out
}

func mediaMimeType(ext string) string {
	if value := mime.TypeByExtension(ext); value != "" {
		return value
	}
	switch ext {
	case ".mp4":
		return "video/mp4"
	case ".ppt":
		return "application/vnd.ms-powerpoint"
	case ".pptx":
		return "application/vnd.openxmlformats-officedocument.presentationml.presentation"
	case ".zip":
		return "application/zip"
	default:
		return "application/octet-stream"
	}
}
