package app

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

var sourceManagedSitePageSlugs = []string{"about", "contact", "videomind", "vknow"}

func (s *Server) publicSitePage(c *gin.Context) {
	slug := strings.TrimSpace(c.Param("slug"))
	var page SitePage
	if err := s.db.Where("slug = ? AND status = ?", slug, StatusPublished).First(&page).Error; err != nil {
		notFoundOrError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"page": page})
}

func (s *Server) listSitePages(c *gin.Context) {
	var items []SitePage
	db, page, pageSize := applyListQuery(c, s.db.Model(&SitePage{}), "slug", "title", "description")
	db = db.Where("slug NOT IN ?", sourceManagedSitePageSlugs)
	paged(c, db.Order("sort_order ASC, id ASC"), &items, page, pageSize)
}

func (s *Server) getSitePage(c *gin.Context) {
	slug := strings.TrimSpace(c.Param("slug"))
	var page SitePage
	if err := s.db.Where("slug = ?", slug).First(&page).Error; err != nil {
		notFoundOrError(c, err)
		return
	}
	c.JSON(http.StatusOK, page)
}

func (s *Server) updateSitePage(c *gin.Context) {
	slug := strings.TrimSpace(c.Param("slug"))
	payload, ok := bindJSON[SitePage](c)
	if !ok {
		return
	}

	var page SitePage
	if err := s.db.Where("slug = ?", slug).First(&page).Error; err != nil {
		notFoundOrError(c, err)
		return
	}

	title := strings.TrimSpace(payload.Title)
	if title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "title is required"})
		return
	}
	content := payload.Content
	if content == nil {
		content = JSONMap{}
	}

	if err := s.db.Model(&page).Updates(map[string]any{
		"title":       title,
		"description": strings.TrimSpace(payload.Description),
		"content":     content,
		"status":      normalizeStatus(payload.Status),
	}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	s.db.Where("slug = ?", slug).First(&page)
	c.JSON(http.StatusOK, page)
}
