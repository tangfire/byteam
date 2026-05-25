package app

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TrashItem struct {
	Resource  string     `json:"resource"`
	Label     string     `json:"label"`
	ID        uint       `json:"id"`
	Title     string     `json:"title"`
	Subtitle  string     `json:"subtitle"`
	Status    string     `json:"status"`
	DeletedAt *time.Time `json:"deletedAt,omitempty"`
}

func (s *Server) listTrash(c *gin.Context) {
	resource := strings.TrimSpace(c.Query("resource"))
	if resource == "" {
		resource = "news"
	}

	page := queryInt(c, "page", 1)
	pageSize := queryInt(c, "pageSize", 20)
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 20
	}
	if pageSize > 100 {
		pageSize = 100
	}

	items, total, ok, err := s.deletedItems(resource, page, pageSize, strings.TrimSpace(c.Query("q")))
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "unsupported trash resource"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"items":    items,
		"total":    total,
		"page":     page,
		"pageSize": pageSize,
	})
}

func (s *Server) restoreTrash(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		return
	}

	restored, supported, err := s.restoreDeleted(c.Param("resource"), id)
	if !supported {
		c.JSON(http.StatusBadRequest, gin.H{"error": "unsupported trash resource"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if !restored {
		c.JSON(http.StatusNotFound, gin.H{"error": "deleted item not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"restored": true})
}

func (s *Server) deletedItems(resource string, page int, pageSize int, q string) ([]TrashItem, int64, bool, error) {
	switch resource {
	case "pages":
		return listDeleted(s.db, resource, "Pages", page, pageSize, q, []string{"slug", "title", "description"}, func(item SitePage) TrashItem {
			return trashItem(resource, "Pages", item.ID, item.Title, compactSubtitle(item.Slug, item.Description), item.Status, item.DeletedAt)
		})
	case "news":
		return listDeleted(s.db, resource, "News", page, pageSize, q, []string{"title", "content", "excerpt"}, func(item NewsItem) TrashItem {
			return trashItem(resource, "News", item.ID, item.Title, compactSubtitle(item.TypeLabel, item.EventDate), item.Status, item.DeletedAt)
		})
	case "people":
		return listDeleted(s.db, resource, "People", page, pageSize, q, []string{"name", "category", "research"}, func(item Person) TrashItem {
			return trashItem(resource, "People", item.ID, item.Name, compactSubtitle(item.Category, item.Research), item.Status, item.DeletedAt)
		})
	case "undergraduates":
		return listDeleted(s.db, resource, "Undergraduates", page, pageSize, q, []string{"name", "grade", "major", "direction"}, func(item UndergraduateEducation) TrashItem {
			return trashItem(resource, "Undergraduates", item.ID, item.Name, compactSubtitle(item.Grade, item.Major), item.Status, item.DeletedAt)
		})
	case "publications":
		return listDeleted(s.db, resource, "Publications", page, pageSize, q, []string{"title", "authors", "venue"}, func(item Publication) TrashItem {
			return trashItem(resource, "Publications", item.ID, item.Title, item.Venue, item.Status, item.DeletedAt)
		})
	case "patents":
		return listDeleted(s.db, resource, "Patents", page, pageSize, q, []string{"authors", "title", "number", "country"}, func(item Patent) TrashItem {
			return trashItem(resource, "Patents", item.ID, item.Title, compactSubtitle(item.Number, item.Country), item.Status, item.DeletedAt)
		})
	case "research-projects":
		return listDeleted(s.db, resource, "Research Projects", page, pageSize, q, []string{"title", "fund", "number", "role"}, func(item ResearchProject) TrashItem {
			return trashItem(resource, "Research Projects", item.ID, item.Title, compactSubtitle(item.Fund, item.Number), item.Status, item.DeletedAt)
		})
	case "media":
		return listDeleted(s.db, resource, "Media", page, pageSize, q, []string{"file_name", "original_name", "mime_type", "url"}, func(item MediaAsset) TrashItem {
			return trashItem(resource, "Media", item.ID, item.OriginalName, compactSubtitle(item.Kind, item.URL), "", item.DeletedAt)
		})
	default:
		return nil, 0, false, nil
	}
}

func listDeleted[T any](db *gorm.DB, resource string, label string, page int, pageSize int, q string, searchable []string, mapItem func(T) TrashItem) ([]TrashItem, int64, bool, error) {
	query := db.Unscoped().Model(new(T)).Where("deleted_at IS NOT NULL")
	if q != "" {
		like := "%" + q + "%"
		conditions := make([]string, 0, len(searchable))
		args := make([]any, 0, len(searchable))
		for _, field := range searchable {
			conditions = append(conditions, field+" LIKE ?")
			args = append(args, like)
		}
		query = query.Where(strings.Join(conditions, " OR "), args...)
	}

	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, true, err
	}

	var rows []T
	if err := query.Order("deleted_at DESC, id DESC").Offset((page - 1) * pageSize).Limit(pageSize).Find(&rows).Error; err != nil {
		return nil, 0, true, err
	}

	items := make([]TrashItem, 0, len(rows))
	for _, row := range rows {
		item := mapItem(row)
		item.Resource = resource
		item.Label = label
		items = append(items, item)
	}
	return items, total, true, nil
}

func (s *Server) restoreDeleted(resource string, id uint) (bool, bool, error) {
	switch resource {
	case "pages":
		return restoreDeletedModel[SitePage](s.db, id)
	case "news":
		return restoreDeletedModel[NewsItem](s.db, id)
	case "people":
		return restoreDeletedModel[Person](s.db, id)
	case "undergraduates":
		return restoreDeletedModel[UndergraduateEducation](s.db, id)
	case "publications":
		return restoreDeletedModel[Publication](s.db, id)
	case "patents":
		return restoreDeletedModel[Patent](s.db, id)
	case "research-projects":
		return restoreDeletedModel[ResearchProject](s.db, id)
	case "media":
		return restoreDeletedModel[MediaAsset](s.db, id)
	default:
		return false, false, nil
	}
}

func restoreDeletedModel[T any](db *gorm.DB, id uint) (bool, bool, error) {
	result := db.Unscoped().
		Model(new(T)).
		Where("id = ? AND deleted_at IS NOT NULL", id).
		Update("deleted_at", nil)
	if result.Error != nil {
		return false, true, result.Error
	}
	return result.RowsAffected > 0, true, nil
}

func trashItem(resource string, label string, id uint, title string, subtitle string, status string, deletedAt gorm.DeletedAt) TrashItem {
	var deleted *time.Time
	if deletedAt.Valid {
		deleted = &deletedAt.Time
	}
	return TrashItem{
		Resource:  resource,
		Label:     label,
		ID:        id,
		Title:     title,
		Subtitle:  subtitle,
		Status:    status,
		DeletedAt: deleted,
	}
}

func compactSubtitle(parts ...string) string {
	cleaned := make([]string, 0, len(parts))
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part != "" {
			cleaned = append(cleaned, part)
		}
	}
	return strings.Join(cleaned, " / ")
}
