package data

import (
	"time"

	"github.com/rizadwiandhika/miniproject-backend-alterra/features/bookmarks"
)

type Bookmark struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint `gorm:"not null"`
	ArticleID uint `gorm:"not null"`
	CreatedAt time.Time
}

func (b *Bookmark) ToBookmarkCore() bookmarks.BookmarkCore {
	return bookmarks.BookmarkCore{
		ID:        b.ID,
		CreatedAt: b.CreatedAt,
		UserID:    b.UserID,
		ArticleID: b.ArticleID,
	}
}

func ToSliceBookmarkCore(b []Bookmark) []bookmarks.BookmarkCore {
	bookmarkCores := make([]bookmarks.BookmarkCore, len(b))
	for i, v := range b {
		bookmarkCores[i] = v.ToBookmarkCore()
	}
	return bookmarkCores
}
