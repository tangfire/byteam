package app

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type placeRequest struct {
	TargetID uint   `json:"targetId"`
	Position string `json:"position"`
}

type sortableRow struct {
	ID        uint
	SortOrder int
}

func bindPlaceRequest(c *gin.Context) (placeRequest, bool) {
	payload, ok := bindJSON[placeRequest](c)
	if !ok {
		return placeRequest{}, false
	}
	if payload.TargetID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "targetId is required"})
		return placeRequest{}, false
	}
	if payload.Position != "after" {
		payload.Position = "before"
	}
	return payload, true
}

func placeSortableRows(tx *gorm.DB, model any, db *gorm.DB, itemID uint, targetID uint, position string) error {
	var rows []sortableRow
	if err := db.Order("sort_order ASC, id ASC").Find(&rows).Error; err != nil {
		return err
	}
	ordered, err := reorderSortableRows(rows, itemID, targetID, position)
	if err != nil {
		return err
	}
	return updateSortableRows(tx, model, ordered)
}

func renumberSortableRowsWithFirst(tx *gorm.DB, model any, db *gorm.DB, firstID uint) error {
	var rows []sortableRow
	if err := db.Where("id <> ?", firstID).Order("sort_order ASC, id ASC").Find(&rows).Error; err != nil {
		return err
	}
	rows = append([]sortableRow{{ID: firstID}}, rows...)
	return updateSortableRows(tx, model, rows)
}

func renumberSortableGroup(tx *gorm.DB, model any, db *gorm.DB) error {
	var rows []sortableRow
	if err := db.Order("sort_order ASC, id ASC").Find(&rows).Error; err != nil {
		return err
	}
	return updateSortableRows(tx, model, rows)
}

func reorderSortableRows(rows []sortableRow, itemID uint, targetID uint, position string) ([]sortableRow, error) {
	if itemID == targetID {
		return rows, nil
	}

	itemIndex := -1
	targetIndex := -1
	for index, row := range rows {
		if row.ID == itemID {
			itemIndex = index
		}
		if row.ID == targetID {
			targetIndex = index
		}
	}
	if itemIndex < 0 || targetIndex < 0 {
		return nil, gorm.ErrRecordNotFound
	}

	moved := rows[itemIndex]
	rows = append(rows[:itemIndex], rows[itemIndex+1:]...)
	if itemIndex < targetIndex {
		targetIndex--
	}
	if position == "after" {
		targetIndex++
	}
	if targetIndex < 0 {
		targetIndex = 0
	}
	if targetIndex > len(rows) {
		targetIndex = len(rows)
	}

	rows = append(rows[:targetIndex], append([]sortableRow{moved}, rows[targetIndex:]...)...)
	return rows, nil
}

func updateSortableRows(tx *gorm.DB, model any, rows []sortableRow) error {
	for index := range rows {
		if err := tx.Model(model).
			Where("id = ?", rows[index].ID).
			Update("sort_order", index+1).Error; err != nil {
			return err
		}
	}
	return nil
}

func badSortGroup(c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{"error": "只能在同一展示分组内拖动排序"})
}

func (s *Server) placePerson(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		return
	}
	payload, ok := bindPlaceRequest(c)
	if !ok {
		return
	}

	var item Person
	var target Person
	if err := s.db.First(&item, id).Error; err != nil {
		notFoundOrError(c, err)
		return
	}
	if err := s.db.First(&target, payload.TargetID).Error; err != nil {
		notFoundOrError(c, err)
		return
	}
	if item.Category != target.Category {
		badSortGroup(c)
		return
	}

	if err := s.db.Transaction(func(tx *gorm.DB) error {
		return placeSortableRows(tx, &Person{}, tx.Model(&Person{}).Where("category = ?", item.Category), item.ID, target.ID, payload.Position)
	}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	s.db.First(&item, id)
	c.JSON(http.StatusOK, item)
}

func (s *Server) placeUndergraduate(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		return
	}
	payload, ok := bindPlaceRequest(c)
	if !ok {
		return
	}

	var item UndergraduateEducation
	var target UndergraduateEducation
	if err := s.db.First(&item, id).Error; err != nil {
		notFoundOrError(c, err)
		return
	}
	if err := s.db.First(&target, payload.TargetID).Error; err != nil {
		notFoundOrError(c, err)
		return
	}

	if err := s.db.Transaction(func(tx *gorm.DB) error {
		return placeSortableRows(tx, &UndergraduateEducation{}, tx.Model(&UndergraduateEducation{}), item.ID, target.ID, payload.Position)
	}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	s.db.First(&item, id)
	c.JSON(http.StatusOK, item)
}

func (s *Server) placePublication(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		return
	}
	payload, ok := bindPlaceRequest(c)
	if !ok {
		return
	}

	var item Publication
	var target Publication
	if err := s.db.First(&item, id).Error; err != nil {
		notFoundOrError(c, err)
		return
	}
	if err := s.db.First(&target, payload.TargetID).Error; err != nil {
		notFoundOrError(c, err)
		return
	}
	if item.Year != target.Year || item.Kind != target.Kind {
		badSortGroup(c)
		return
	}

	if err := s.db.Transaction(func(tx *gorm.DB) error {
		return placeSortableRows(tx, &Publication{}, tx.Model(&Publication{}).Where("year = ? AND kind = ?", item.Year, item.Kind), item.ID, target.ID, payload.Position)
	}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	preloadPublicationLinks(s.db).First(&item, id)
	c.JSON(http.StatusOK, item)
}

func (s *Server) placePatent(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		return
	}
	payload, ok := bindPlaceRequest(c)
	if !ok {
		return
	}

	var item Patent
	var target Patent
	if err := s.db.First(&item, id).Error; err != nil {
		notFoundOrError(c, err)
		return
	}
	if err := s.db.First(&target, payload.TargetID).Error; err != nil {
		notFoundOrError(c, err)
		return
	}
	if item.Category != target.Category {
		badSortGroup(c)
		return
	}

	if err := s.db.Transaction(func(tx *gorm.DB) error {
		return placeSortableRows(tx, &Patent{}, tx.Model(&Patent{}).Where("category = ?", item.Category), item.ID, target.ID, payload.Position)
	}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	s.db.First(&item, id)
	c.JSON(http.StatusOK, item)
}

func (s *Server) placeResearchProject(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		return
	}
	payload, ok := bindPlaceRequest(c)
	if !ok {
		return
	}

	var item ResearchProject
	var target ResearchProject
	if err := s.db.First(&item, id).Error; err != nil {
		notFoundOrError(c, err)
		return
	}
	if err := s.db.First(&target, payload.TargetID).Error; err != nil {
		notFoundOrError(c, err)
		return
	}

	if err := s.db.Transaction(func(tx *gorm.DB) error {
		return placeSortableRows(tx, &ResearchProject{}, tx.Model(&ResearchProject{}), item.ID, target.ID, payload.Position)
	}); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			notFoundOrError(c, err)
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	s.db.First(&item, id)
	c.JSON(http.StatusOK, item)
}
