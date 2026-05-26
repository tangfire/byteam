package app

import (
	"net/http"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
)

var sourceManagedSitePageSlugs = []string{"about", "contact", "videomind", "vknow"}
var videoPageSlugPattern = regexp.MustCompile(`^video-[a-z0-9]+(?:-[a-z0-9]+)*$`)

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

func (s *Server) createSitePage(c *gin.Context) {
	payload, ok := bindJSON[SitePage](c)
	if !ok {
		return
	}

	slug := strings.ToLower(strings.TrimSpace(payload.Slug))
	if !videoPageSlugPattern.MatchString(slug) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "video page slug must look like video-paper-title"})
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
	if strings.TrimSpace(asString(content["title"])) == "" {
		content["title"] = title
	}

	page := SitePage{
		Slug:        slug,
		Title:       title,
		Description: strings.TrimSpace(payload.Description),
		Content:     content,
		Status:      normalizeStatus(payload.Status),
		SortOrder:   nextSitePageSortOrder(s),
	}
	if page.Description == "" {
		page.Description = "Publication video"
	}

	if err := s.db.Create(&page).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, page)
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
	if isSourceManagedSitePage(slug) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "this page is source managed"})
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

func (s *Server) deleteSitePage(c *gin.Context) {
	slug := strings.TrimSpace(c.Param("slug"))
	if isSourceManagedSitePage(slug) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "this page is source managed"})
		return
	}
	if !strings.HasPrefix(slug, "video-") {
		c.JSON(http.StatusBadRequest, gin.H{"error": "only video pages can be deleted from this editor"})
		return
	}

	var page SitePage
	if err := s.db.Where("slug = ?", slug).First(&page).Error; err != nil {
		notFoundOrError(c, err)
		return
	}
	if err := s.db.Delete(&page).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"deleted": true})
}

func isSourceManagedSitePage(slug string) bool {
	for _, item := range sourceManagedSitePageSlugs {
		if slug == item {
			return true
		}
	}
	return false
}

func nextSitePageSortOrder(s *Server) int {
	var maxOrder int
	s.db.Model(&SitePage{}).Select("COALESCE(MAX(sort_order), 0)").Scan(&maxOrder)
	return maxOrder + 1
}

func asString(value any) string {
	if text, ok := value.(string); ok {
		return text
	}
	return ""
}
