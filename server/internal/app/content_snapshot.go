package app

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"time"

	"gorm.io/gorm"
)

const minSnapshotBytes = 2000

type contentSnapshot struct {
	SitePages        []snapshotSitePage        `json:"sitePages"`
	News             []snapshotNews            `json:"news"`
	Media            []snapshotMedia           `json:"media"`
	People           []snapshotPerson          `json:"people"`
	Patents          []snapshotPatent          `json:"patents"`
	Publications     []snapshotPublication     `json:"publications"`
	Undergraduates   []snapshotUndergraduate   `json:"undergraduates"`
	PublicationLinks []snapshotPublicationLink `json:"publicationLinks"`
	ResearchProjects []snapshotResearchProject `json:"researchProjects"`
}

type snapshotSitePage struct {
	ID          uint    `json:"id"`
	Slug        string  `json:"slug"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Content     JSONMap `json:"content"`
	Status      string  `json:"status"`
	CreatedAt   string  `json:"createdAt"`
	DeletedAt   *string `json:"deletedAt"`
	SortOrder   int     `json:"sortOrder"`
	UpdatedAt   string  `json:"updatedAt"`
}

type snapshotNews struct {
	ID        uint    `json:"id"`
	Date      string  `json:"date"`
	Type      string  `json:"type"`
	Color     string  `json:"color"`
	Title     string  `json:"title"`
	Status    string  `json:"status"`
	Content   string  `json:"content"`
	Excerpt   string  `json:"excerpt"`
	CreatedAt string  `json:"createdAt"`
	DeletedAt *string `json:"deletedAt"`
	SortOrder int     `json:"sortOrder"`
	TypeLabel string  `json:"typeLabel"`
	UpdatedAt string  `json:"updatedAt"`
}

type snapshotMedia struct {
	ID           uint    `json:"id"`
	URL          string  `json:"url"`
	Kind         string  `json:"kind"`
	Size         int64   `json:"size"`
	FileName     string  `json:"fileName"`
	MimeType     string  `json:"mimeType"`
	CreatedAt    string  `json:"createdAt"`
	DeletedAt    *string `json:"deletedAt"`
	DisplayName  string  `json:"displayName"`
	UpdatedAt    string  `json:"updatedAt"`
	OriginalName string  `json:"originalName"`
}

type snapshotPerson struct {
	ID             uint    `json:"id"`
	Name           string  `json:"name"`
	Status         string  `json:"status"`
	Category       string  `json:"category"`
	Research       string  `json:"research"`
	AvatarURL      string  `json:"avatarUrl"`
	AvatarObjectX  int     `json:"avatarObjectX"`
	AvatarObjectY  int     `json:"avatarObjectY"`
	AvatarScale    int     `json:"avatarScale"`
	CreatedAt      string  `json:"createdAt"`
	DeletedAt      *string `json:"deletedAt"`
	SortOrder      int     `json:"sortOrder"`
	UpdatedAt      string  `json:"updatedAt"`
	GraduationDate string  `json:"graduationDate"`
}

type snapshotPatent struct {
	ID        uint    `json:"id"`
	Date      string  `json:"date"`
	Title     string  `json:"title"`
	Number    string  `json:"number"`
	Status    string  `json:"status"`
	Authors   string  `json:"authors"`
	Country   string  `json:"country"`
	Category  string  `json:"category"`
	CreatedAt string  `json:"createdAt"`
	DeletedAt *string `json:"deletedAt"`
	SortOrder int     `json:"sortOrder"`
	UpdatedAt string  `json:"updatedAt"`
}

type snapshotPublication struct {
	ID        uint    `json:"id"`
	Kind      string  `json:"kind"`
	Year      int     `json:"year"`
	Image     string  `json:"image"`
	Title     string  `json:"title"`
	Venue     string  `json:"venue"`
	Status    string  `json:"status"`
	Authors   string  `json:"authors"`
	Featured  int     `json:"featured"`
	CreatedAt string  `json:"createdAt"`
	DeletedAt *string `json:"deletedAt"`
	SortOrder int     `json:"sortOrder"`
	UpdatedAt string  `json:"updatedAt"`
}

type snapshotUndergraduate struct {
	ID           uint     `json:"id"`
	Name         string   `json:"name"`
	Grade        string   `json:"grade"`
	Major        string   `json:"major"`
	Status       string   `json:"status"`
	CreatedAt    string   `json:"createdAt"`
	DeletedAt    *string  `json:"deletedAt"`
	Direction    string   `json:"direction"`
	SortOrder    int      `json:"sortOrder"`
	UpdatedAt    string   `json:"updatedAt"`
	Achievements []string `json:"achievements"`
}

type snapshotPublicationLink struct {
	ID            uint   `json:"id"`
	URL           string `json:"url"`
	Type          string `json:"type"`
	Label         string `json:"label"`
	CreatedAt     string `json:"createdAt"`
	RouteName     string `json:"routeName"`
	SortOrder     int    `json:"sortOrder"`
	UpdatedAt     string `json:"updatedAt"`
	PublicationID uint   `json:"publicationId"`
}

type snapshotResearchProject struct {
	ID            uint    `json:"id"`
	Fund          string  `json:"fund"`
	Role          string  `json:"role"`
	Title         string  `json:"title"`
	Amount        string  `json:"amount"`
	Number        string  `json:"number"`
	Period        string  `json:"period"`
	Status        string  `json:"status"`
	CreatedAt     string  `json:"createdAt"`
	DeletedAt     *string `json:"deletedAt"`
	SortOrder     int     `json:"sortOrder"`
	UpdatedAt     string  `json:"updatedAt"`
	ProjectStatus string  `json:"projectStatus"`
}

func (s *Server) exportContentSnapshot(root string) (contentSnapshot, string, error) {
	data, err := s.buildContentSnapshot()
	if err != nil {
		return data, "", err
	}
	path, err := writeContentSnapshot(root, data)
	return data, path, err
}

func (s *Server) buildContentSnapshot() (contentSnapshot, error) {
	data := contentSnapshot{}

	var pages []SitePage
	if err := s.db.Unscoped().Order("sort_order ASC, id ASC").Find(&pages).Error; err != nil {
		return data, err
	}
	data.SitePages = make([]snapshotSitePage, 0, len(pages))
	for _, item := range pages {
		times := snapshotTimes(item.CreatedAt, item.UpdatedAt, item.DeletedAt)
		data.SitePages = append(data.SitePages, snapshotSitePage{
			ID: item.ID, Slug: item.Slug, Title: item.Title, Description: item.Description,
			Content: item.Content, Status: item.Status, CreatedAt: times.createdAt,
			DeletedAt: times.deletedAt, SortOrder: item.SortOrder, UpdatedAt: times.updatedAt,
		})
	}

	var news []NewsItem
	if err := s.db.Unscoped().Order("sort_order ASC, event_date DESC, id ASC").Find(&news).Error; err != nil {
		return data, err
	}
	data.News = make([]snapshotNews, 0, len(news))
	for _, item := range news {
		times := snapshotTimes(item.CreatedAt, item.UpdatedAt, item.DeletedAt)
		data.News = append(data.News, snapshotNews{
			ID: item.ID, Date: item.EventDate, Type: item.Type, Color: item.Color,
			Title: item.Title, Status: item.Status, Content: item.Content, Excerpt: item.Excerpt,
			CreatedAt: times.createdAt, DeletedAt: times.deletedAt, SortOrder: item.SortOrder,
			TypeLabel: item.TypeLabel, UpdatedAt: times.updatedAt,
		})
	}

	var media []MediaAsset
	if err := s.db.Unscoped().Order("id ASC").Find(&media).Error; err != nil {
		return data, err
	}
	data.Media = make([]snapshotMedia, 0, len(media))
	for _, item := range media {
		times := snapshotTimes(item.CreatedAt, item.UpdatedAt, item.DeletedAt)
		data.Media = append(data.Media, snapshotMedia{
			ID: item.ID, URL: item.URL, Kind: item.Kind, Size: item.Size, FileName: item.FileName,
			MimeType: item.MimeType, CreatedAt: times.createdAt, DeletedAt: times.deletedAt,
			DisplayName: item.DisplayName, UpdatedAt: times.updatedAt, OriginalName: item.OriginalName,
		})
	}

	var people []Person
	if err := s.db.Unscoped().Order("sort_order ASC, category ASC, id ASC").Find(&people).Error; err != nil {
		return data, err
	}
	data.People = make([]snapshotPerson, 0, len(people))
	for _, item := range people {
		times := snapshotTimes(item.CreatedAt, item.UpdatedAt, item.DeletedAt)
		data.People = append(data.People, snapshotPerson{
			ID: item.ID, Name: item.Name, Status: item.Status, Category: item.Category,
			Research: item.Research, AvatarURL: item.AvatarURL, CreatedAt: times.createdAt,
			AvatarObjectX: item.AvatarObjectX, AvatarObjectY: item.AvatarObjectY,
			AvatarScale: item.AvatarScale, DeletedAt: times.deletedAt,
			SortOrder: item.SortOrder, UpdatedAt: times.updatedAt, GraduationDate: item.GraduationDate,
		})
	}

	var patents []Patent
	if err := s.db.Unscoped().Order("sort_order ASC, date DESC, id ASC").Find(&patents).Error; err != nil {
		return data, err
	}
	data.Patents = make([]snapshotPatent, 0, len(patents))
	for _, item := range patents {
		times := snapshotTimes(item.CreatedAt, item.UpdatedAt, item.DeletedAt)
		data.Patents = append(data.Patents, snapshotPatent{
			ID: item.ID, Date: item.Date, Title: item.Title, Number: item.Number,
			Status: item.Status, Authors: item.Authors, Country: item.Country,
			Category: item.Category, CreatedAt: times.createdAt, DeletedAt: times.deletedAt,
			SortOrder: item.SortOrder, UpdatedAt: times.updatedAt,
		})
	}

	var publications []Publication
	if err := s.db.Unscoped().Order("sort_order ASC, year DESC, id ASC").Find(&publications).Error; err != nil {
		return data, err
	}
	data.Publications = make([]snapshotPublication, 0, len(publications))
	for _, item := range publications {
		times := snapshotTimes(item.CreatedAt, item.UpdatedAt, item.DeletedAt)
		data.Publications = append(data.Publications, snapshotPublication{
			ID: item.ID, Kind: item.Kind, Year: item.Year, Image: item.ImageURL,
			Title: item.Title, Venue: item.Venue, Status: item.Status, Authors: item.Authors,
			Featured: boolInt(item.Featured), CreatedAt: times.createdAt, DeletedAt: times.deletedAt,
			SortOrder: item.SortOrder, UpdatedAt: times.updatedAt,
		})
	}

	var undergraduates []UndergraduateEducation
	if err := s.db.Unscoped().Order("sort_order ASC, grade DESC, id ASC").Find(&undergraduates).Error; err != nil {
		return data, err
	}
	data.Undergraduates = make([]snapshotUndergraduate, 0, len(undergraduates))
	for _, item := range undergraduates {
		times := snapshotTimes(item.CreatedAt, item.UpdatedAt, item.DeletedAt)
		data.Undergraduates = append(data.Undergraduates, snapshotUndergraduate{
			ID: item.ID, Name: item.Name, Grade: item.Grade, Major: item.Major,
			Status: item.Status, CreatedAt: times.createdAt, DeletedAt: times.deletedAt,
			Direction: item.Direction, SortOrder: item.SortOrder, UpdatedAt: times.updatedAt,
			Achievements: []string(item.Achievements),
		})
	}

	var links []PublicationLink
	if err := s.db.Unscoped().Order("publication_id ASC, sort_order ASC, id ASC").Find(&links).Error; err != nil {
		return data, err
	}
	data.PublicationLinks = make([]snapshotPublicationLink, 0, len(links))
	for _, item := range links {
		data.PublicationLinks = append(data.PublicationLinks, snapshotPublicationLink{
			ID: item.ID, URL: item.URL, Type: item.Type, Label: item.Label,
			CreatedAt: formatSnapshotTime(item.CreatedAt), RouteName: item.RouteName,
			SortOrder: item.SortOrder, UpdatedAt: formatSnapshotTime(item.UpdatedAt),
			PublicationID: item.PublicationID,
		})
	}

	var projects []ResearchProject
	if err := s.db.Unscoped().Order("sort_order ASC, id ASC").Find(&projects).Error; err != nil {
		return data, err
	}
	data.ResearchProjects = make([]snapshotResearchProject, 0, len(projects))
	for _, item := range projects {
		times := snapshotTimes(item.CreatedAt, item.UpdatedAt, item.DeletedAt)
		data.ResearchProjects = append(data.ResearchProjects, snapshotResearchProject{
			ID: item.ID, Fund: item.Fund, Role: item.Role, Title: item.Title,
			Amount: item.Amount, Number: item.Number, Period: item.Period, Status: item.Status,
			CreatedAt: times.createdAt, DeletedAt: times.deletedAt, SortOrder: item.SortOrder,
			UpdatedAt: times.updatedAt, ProjectStatus: item.ProjectStatus,
		})
	}

	return data, nil
}

type snapshotTimeValues struct {
	createdAt string
	updatedAt string
	deletedAt *string
}

func snapshotTimes(createdAt time.Time, updatedAt time.Time, deletedAt gorm.DeletedAt) snapshotTimeValues {
	var deleted *string
	if deletedAt.Valid {
		value := formatSnapshotTime(deletedAt.Time)
		deleted = &value
	}
	return snapshotTimeValues{
		createdAt: formatSnapshotTime(createdAt),
		updatedAt: formatSnapshotTime(updatedAt),
		deletedAt: deleted,
	}
}

func formatSnapshotTime(value time.Time) string {
	if value.IsZero() {
		return ""
	}
	return value.Format("2006-01-02T15:04:05")
}

func boolInt(value bool) int {
	if value {
		return 1
	}
	return 0
}

func writeContentSnapshot(root string, data contentSnapshot) (string, error) {
	contentDir := filepath.Join(root, "storage", "content")
	if err := os.MkdirAll(filepath.Join(contentDir, "checkpoints", "weekly"), 0o755); err != nil {
		return "", err
	}
	if err := os.MkdirAll(filepath.Join(contentDir, "checkpoints", "monthly"), 0o755); err != nil {
		return "", err
	}

	snapshotPath := filepath.Join(contentDir, "content.json")
	tmpPath := snapshotPath + ".tmp"
	var buffer bytes.Buffer
	encoder := json.NewEncoder(&buffer)
	encoder.SetEscapeHTML(false)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(data); err != nil {
		return "", err
	}
	raw := buffer.Bytes()
	if len(raw) < minSnapshotBytes {
		return "", errSnapshotTooSmall{}
	}
	if err := rejectLargeSnapshotShrink(snapshotPath, len(raw)); err != nil {
		return "", err
	}
	if err := os.WriteFile(tmpPath, raw, 0o644); err != nil {
		return "", err
	}
	if err := writeContentCheckpoints(contentDir, raw); err != nil {
		_ = os.Remove(tmpPath)
		return "", err
	}
	if err := os.Rename(tmpPath, snapshotPath); err != nil {
		_ = os.Remove(tmpPath)
		return "", err
	}
	return snapshotPath, nil
}

func writeContentCheckpoints(contentDir string, raw []byte) error {
	now := time.Now()
	weekly := filepath.Join(contentDir, "checkpoints", "weekly", now.Format("2006")+"-W"+weekSunday(now)+".json")
	monthly := filepath.Join(contentDir, "checkpoints", "monthly", now.Format("2006-01")+".json")
	if err := writeCheckpointIfMissing(weekly, raw); err != nil {
		return err
	}
	return writeCheckpointIfMissing(monthly, raw)
}

func writeCheckpointIfMissing(path string, raw []byte) error {
	info, err := os.Stat(path)
	if err == nil && !info.IsDir() && info.Size() >= minSnapshotBytes {
		return nil
	}
	if err != nil && !os.IsNotExist(err) {
		return err
	}
	return os.WriteFile(path, raw, 0o644)
}

func rejectLargeSnapshotShrink(snapshotPath string, nextSize int) error {
	info, err := os.Stat(snapshotPath)
	if err != nil || info.IsDir() || info.Size() <= 0 {
		return nil
	}
	minimum := int(info.Size()) * 30 / 100
	if minimum > minSnapshotBytes && nextSize < minimum {
		return errSnapshotShrank{previous: info.Size(), next: int64(nextSize)}
	}
	return nil
}

type errSnapshotTooSmall struct{}

func (errSnapshotTooSmall) Error() string {
	return "generated content snapshot is unexpectedly small"
}

type errSnapshotShrank struct {
	previous int64
	next     int64
}

func (errSnapshotShrank) Error() string {
	return "generated content snapshot is much smaller than the previous snapshot"
}

func weekSunday(value time.Time) string {
	_, week := value.AddDate(0, 0, int(time.Sunday-value.Weekday())).ISOWeek()
	return twoDigits(week)
}

func twoDigits(value int) string {
	if value < 10 {
		return "0" + string(rune('0'+value))
	}
	return string(rune('0'+value/10)) + string(rune('0'+value%10))
}
