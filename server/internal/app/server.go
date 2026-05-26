package app

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Server struct {
	cfg Config
	db  *gorm.DB
	r   *gin.Engine
}

func NewServer(cfg Config) (*Server, error) {
	if err := os.MkdirAll(cfg.UploadDir, 0o755); err != nil {
		return nil, err
	}

	db, err := gorm.Open(mysql.Open(cfg.DatabaseDSN), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
	})
	if err != nil {
		return nil, err
	}

	server := &Server{cfg: cfg, db: db}
	if err := server.migrateAndSeed(); err != nil {
		return nil, err
	}

	server.r = server.routes()
	return server, nil
}

func (s *Server) Run() error {
	return s.r.Run(":" + s.cfg.Port)
}

func (s *Server) routes() *gin.Engine {
	r := gin.Default()
	r.MaxMultipartMemory = s.cfg.MaxUploadSize
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "http://127.0.0.1:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.Static("/uploads", s.cfg.UploadDir)
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	public := r.Group("/api/public")
	{
		public.GET("/home", s.publicHome)
		public.GET("/pages/:slug", s.publicSitePage)
		public.GET("/news", s.publicNews)
		public.GET("/people", s.publicPeople)
		public.GET("/undergraduates", s.publicUndergraduates)
		public.GET("/publications", s.publicPublications)
		public.GET("/patents", s.publicPatents)
		public.GET("/research-projects", s.publicResearchProjects)
	}

	admin := r.Group("/api/admin")
	{
		admin.POST("/auth/login", s.login)
		admin.GET("/me", s.authMiddleware(), s.me)

		protected := admin.Group("", s.authMiddleware())
		protected.GET("/summary", s.adminSummary)
		protected.GET("/maintenance/status", s.maintenanceStatus)
		protected.POST("/maintenance/backup", s.runBackup)
		protected.POST("/maintenance/git-sync", s.runGitSync)
		protected.GET("/trash", s.listTrash)
		protected.POST("/trash/:resource/:id/restore", s.restoreTrash)

		protected.GET("/pages", s.listSitePages)
		protected.POST("/pages", s.createSitePage)
		protected.GET("/pages/:slug", s.getSitePage)
		protected.PUT("/pages/:slug", s.updateSitePage)
		protected.DELETE("/pages/:slug", s.deleteSitePage)

		protected.GET("/news", s.listNews)
		protected.POST("/news", s.createNews)
		protected.PUT("/news/:id", s.updateNews)
		protected.DELETE("/news/:id", s.deleteNews)

		protected.GET("/people", s.listPeople)
		protected.POST("/people", s.createPerson)
		protected.PUT("/people/:id", s.updatePerson)
		protected.POST("/people/:id/place", s.placePerson)
		protected.DELETE("/people/:id", s.deletePerson)

		protected.GET("/undergraduates", s.listUndergraduates)
		protected.POST("/undergraduates", s.createUndergraduate)
		protected.PUT("/undergraduates/:id", s.updateUndergraduate)
		protected.POST("/undergraduates/:id/place", s.placeUndergraduate)
		protected.DELETE("/undergraduates/:id", s.deleteUndergraduate)

		protected.GET("/publications", s.listPublications)
		protected.POST("/publications", s.createPublication)
		protected.PUT("/publications/:id", s.updatePublication)
		protected.POST("/publications/:id/place", s.placePublication)
		protected.DELETE("/publications/:id", s.deletePublication)

		protected.GET("/patents", s.listPatents)
		protected.POST("/patents", s.createPatent)
		protected.PUT("/patents/:id", s.updatePatent)
		protected.POST("/patents/:id/place", s.placePatent)
		protected.DELETE("/patents/:id", s.deletePatent)

		protected.GET("/research-projects", s.listResearchProjects)
		protected.POST("/research-projects", s.createResearchProject)
		protected.PUT("/research-projects/:id", s.updateResearchProject)
		protected.POST("/research-projects/:id/place", s.placeResearchProject)
		protected.DELETE("/research-projects/:id", s.deleteResearchProject)

		protected.GET("/media", s.listMedia)
		protected.POST("/media", s.uploadMedia)
		protected.POST("/media/import-public", s.importPublicMedia)
		protected.PUT("/media/:id", s.updateMedia)
		protected.DELETE("/media/:id", s.deleteMedia)
	}

	return r
}

func (s *Server) migrateAndSeed() error {
	if err := s.db.AutoMigrate(
		&Admin{},
		&SitePage{},
		&MediaAsset{},
		&NewsItem{},
		&Person{},
		&UndergraduateEducation{},
		&Publication{},
		&PublicationLink{},
		&Patent{},
		&ResearchProject{},
	); err != nil {
		return err
	}

	if err := s.ensureAdmin(); err != nil {
		return err
	}
	if err := s.seedContent(); err != nil {
		return err
	}
	return s.backfillMediaDisplayNames()
}

func bindJSON[T any](c *gin.Context) (T, bool) {
	var payload T
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return payload, false
	}
	return payload, true
}

func parseID(c *gin.Context) (uint, bool) {
	var id uint
	if _, err := fmt.Sscanf(c.Param("id"), "%d", &id); err != nil || id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return 0, false
	}
	return id, true
}

func applyListQuery(c *gin.Context, db *gorm.DB, searchable ...string) (*gorm.DB, int, int) {
	page := queryInt(c, "page", 1)
	pageSize := queryInt(c, "pageSize", 20)
	if pageSize > 100 {
		pageSize = 100
	}
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 20
	}

	status := strings.TrimSpace(c.Query("status"))
	if status != "" {
		db = db.Where("status = ?", status)
	}

	q := strings.TrimSpace(c.Query("q"))
	if q != "" && len(searchable) > 0 {
		like := "%" + q + "%"
		conditions := make([]string, 0, len(searchable))
		args := make([]any, 0, len(searchable))
		for _, field := range searchable {
			conditions = append(conditions, field+" LIKE ?")
			args = append(args, like)
		}
		db = db.Where(strings.Join(conditions, " OR "), args...)
	}

	return db, page, pageSize
}

func paged(c *gin.Context, db *gorm.DB, out any, page int, pageSize int) {
	var total int64
	db.Count(&total)
	db.Offset((page - 1) * pageSize).Limit(pageSize).Find(out)
	c.JSON(http.StatusOK, gin.H{
		"items":    out,
		"total":    total,
		"page":     page,
		"pageSize": pageSize,
	})
}

func queryInt(c *gin.Context, key string, fallback int) int {
	value := c.Query(key)
	if value == "" {
		return fallback
	}
	var parsed int
	if _, err := fmt.Sscanf(value, "%d", &parsed); err != nil {
		return fallback
	}
	return parsed
}

func notFoundOrError(c *gin.Context, err error) {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
}
