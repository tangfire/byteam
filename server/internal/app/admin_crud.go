package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (s *Server) adminSummary(c *gin.Context) {
	counts := map[string]int64{}
	counts["news"] = countModel(s, &NewsItem{})
	counts["people"] = countModel(s, &Person{})
	counts["undergraduates"] = countModel(s, &UndergraduateEducation{})
	counts["publications"] = countModel(s, &Publication{})
	counts["patents"] = countModel(s, &Patent{})
	counts["researchProjects"] = countModel(s, &ResearchProject{})
	counts["media"] = countModel(s, &MediaAsset{})
	c.JSON(http.StatusOK, gin.H{"counts": counts})
}

func countModel(s *Server, model any) int64 {
	var count int64
	s.db.Model(model).Count(&count)
	return count
}

func (s *Server) listNews(c *gin.Context) {
	var items []NewsItem
	db, page, pageSize := applyListQuery(c, s.db.Model(&NewsItem{}), "title", "content", "excerpt")
	paged(c, db.Order("event_date DESC, sort_order ASC, id DESC"), &items, page, pageSize)
}

func (s *Server) createNews(c *gin.Context) {
	payload, ok := bindJSON[NewsItem](c)
	if !ok {
		return
	}
	payload.ID = 0
	payload.Status = normalizeStatus(payload.Status)
	if payload.Color == "" {
		payload.Color = "#7d1231"
	}
	if payload.TypeLabel == "" {
		payload.TypeLabel = newsTypeLabel(payload.Type)
	}
	if err := s.db.Create(&payload).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, payload)
}

func (s *Server) updateNews(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		return
	}
	payload, ok := bindJSON[NewsItem](c)
	if !ok {
		return
	}
	var item NewsItem
	if err := s.db.First(&item, id).Error; err != nil {
		notFoundOrError(c, err)
		return
	}
	payload.ID = item.ID
	payload.CreatedAt = item.CreatedAt
	payload.Status = normalizeStatus(payload.Status)
	if payload.TypeLabel == "" {
		payload.TypeLabel = newsTypeLabel(payload.Type)
	}
	if err := s.db.Save(&payload).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	s.db.First(&item, id)
	c.JSON(http.StatusOK, item)
}

func (s *Server) deleteNews(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		return
	}
	c.JSON(http.StatusOK, gin.H{"deleted": s.db.Delete(&NewsItem{}, id).Error == nil})
}

func newsTypeLabel(kind string) string {
	switch kind {
	case "publication":
		return "Publication"
	case "team":
		return "Team Update"
	case "award":
		return "Award"
	case "event":
		return "Event"
	default:
		return "General"
	}
}

func (s *Server) listPeople(c *gin.Context) {
	var items []Person
	db, page, pageSize := applyListQuery(c, s.db.Model(&Person{}), "name", "research", "category")
	if category := c.Query("category"); category != "" {
		db = db.Where("category = ?", category)
	}
	paged(c, db.Order("category ASC, sort_order ASC, id ASC"), &items, page, pageSize)
}

func (s *Server) createPerson(c *gin.Context) {
	payload, ok := bindJSON[Person](c)
	if !ok {
		return
	}
	payload.ID = 0
	payload.Status = normalizeStatus(payload.Status)
	if err := s.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&payload).Error; err != nil {
			return err
		}
		return renumberSortableRowsWithFirst(tx, &Person{}, tx.Model(&Person{}).Where("category = ?", payload.Category), payload.ID)
	}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, payload)
}

func (s *Server) updatePerson(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		return
	}
	payload, ok := bindJSON[Person](c)
	if !ok {
		return
	}
	var item Person
	if err := s.db.First(&item, id).Error; err != nil {
		notFoundOrError(c, err)
		return
	}
	payload.ID = item.ID
	payload.CreatedAt = item.CreatedAt
	payload.Status = normalizeStatus(payload.Status)
	if err := s.db.Transaction(func(tx *gorm.DB) error {
		groupChanged := item.Category != payload.Category
		payload.SortOrder = item.SortOrder
		if groupChanged {
			payload.SortOrder = 0
		}
		if err := tx.Save(&payload).Error; err != nil {
			return err
		}
		if groupChanged {
			if err := renumberSortableRowsWithFirst(tx, &Person{}, tx.Model(&Person{}).Where("category = ?", payload.Category), payload.ID); err != nil {
				return err
			}
			return renumberSortableGroup(tx, &Person{}, tx.Model(&Person{}).Where("category = ?", item.Category))
		}
		return nil
	}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	s.db.First(&item, id)
	c.JSON(http.StatusOK, item)
}

func (s *Server) deletePerson(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		return
	}
	c.JSON(http.StatusOK, gin.H{"deleted": s.db.Delete(&Person{}, id).Error == nil})
}

func (s *Server) listUndergraduates(c *gin.Context) {
	var items []UndergraduateEducation
	db, page, pageSize := applyListQuery(c, s.db.Model(&UndergraduateEducation{}), "name", "grade", "major", "direction")
	paged(c, db.Order("sort_order ASC, id ASC"), &items, page, pageSize)
}

func (s *Server) createUndergraduate(c *gin.Context) {
	payload, ok := bindJSON[UndergraduateEducation](c)
	if !ok {
		return
	}
	payload.ID = 0
	payload.Status = normalizeStatus(payload.Status)
	if err := s.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&payload).Error; err != nil {
			return err
		}
		return renumberSortableRowsWithFirst(tx, &UndergraduateEducation{}, tx.Model(&UndergraduateEducation{}), payload.ID)
	}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, payload)
}

func (s *Server) updateUndergraduate(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		return
	}
	payload, ok := bindJSON[UndergraduateEducation](c)
	if !ok {
		return
	}
	var item UndergraduateEducation
	if err := s.db.First(&item, id).Error; err != nil {
		notFoundOrError(c, err)
		return
	}
	payload.ID = item.ID
	payload.CreatedAt = item.CreatedAt
	payload.Status = normalizeStatus(payload.Status)
	payload.SortOrder = item.SortOrder
	if err := s.db.Save(&payload).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	s.db.First(&item, id)
	c.JSON(http.StatusOK, item)
}

func (s *Server) deleteUndergraduate(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		return
	}
	c.JSON(http.StatusOK, gin.H{"deleted": s.db.Delete(&UndergraduateEducation{}, id).Error == nil})
}

func (s *Server) listPublications(c *gin.Context) {
	var items []Publication
	db, page, pageSize := applyListQuery(c, s.db.Model(&Publication{}).Preload("Links"), "title", "authors", "venue")
	paged(c, db.Order("year DESC, kind ASC, sort_order ASC, id ASC"), &items, page, pageSize)
}

func (s *Server) createPublication(c *gin.Context) {
	payload, ok := bindJSON[Publication](c)
	if !ok {
		return
	}
	payload.ID = 0
	payload.Status = normalizeStatus(payload.Status)
	if err := s.createPublicationAtTop(&payload); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	s.db.Preload("Links").First(&payload, payload.ID)
	c.JSON(http.StatusCreated, payload)
}

func (s *Server) createPublicationAtTop(payload *Publication) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		kind := normalizePublicationKind(payload.Kind)
		payload.Kind = kind
		payload.SortOrder = 0
		if err := tx.Create(payload).Error; err != nil {
			return err
		}
		var group []Publication
		if err := tx.Where("year = ? AND kind = ? AND id <> ?", payload.Year, kind, payload.ID).
			Order("sort_order ASC, id ASC").
			Find(&group).Error; err != nil {
			return err
		}
		return renumberPublications(tx, append([]Publication{*payload}, group...))
	})
}

func (s *Server) updatePublication(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		return
	}
	payload, ok := bindJSON[Publication](c)
	if !ok {
		return
	}
	var item Publication
	if err := s.db.Preload("Links").First(&item, id).Error; err != nil {
		notFoundOrError(c, err)
		return
	}
	payload.ID = item.ID
	payload.CreatedAt = item.CreatedAt
	payload.Status = normalizeStatus(payload.Status)
	payload.Kind = normalizePublicationKind(payload.Kind)
	if err := s.db.Transaction(func(tx *gorm.DB) error {
		groupChanged := item.Year != payload.Year || item.Kind != payload.Kind
		sortOrder := item.SortOrder
		if payload.SortOrder > 0 {
			sortOrder = payload.SortOrder
		}
		if groupChanged {
			sortOrder = 0
		}
		if err := tx.Model(&item).Updates(map[string]any{
			"image_url":  payload.ImageURL,
			"title":      payload.Title,
			"authors":    payload.Authors,
			"venue":      payload.Venue,
			"year":       payload.Year,
			"kind":       payload.Kind,
			"status":     payload.Status,
			"featured":   payload.Featured,
			"sort_order": sortOrder,
		}).Error; err != nil {
			return err
		}
		if err := tx.Where("publication_id = ?", item.ID).Delete(&PublicationLink{}).Error; err != nil {
			return err
		}
		for i := range payload.Links {
			payload.Links[i].ID = 0
			payload.Links[i].PublicationID = item.ID
		}
		if len(payload.Links) > 0 {
			return tx.Create(&payload.Links).Error
		}
		if err := renumberPublicationGroup(tx, payload.Year, payload.Kind); err != nil {
			return err
		}
		if groupChanged {
			return renumberPublicationGroup(tx, item.Year, item.Kind)
		}
		return nil
	}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	s.db.Preload("Links").First(&item, id)
	c.JSON(http.StatusOK, item)
}

func normalizePublicationKind(kind string) string {
	if kind == "journal" {
		return "journal"
	}
	return "conference"
}

func renumberPublicationGroup(tx *gorm.DB, year int, kind string) error {
	var group []Publication
	if err := tx.Where("year = ? AND kind = ?", year, kind).
		Order("sort_order ASC, id ASC").
		Find(&group).Error; err != nil {
		return err
	}
	return renumberPublications(tx, group)
}

func renumberPublications(tx *gorm.DB, group []Publication) error {
	for i := range group {
		if err := tx.Model(&Publication{}).
			Where("id = ?", group[i].ID).
			Update("sort_order", i+1).Error; err != nil {
			return err
		}
	}
	return nil
}

func (s *Server) deletePublication(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		return
	}
	c.JSON(http.StatusOK, gin.H{"deleted": s.db.Delete(&Publication{}, id).Error == nil})
}

func (s *Server) listPatents(c *gin.Context) {
	var items []Patent
	db, page, pageSize := applyListQuery(c, s.db.Model(&Patent{}), "authors", "title", "number", "country")
	paged(c, db.Order("category ASC, sort_order ASC, id ASC"), &items, page, pageSize)
}

func (s *Server) createPatent(c *gin.Context) {
	payload, ok := bindJSON[Patent](c)
	if !ok {
		return
	}
	payload.ID = 0
	payload.Status = normalizeStatus(payload.Status)
	if err := s.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&payload).Error; err != nil {
			return err
		}
		return renumberSortableRowsWithFirst(tx, &Patent{}, tx.Model(&Patent{}).Where("category = ?", payload.Category), payload.ID)
	}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, payload)
}

func (s *Server) updatePatent(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		return
	}
	payload, ok := bindJSON[Patent](c)
	if !ok {
		return
	}
	var item Patent
	if err := s.db.First(&item, id).Error; err != nil {
		notFoundOrError(c, err)
		return
	}
	payload.ID = item.ID
	payload.CreatedAt = item.CreatedAt
	payload.Status = normalizeStatus(payload.Status)
	if err := s.db.Transaction(func(tx *gorm.DB) error {
		groupChanged := item.Category != payload.Category
		payload.SortOrder = item.SortOrder
		if groupChanged {
			payload.SortOrder = 0
		}
		if err := tx.Save(&payload).Error; err != nil {
			return err
		}
		if groupChanged {
			if err := renumberSortableRowsWithFirst(tx, &Patent{}, tx.Model(&Patent{}).Where("category = ?", payload.Category), payload.ID); err != nil {
				return err
			}
			return renumberSortableGroup(tx, &Patent{}, tx.Model(&Patent{}).Where("category = ?", item.Category))
		}
		return nil
	}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	s.db.First(&item, id)
	c.JSON(http.StatusOK, item)
}

func (s *Server) deletePatent(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		return
	}
	c.JSON(http.StatusOK, gin.H{"deleted": s.db.Delete(&Patent{}, id).Error == nil})
}

func (s *Server) listResearchProjects(c *gin.Context) {
	var items []ResearchProject
	db, page, pageSize := applyListQuery(c, s.db.Model(&ResearchProject{}), "title", "fund", "number", "role")
	paged(c, db.Order("sort_order ASC, id ASC"), &items, page, pageSize)
}

func (s *Server) createResearchProject(c *gin.Context) {
	payload, ok := bindJSON[ResearchProject](c)
	if !ok {
		return
	}
	payload.ID = 0
	payload.Status = normalizeStatus(payload.Status)
	if err := s.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&payload).Error; err != nil {
			return err
		}
		return renumberSortableRowsWithFirst(tx, &ResearchProject{}, tx.Model(&ResearchProject{}), payload.ID)
	}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, payload)
}

func (s *Server) updateResearchProject(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		return
	}
	payload, ok := bindJSON[ResearchProject](c)
	if !ok {
		return
	}
	var item ResearchProject
	if err := s.db.First(&item, id).Error; err != nil {
		notFoundOrError(c, err)
		return
	}
	payload.ID = item.ID
	payload.CreatedAt = item.CreatedAt
	payload.Status = normalizeStatus(payload.Status)
	payload.SortOrder = item.SortOrder
	if err := s.db.Save(&payload).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	s.db.First(&item, id)
	c.JSON(http.StatusOK, item)
}

func (s *Server) deleteResearchProject(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		return
	}
	c.JSON(http.StatusOK, gin.H{"deleted": s.db.Delete(&ResearchProject{}, id).Error == nil})
}
