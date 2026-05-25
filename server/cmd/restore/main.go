package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"byml/server/internal/app"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type snapshotBool bool

type snapshot struct {
	News             []snapshotNews            `json:"news"`
	Media            []snapshotMedia           `json:"media"`
	People           []snapshotPerson          `json:"people"`
	Undergraduates   []snapshotUndergraduate   `json:"undergraduates"`
	Publications     []snapshotPublication     `json:"publications"`
	PublicationLinks []snapshotPublicationLink `json:"publicationLinks"`
	Patents          []snapshotPatent          `json:"patents"`
	ResearchProjects []snapshotResearchProject `json:"researchProjects"`
}

type snapshotTimestamps struct {
	CreatedAt string  `json:"createdAt"`
	UpdatedAt string  `json:"updatedAt"`
	DeletedAt *string `json:"deletedAt"`
}

type snapshotNews struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Excerpt   string `json:"excerpt"`
	Type      string `json:"type"`
	TypeLabel string `json:"typeLabel"`
	Date      string `json:"date"`
	Color     string `json:"color"`
	Status    string `json:"status"`
	SortOrder int    `json:"sortOrder"`
	snapshotTimestamps
}

type snapshotMedia struct {
	ID           uint   `json:"id"`
	FileName     string `json:"fileName"`
	OriginalName string `json:"originalName"`
	DisplayName  string `json:"displayName"`
	URL          string `json:"url"`
	MimeType     string `json:"mimeType"`
	Size         int64  `json:"size"`
	Kind         string `json:"kind"`
	snapshotTimestamps
}

type snapshotPerson struct {
	ID             uint   `json:"id"`
	Name           string `json:"name"`
	AvatarURL      string `json:"avatarUrl"`
	Category       string `json:"category"`
	Research       string `json:"research"`
	GraduationDate string `json:"graduationDate"`
	Status         string `json:"status"`
	SortOrder      int    `json:"sortOrder"`
	snapshotTimestamps
}

type snapshotUndergraduate struct {
	ID           uint     `json:"id"`
	Name         string   `json:"name"`
	Grade        string   `json:"grade"`
	Major        string   `json:"major"`
	Direction    string   `json:"direction"`
	Achievements []string `json:"achievements"`
	Status       string   `json:"status"`
	SortOrder    int      `json:"sortOrder"`
	snapshotTimestamps
}

type snapshotPublication struct {
	ID        uint         `json:"id"`
	Image     string       `json:"image"`
	Title     string       `json:"title"`
	Authors   string       `json:"authors"`
	Venue     string       `json:"venue"`
	Year      int          `json:"year"`
	Kind      string       `json:"kind"`
	Status    string       `json:"status"`
	Featured  snapshotBool `json:"featured"`
	SortOrder int          `json:"sortOrder"`
	snapshotTimestamps
}

type snapshotPublicationLink struct {
	ID            uint   `json:"id"`
	PublicationID uint   `json:"publicationId"`
	Type          string `json:"type"`
	Label         string `json:"label"`
	URL           string `json:"url"`
	RouteName     string `json:"routeName"`
	SortOrder     int    `json:"sortOrder"`
	CreatedAt     string `json:"createdAt"`
	UpdatedAt     string `json:"updatedAt"`
}

type snapshotPatent struct {
	ID        uint   `json:"id"`
	Authors   string `json:"authors"`
	Title     string `json:"title"`
	Date      string `json:"date"`
	Country   string `json:"country"`
	Number    string `json:"number"`
	Category  string `json:"category"`
	Status    string `json:"status"`
	SortOrder int    `json:"sortOrder"`
	snapshotTimestamps
}

type snapshotResearchProject struct {
	ID            uint   `json:"id"`
	Title         string `json:"title"`
	Fund          string `json:"fund"`
	Number        string `json:"number"`
	Period        string `json:"period"`
	Amount        string `json:"amount"`
	ProjectStatus string `json:"projectStatus"`
	Role          string `json:"role"`
	Status        string `json:"status"`
	SortOrder     int    `json:"sortOrder"`
	snapshotTimestamps
}

func main() {
	file := flag.String("file", "../storage/content/content.json", "content snapshot JSON file")
	dryRun := flag.Bool("dry-run", false, "validate the snapshot without changing the database")
	flag.Parse()

	cfg := app.LoadConfig()
	data, err := readSnapshot(*file)
	if err != nil {
		log.Fatalf("failed to read snapshot: %v", err)
	}

	if *dryRun {
		printSummary(data, "validated")
		return
	}

	db, err := gorm.Open(mysql.Open(cfg.DatabaseDSN), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
	})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	if err := restoreSnapshot(db, cfg, data); err != nil {
		log.Fatalf("failed to restore snapshot: %v", err)
	}
	printSummary(data, "restored")
}

func readSnapshot(file string) (snapshot, error) {
	var data snapshot
	raw, err := os.ReadFile(file)
	if err != nil {
		return data, err
	}
	if err := json.Unmarshal(raw, &data); err != nil {
		return data, err
	}
	return data, nil
}

func (b *snapshotBool) UnmarshalJSON(raw []byte) error {
	var value any
	if err := json.Unmarshal(raw, &value); err != nil {
		return err
	}
	switch typed := value.(type) {
	case bool:
		*b = snapshotBool(typed)
	case float64:
		*b = snapshotBool(typed != 0)
	case string:
		normalized := strings.ToLower(strings.TrimSpace(typed))
		*b = normalized == "true" || normalized == "1" || normalized == "yes"
	case nil:
		*b = false
	default:
		return fmt.Errorf("unsupported bool value %T", value)
	}
	return nil
}

func restoreSnapshot(db *gorm.DB, cfg app.Config, data snapshot) error {
	if err := db.AutoMigrate(
		&app.MediaAsset{},
		&app.NewsItem{},
		&app.Person{},
		&app.UndergraduateEducation{},
		&app.Publication{},
		&app.PublicationLink{},
		&app.Patent{},
		&app.ResearchProject{},
	); err != nil {
		return err
	}

	return db.Transaction(func(tx *gorm.DB) error {
		if err := clearContentTables(tx); err != nil {
			return err
		}
		if err := insertAll(tx, mapNews(data.News)); err != nil {
			return err
		}
		if err := insertAll(tx, mapMedia(cfg, data.Media)); err != nil {
			return err
		}
		if err := insertAll(tx, mapPeople(data.People)); err != nil {
			return err
		}
		if err := insertAll(tx, mapUndergraduates(data.Undergraduates)); err != nil {
			return err
		}
		if err := insertAll(tx, mapPublications(data.Publications)); err != nil {
			return err
		}
		if err := insertAll(tx, mapPublicationLinks(data.PublicationLinks)); err != nil {
			return err
		}
		if err := insertAll(tx, mapPatents(data.Patents)); err != nil {
			return err
		}
		return insertAll(tx, mapResearchProjects(data.ResearchProjects))
	})
}

func clearContentTables(tx *gorm.DB) error {
	tables := []string{
		"publication_links",
		"news_items",
		"media_assets",
		"people",
		"undergraduate_educations",
		"publications",
		"patents",
		"research_projects",
	}
	if err := tx.Exec("SET FOREIGN_KEY_CHECKS = 0").Error; err != nil {
		return err
	}
	for _, table := range tables {
		if err := tx.Exec("DELETE FROM " + table).Error; err != nil {
			return err
		}
		if err := tx.Exec("ALTER TABLE " + table + " AUTO_INCREMENT = 1").Error; err != nil {
			return err
		}
	}
	return tx.Exec("SET FOREIGN_KEY_CHECKS = 1").Error
}

func insertAll[T any](tx *gorm.DB, items []T) error {
	if len(items) == 0 {
		return nil
	}
	return tx.Create(&items).Error
}

func mapNews(items []snapshotNews) []app.NewsItem {
	out := make([]app.NewsItem, 0, len(items))
	for _, item := range items {
		createdAt, updatedAt, deletedAt := parseTimestamps(item.snapshotTimestamps)
		out = append(out, app.NewsItem{
			ID:        item.ID,
			Title:     item.Title,
			Content:   item.Content,
			Excerpt:   item.Excerpt,
			Type:      item.Type,
			TypeLabel: item.TypeLabel,
			EventDate: item.Date,
			Color:     item.Color,
			Status:    normalizeSnapshotStatus(item.Status),
			SortOrder: item.SortOrder,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
			DeletedAt: deletedAt,
		})
	}
	return out
}

func mapMedia(cfg app.Config, items []snapshotMedia) []app.MediaAsset {
	out := make([]app.MediaAsset, 0, len(items))
	for _, item := range items {
		createdAt, updatedAt, deletedAt := parseTimestamps(item.snapshotTimestamps)
		out = append(out, app.MediaAsset{
			ID:           item.ID,
			FileName:     item.FileName,
			OriginalName: item.OriginalName,
			DisplayName:  item.DisplayName,
			URL:          item.URL,
			Path:         inferMediaPath(cfg, item.URL),
			MimeType:     item.MimeType,
			Size:         item.Size,
			Kind:         item.Kind,
			CreatedAt:    createdAt,
			UpdatedAt:    updatedAt,
			DeletedAt:    deletedAt,
		})
	}
	return out
}

func mapPeople(items []snapshotPerson) []app.Person {
	out := make([]app.Person, 0, len(items))
	for _, item := range items {
		createdAt, updatedAt, deletedAt := parseTimestamps(item.snapshotTimestamps)
		out = append(out, app.Person{
			ID:             item.ID,
			Name:           item.Name,
			AvatarURL:      item.AvatarURL,
			Category:       item.Category,
			Research:       item.Research,
			GraduationDate: item.GraduationDate,
			Status:         normalizeSnapshotStatus(item.Status),
			SortOrder:      item.SortOrder,
			CreatedAt:      createdAt,
			UpdatedAt:      updatedAt,
			DeletedAt:      deletedAt,
		})
	}
	return out
}

func mapUndergraduates(items []snapshotUndergraduate) []app.UndergraduateEducation {
	out := make([]app.UndergraduateEducation, 0, len(items))
	for _, item := range items {
		createdAt, updatedAt, deletedAt := parseTimestamps(item.snapshotTimestamps)
		out = append(out, app.UndergraduateEducation{
			ID:           item.ID,
			Name:         item.Name,
			Grade:        item.Grade,
			Major:        item.Major,
			Direction:    item.Direction,
			Achievements: app.StringList(item.Achievements),
			Status:       normalizeSnapshotStatus(item.Status),
			SortOrder:    item.SortOrder,
			CreatedAt:    createdAt,
			UpdatedAt:    updatedAt,
			DeletedAt:    deletedAt,
		})
	}
	return out
}

func mapPublications(items []snapshotPublication) []app.Publication {
	out := make([]app.Publication, 0, len(items))
	for _, item := range items {
		createdAt, updatedAt, deletedAt := parseTimestamps(item.snapshotTimestamps)
		out = append(out, app.Publication{
			ID:        item.ID,
			ImageURL:  item.Image,
			Title:     item.Title,
			Authors:   item.Authors,
			Venue:     item.Venue,
			Year:      item.Year,
			Kind:      item.Kind,
			Status:    normalizeSnapshotStatus(item.Status),
			Featured:  bool(item.Featured),
			SortOrder: item.SortOrder,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
			DeletedAt: deletedAt,
		})
	}
	return out
}

func mapPublicationLinks(items []snapshotPublicationLink) []app.PublicationLink {
	out := make([]app.PublicationLink, 0, len(items))
	for _, item := range items {
		out = append(out, app.PublicationLink{
			ID:            item.ID,
			PublicationID: item.PublicationID,
			Type:          item.Type,
			Label:         item.Label,
			URL:           item.URL,
			RouteName:     item.RouteName,
			SortOrder:     item.SortOrder,
			CreatedAt:     parseSnapshotTime(item.CreatedAt),
			UpdatedAt:     parseSnapshotTime(item.UpdatedAt),
		})
	}
	return out
}

func mapPatents(items []snapshotPatent) []app.Patent {
	out := make([]app.Patent, 0, len(items))
	for _, item := range items {
		createdAt, updatedAt, deletedAt := parseTimestamps(item.snapshotTimestamps)
		out = append(out, app.Patent{
			ID:        item.ID,
			Authors:   item.Authors,
			Title:     item.Title,
			Date:      item.Date,
			Country:   item.Country,
			Number:    item.Number,
			Category:  item.Category,
			Status:    normalizeSnapshotStatus(item.Status),
			SortOrder: item.SortOrder,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
			DeletedAt: deletedAt,
		})
	}
	return out
}

func mapResearchProjects(items []snapshotResearchProject) []app.ResearchProject {
	out := make([]app.ResearchProject, 0, len(items))
	for _, item := range items {
		createdAt, updatedAt, deletedAt := parseTimestamps(item.snapshotTimestamps)
		out = append(out, app.ResearchProject{
			ID:            item.ID,
			Title:         item.Title,
			Fund:          item.Fund,
			Number:        item.Number,
			Period:        item.Period,
			Amount:        item.Amount,
			ProjectStatus: item.ProjectStatus,
			Role:          item.Role,
			Status:        normalizeSnapshotStatus(item.Status),
			SortOrder:     item.SortOrder,
			CreatedAt:     createdAt,
			UpdatedAt:     updatedAt,
			DeletedAt:     deletedAt,
		})
	}
	return out
}

func parseTimestamps(item snapshotTimestamps) (time.Time, time.Time, gorm.DeletedAt) {
	return parseSnapshotTime(item.CreatedAt), parseSnapshotTime(item.UpdatedAt), parseDeletedAt(item.DeletedAt)
}

func parseDeletedAt(value *string) gorm.DeletedAt {
	if value == nil || strings.TrimSpace(*value) == "" {
		return gorm.DeletedAt{}
	}
	return gorm.DeletedAt{Time: parseSnapshotTime(*value), Valid: true}
}

func parseSnapshotTime(value string) time.Time {
	value = strings.TrimSpace(value)
	if value == "" {
		return time.Time{}
	}
	layouts := []string{
		time.RFC3339,
		"2006-01-02T15:04:05",
		"2006-01-02 15:04:05",
	}
	for _, layout := range layouts {
		parsed, err := time.ParseInLocation(layout, value, time.Local)
		if err == nil {
			return parsed
		}
	}
	log.Fatalf("unsupported snapshot timestamp %q", value)
	return time.Time{}
}

func normalizeSnapshotStatus(status string) string {
	if strings.TrimSpace(status) == app.StatusDraft {
		return app.StatusDraft
	}
	return app.StatusPublished
}

func inferMediaPath(cfg app.Config, rawURL string) string {
	pathValue := rawURL
	if parsed, err := url.Parse(rawURL); err == nil && parsed.Path != "" {
		pathValue = parsed.Path
	}
	if strings.HasPrefix(pathValue, "/uploads/") {
		return filepath.Join(cfg.UploadDir, strings.TrimPrefix(pathValue, "/uploads/"))
	}
	return filepath.Join(cfg.PublicDir, strings.TrimPrefix(pathValue, "/"))
}

func printSummary(data snapshot, action string) {
	fmt.Printf("%s content snapshot: news=%d people=%d undergraduates=%d publications=%d publicationLinks=%d patents=%d researchProjects=%d media=%d\n",
		action,
		len(data.News),
		len(data.People),
		len(data.Undergraduates),
		len(data.Publications),
		len(data.PublicationLinks),
		len(data.Patents),
		len(data.ResearchProjects),
		len(data.Media),
	)
}
