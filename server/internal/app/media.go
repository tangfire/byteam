package app

import (
	"fmt"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
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
	if mimeType == "" {
		mimeType = mime.TypeByExtension(ext)
	}
	if mimeType == "" {
		mimeType = "application/octet-stream"
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
