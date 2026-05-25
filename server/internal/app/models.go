package app

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"gorm.io/gorm"
)

const (
	StatusDraft     = "draft"
	StatusPublished = "published"
)

type StringList []string

func (s StringList) Value() (driver.Value, error) {
	if s == nil {
		return "[]", nil
	}
	data, err := json.Marshal(s)
	return string(data), err
}

func (s *StringList) Scan(value any) error {
	if value == nil {
		*s = StringList{}
		return nil
	}
	var data []byte
	switch typed := value.(type) {
	case []byte:
		data = typed
	case string:
		data = []byte(typed)
	default:
		return fmt.Errorf("unsupported StringList value %T", value)
	}
	return json.Unmarshal(data, s)
}

type Admin struct {
	ID           uint           `json:"id" gorm:"primaryKey"`
	Username     string         `json:"username" gorm:"uniqueIndex;size:120;not null"`
	PasswordHash string         `json:"-" gorm:"size:255;not null"`
	CreatedAt    time.Time      `json:"createdAt"`
	UpdatedAt    time.Time      `json:"updatedAt"`
	DeletedAt    gorm.DeletedAt `json:"-" gorm:"index"`
}

type MediaAsset struct {
	ID           uint           `json:"id" gorm:"primaryKey"`
	FileName     string         `json:"fileName" gorm:"size:255;not null"`
	OriginalName string         `json:"originalName" gorm:"size:255;not null"`
	URL          string         `json:"url" gorm:"size:500;not null"`
	Path         string         `json:"-" gorm:"size:500;not null"`
	MimeType     string         `json:"mimeType" gorm:"size:160;not null"`
	Size         int64          `json:"size"`
	Kind         string         `json:"kind" gorm:"size:40;index"`
	CreatedAt    time.Time      `json:"createdAt"`
	UpdatedAt    time.Time      `json:"updatedAt"`
	DeletedAt    gorm.DeletedAt `json:"-" gorm:"index"`
}

type NewsItem struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Title     string         `json:"title" gorm:"size:255;not null;uniqueIndex"`
	Content   string         `json:"content" gorm:"type:text;not null"`
	Excerpt   string         `json:"excerpt" gorm:"size:500"`
	Type      string         `json:"type" gorm:"size:40;index"`
	TypeLabel string         `json:"typeLabel" gorm:"size:80"`
	EventDate string         `json:"date" gorm:"size:20;index"`
	Color     string         `json:"color" gorm:"size:40"`
	Status    string         `json:"status" gorm:"size:20;index;default:published"`
	SortOrder int            `json:"sortOrder" gorm:"index"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type Person struct {
	ID             uint           `json:"id" gorm:"primaryKey"`
	Name           string         `json:"name" gorm:"size:180;not null;uniqueIndex"`
	AvatarURL      string         `json:"avatarUrl" gorm:"size:500"`
	Category       string         `json:"category" gorm:"size:60;index"`
	Research       string         `json:"research" gorm:"size:255"`
	GraduationDate string         `json:"graduationDate" gorm:"size:80"`
	Status         string         `json:"status" gorm:"size:20;index;default:published"`
	SortOrder      int            `json:"sortOrder" gorm:"index"`
	CreatedAt      time.Time      `json:"createdAt"`
	UpdatedAt      time.Time      `json:"updatedAt"`
	DeletedAt      gorm.DeletedAt `json:"-" gorm:"index"`
}

type UndergraduateEducation struct {
	ID           uint           `json:"id" gorm:"primaryKey"`
	Name         string         `json:"name" gorm:"size:180;not null;uniqueIndex"`
	Grade        string         `json:"grade" gorm:"size:80"`
	Major        string         `json:"major" gorm:"size:180"`
	Direction    string         `json:"direction" gorm:"size:255"`
	Achievements StringList     `json:"achievements" gorm:"type:json"`
	Status       string         `json:"status" gorm:"size:20;index;default:published"`
	SortOrder    int            `json:"sortOrder" gorm:"index"`
	CreatedAt    time.Time      `json:"createdAt"`
	UpdatedAt    time.Time      `json:"updatedAt"`
	DeletedAt    gorm.DeletedAt `json:"-" gorm:"index"`
}

type Publication struct {
	ID        uint              `json:"id" gorm:"primaryKey"`
	ImageURL  string            `json:"image" gorm:"size:500"`
	Title     string            `json:"title" gorm:"size:500;not null;uniqueIndex"`
	Authors   string            `json:"authors" gorm:"type:text"`
	Venue     string            `json:"venue" gorm:"size:500"`
	Year      int               `json:"year" gorm:"index"`
	Kind      string            `json:"kind" gorm:"size:40;index"`
	Status    string            `json:"status" gorm:"size:20;index;default:published"`
	Featured  bool              `json:"featured" gorm:"index"`
	SortOrder int               `json:"sortOrder" gorm:"index"`
	Links     []PublicationLink `json:"links" gorm:"constraint:OnDelete:CASCADE"`
	CreatedAt time.Time         `json:"createdAt"`
	UpdatedAt time.Time         `json:"updatedAt"`
	DeletedAt gorm.DeletedAt    `json:"-" gorm:"index"`
}

type PublicationLink struct {
	ID            uint      `json:"id" gorm:"primaryKey"`
	PublicationID uint      `json:"publicationId" gorm:"index;not null"`
	Type          string    `json:"type" gorm:"size:40;not null"`
	Label         string    `json:"label" gorm:"size:80"`
	URL           string    `json:"url" gorm:"size:500"`
	RouteName     string    `json:"routeName" gorm:"size:120"`
	SortOrder     int       `json:"sortOrder"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}

type Patent struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Authors   string         `json:"authors" gorm:"type:text;not null"`
	Title     string         `json:"title" gorm:"type:varchar(700);not null;uniqueIndex"`
	Date      string         `json:"date" gorm:"size:40"`
	Country   string         `json:"country" gorm:"size:120"`
	Number    string         `json:"number" gorm:"size:180;index"`
	Category  string         `json:"category" gorm:"size:60;index"`
	Status    string         `json:"status" gorm:"size:20;index;default:published"`
	SortOrder int            `json:"sortOrder" gorm:"index"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type ResearchProject struct {
	ID            uint           `json:"id" gorm:"primaryKey"`
	Title         string         `json:"title" gorm:"size:500;not null;uniqueIndex"`
	Fund          string         `json:"fund" gorm:"size:500"`
	Number        string         `json:"number" gorm:"size:180;index"`
	Period        string         `json:"period" gorm:"size:180"`
	Amount        string         `json:"amount" gorm:"size:120"`
	ProjectStatus string         `json:"projectStatus" gorm:"column:project_status;size:80;index"`
	Role          string         `json:"role" gorm:"size:120"`
	Status        string         `json:"status" gorm:"size:20;index;default:published"`
	SortOrder     int            `json:"sortOrder" gorm:"index"`
	CreatedAt     time.Time      `json:"createdAt"`
	UpdatedAt     time.Time      `json:"updatedAt"`
	DeletedAt     gorm.DeletedAt `json:"-" gorm:"index"`
}

func normalizeStatus(status string) string {
	status = strings.TrimSpace(status)
	if status == StatusDraft {
		return StatusDraft
	}
	return StatusPublished
}
