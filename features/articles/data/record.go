package data

import (
	"time"

	"github.com/rizadwiandhika/miniproject-backend-alterra/features/articles"
)

type Article struct {
	ID        uint   `gorm:"primaryKey"`
	AuthorID  uint   `gorm:"not null"`
	Tags      []Tag  `gorm:"foreignKey:ArticleID;references:ID"`
	Title     string `gorm:"size:64"`
	Subtitle  string `gorm:"size:128"`
	Content   string
	Thumbnail string
	Nsfw      bool `gorm:"default:false"`
	UpdatedAt time.Time
	CreatedAt time.Time
}

type Tag struct {
	ID        uint   `gorm:"primaryKey"`
	ArticleID uint   `gorm:"not null"`
	Tag       string `gorm:"size:32"`
	UpdatedAt time.Time
	CreatedAt time.Time
}

func toArticleCore(a *Article) articles.ArticleCore {
	return articles.ArticleCore{
		ID:        a.ID,
		AuthorID:  a.AuthorID,
		Tags:      toSliceTagCore(a.Tags),
		Title:     a.Title,
		Subtitle:  a.Subtitle,
		Content:   a.Content,
		Thumbnail: a.Thumbnail,
		Nsfw:      a.Nsfw,
		UpdatedAt: a.UpdatedAt,
		CreatedAt: a.CreatedAt,
	}
}

func toSliceArticleCore(a []Article) []articles.ArticleCore {
	articles := make([]articles.ArticleCore, len(a))
	for i, v := range a {
		articles[i] = toArticleCore(&v)
	}

	return articles
}

func toTagCore(t *Tag) articles.TagCore {
	return articles.TagCore{
		Tag: t.Tag,
	}
}

func toSliceTagCore(t []Tag) []articles.TagCore {
	tags := make([]articles.TagCore, len(t))

	for i, v := range t {
		tags[i] = toTagCore(&v)
	}

	return tags
}
