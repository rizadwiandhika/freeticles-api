package data

import (
	"time"

	"github.com/rizadwiandhika/miniproject-backend-alterra/features/reactions"
)

type Comment struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint `gorm:"not null"`
	ArticleID uint `gorm:"not null"`
	Commentar string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Like struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint `gorm:"not null"`
	ArticleID uint `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Report struct {
	ID           uint       `gorm:"primaryKey"`
	UserID       uint       `gorm:"not null"`
	ArticleID    uint       `gorm:"not null"`
	ReportTypeID uint       `gorm:"not null"`
	ReportType   ReportType `gorm:"foreignKey:ReportTypeID;references:ID"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type ReportType struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"unique;size:32"`
	Description string `gorm:"size:128"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type SliceComment []Comment
type SliceLike []Like
type SliceReport []Report
type SliceReportType []ReportType

func (c *Comment) toCommentCore() reactions.CommentCore {
	return reactions.CommentCore{
		ID:        c.ID,
		UserID:    c.UserID,
		ArticleID: c.ArticleID,
		CreatedAt: c.CreatedAt,
	}
}

func (l *Like) toLikeCore() reactions.LikeCore {
	return reactions.LikeCore{
		ID:        l.ID,
		UserID:    l.UserID,
		ArticleID: l.ArticleID,
	}
}

func (r *Report) toReportCore() reactions.ReportCore {
	return reactions.ReportCore{
		ID:           r.ID,
		UserID:       r.UserID,
		ArticleID:    r.ArticleID,
		ReportTypeID: r.ReportTypeID,
		CreatedAt:    r.CreatedAt,
	}
}

func (c SliceComment) toSliceCommentCore() []reactions.CommentCore {
	commentCores := make([]reactions.CommentCore, len(c))
	for i, comment := range c {
		commentCores[i] = comment.toCommentCore()
	}
	return commentCores
}

func (l SliceLike) toSliceLikeCore() []reactions.LikeCore {
	likeCores := make([]reactions.LikeCore, len(l))
	for i, comment := range l {
		likeCores[i] = comment.toLikeCore()
	}
	return likeCores
}

func (r SliceReport) toSliceReportCore() []reactions.ReportCore {
	reportCores := make([]reactions.ReportCore, len(r))
	for i, comment := range r {
		reportCores[i] = comment.toReportCore()
	}
	return reportCores
}
