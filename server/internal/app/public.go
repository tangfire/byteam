package app

import (
	"net/http"
	"sort"

	"github.com/gin-gonic/gin"
)

func (s *Server) publicHome(c *gin.Context) {
	var latestNews []NewsItem
	s.db.Where("status = ?", StatusPublished).
		Order("event_date DESC, sort_order ASC, id DESC").
		Limit(4).
		Find(&latestNews)

	var featured []Publication
	s.db.Preload("Links").
		Where("status = ? AND featured = ?", StatusPublished, true).
		Order("year DESC, sort_order ASC, id DESC").
		Limit(5).
		Find(&featured)
	if len(featured) == 0 {
		s.db.Preload("Links").
			Where("status = ?", StatusPublished).
			Order("year DESC, sort_order ASC, id DESC").
			Limit(5).
			Find(&featured)
	}

	var publicationsCount int64
	var projectsCount int64
	var peopleCount int64
	s.db.Model(&Publication{}).Where("status = ?", StatusPublished).Count(&publicationsCount)
	s.db.Model(&ResearchProject{}).Where("status = ?", StatusPublished).Count(&projectsCount)
	s.db.Model(&Person{}).Where("status = ? AND category = ?", StatusPublished, "graduate").Count(&peopleCount)

	c.JSON(http.StatusOK, gin.H{
		"latestNews":           latestNews,
		"featuredPublications": featured,
		"stats": gin.H{
			"publications": publicationsCount,
			"projects":     projectsCount,
			"teamMembers":  peopleCount,
		},
	})
}

func (s *Server) publicNews(c *gin.Context) {
	var items []NewsItem
	s.db.Where("status = ?", StatusPublished).
		Order("event_date DESC, sort_order ASC, id DESC").
		Find(&items)
	c.JSON(http.StatusOK, gin.H{"items": items})
}

func (s *Server) publicPeople(c *gin.Context) {
	var items []Person
	db := s.db.Where("status = ?", StatusPublished)
	category := c.Query("category")
	if category != "" {
		db = db.Where("category = ?", category)
	}
	db.Order("sort_order ASC, id ASC").Find(&items)
	c.JSON(http.StatusOK, gin.H{"items": items})
}

func (s *Server) publicUndergraduates(c *gin.Context) {
	var items []UndergraduateEducation
	s.db.Where("status = ?", StatusPublished).
		Order("sort_order ASC, id ASC").
		Find(&items)
	c.JSON(http.StatusOK, gin.H{"items": items})
}

func (s *Server) publicPublications(c *gin.Context) {
	var items []Publication
	s.db.Preload("Links").
		Where("status = ?", StatusPublished).
		Order("year DESC, FIELD(kind, 'journal', 'conference'), sort_order ASC, id ASC").
		Find(&items)
	c.JSON(http.StatusOK, gin.H{
		"items":  items,
		"groups": groupPublications(items),
	})
}

func (s *Server) publicPatents(c *gin.Context) {
	var items []Patent
	s.db.Where("status = ?", StatusPublished).
		Order("FIELD(category, 'granted', 'review', 'standard'), sort_order ASC, id ASC").
		Find(&items)
	c.JSON(http.StatusOK, gin.H{"items": items, "groups": groupPatents(items)})
}

func (s *Server) publicResearchProjects(c *gin.Context) {
	var items []ResearchProject
	s.db.Where("status = ?", StatusPublished).
		Order("sort_order ASC, id ASC").
		Find(&items)
	c.JSON(http.StatusOK, gin.H{"items": items})
}

func groupPublications(items []Publication) map[int]map[string][]Publication {
	groups := map[int]map[string][]Publication{}
	for _, item := range items {
		if _, exists := groups[item.Year]; !exists {
			groups[item.Year] = map[string][]Publication{"journal": {}, "conference": {}}
		}
		groups[item.Year][item.Kind] = append(groups[item.Year][item.Kind], item)
	}
	return groups
}

func groupPatents(items []Patent) map[string][]Patent {
	groups := map[string][]Patent{"granted": {}, "review": {}, "standard": {}}
	for _, item := range items {
		groups[item.Category] = append(groups[item.Category], item)
	}
	return groups
}

func sortedYears(groups map[int]map[string][]Publication) []int {
	years := make([]int, 0, len(groups))
	for year := range groups {
		years = append(years, year)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(years)))
	return years
}
