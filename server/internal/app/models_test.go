package app

import (
	"testing"
	"time"

	"gorm.io/gorm"
)

func TestNormalizeStatus(t *testing.T) {
	if got := normalizeStatus(StatusDraft); got != StatusDraft {
		t.Fatalf("expected draft, got %q", got)
	}
	if got := normalizeStatus(""); got != StatusPublished {
		t.Fatalf("expected published fallback, got %q", got)
	}
	if got := normalizeStatus("unknown"); got != StatusPublished {
		t.Fatalf("expected published for unknown status, got %q", got)
	}
}

func TestSanitizeFileName(t *testing.T) {
	tests := map[string]string{
		"My Paper 2026": "my-paper-2026",
		"../bad name":   "bad-name",
		"中文 文件":         "file",
	}
	for input, expected := range tests {
		if got := sanitizeFileName(input); got != expected {
			t.Fatalf("sanitizeFileName(%q) = %q, want %q", input, got, expected)
		}
	}
}

func TestTrashItemIncludesDeletedAt(t *testing.T) {
	deletedAt := time.Now()
	item := trashItem("news", "News", 12, "Recoverable News", "Publication", StatusPublished, gorm.DeletedAt{
		Time:  deletedAt,
		Valid: true,
	})

	if item.Resource != "news" || item.Label != "News" || item.ID != 12 {
		t.Fatalf("unexpected trash metadata: %+v", item)
	}
	if item.DeletedAt == nil || !item.DeletedAt.Equal(deletedAt) {
		t.Fatalf("expected deleted timestamp, got %+v", item.DeletedAt)
	}
}

func TestCompactSubtitle(t *testing.T) {
	if got := compactSubtitle(" Paper ", "", "2026-05-25"); got != "Paper / 2026-05-25" {
		t.Fatalf("unexpected compact subtitle: %q", got)
	}
}

func TestLegacyVideoRoutePath(t *testing.T) {
	if got := legacyVideoRoutePath("video-player-XiaoqiZheng01"); got != "/video/video-xiaoqi-zheng-01" {
		t.Fatalf("unexpected legacy video path: %q", got)
	}
	if got := legacyVideoRoutePath("unknown"); got != "" {
		t.Fatalf("expected empty path for unknown route, got %q", got)
	}
}
