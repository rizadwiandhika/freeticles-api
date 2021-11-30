package data

import (
	"fmt"
	"time"

	"github.com/rizadwiandhika/miniproject-backend-alterra/features/articles"
	"gorm.io/gorm"
)

type articleRepository struct {
	db *gorm.DB
}

func NewMySQLRepository(db *gorm.DB) *articleRepository {
	return &articleRepository{db}
}

func (ar *articleRepository) SelectArticles(q articles.QueryParams) ([]articles.ArticleCore, error) {
	articles := []Article{}
	tx := ar.db.Debug()

	if q.Keyword != "" {
		keyword := "%" + "anime" + "%"
		tx = tx.Joins(
			"JOIN tags ON (tags.article_id = articles.id AND (tags.tag LIKE ? OR articles.title LIKE ? OR articles.subtitle LIKE ?))",
			keyword,
			keyword,
			keyword,
		)
	}

	if q.Today {
		current := time.Now()
		today := fmt.Sprintf("%d-%02d-%02d", current.Year(), current.Month(), current.Day())
		tx = tx.Where("DATE(articles.created_at) = ?", today) // today will be like '2021-04-22'
	}

	err := tx.Preload("Tags").Limit(q.Limit).Offset(q.Offset).Find(&articles).Error
	return toSliceArticleCore(articles), err
}

func (ar *articleRepository) SelectArticleById(id uint) (articles.ArticleCore, error) {
	article := Article{}

	err := ar.db.Preload("Tags").First(&article, id).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return articles.ArticleCore{}, err
	}

	return toArticleCore(&article), nil
}

func (ar *articleRepository) SelectArticlesByAuthorId(id uint) ([]articles.ArticleCore, error) {
	articles := []Article{}
	err := ar.db.Where("author_id = ?", id).Find(&articles).Error
	if err != nil {
		return nil, err
	}

	return toSliceArticleCore(articles), nil
}

func (ar *articleRepository) DeleteArticleById(id uint) error {
	return ar.db.Delete(&Article{}, id).Error
}

func (ar *articleRepository) DeleteArticlesByUserId(userID uint) error {
	return ar.db.Where("author_id = ?", userID).Delete(Article{}).Error
}

func (ar *articleRepository) DeleteArticleTags(id uint) error {
	return ar.db.Where("article_id = ?", id).Delete(Tag{}).Error
}

func (ar *articleRepository) DeleteTagByArticleIds(id []uint) error {
	return ar.db.Where("article_id IN (?)", id).Delete(Tag{}).Error
}

func (ar *articleRepository) InsertArticle(article articles.ArticleCore) (articles.ArticleCore, error) {
	tags := make([]Tag, len(article.Tags))
	for i, v := range article.Tags {
		tags[i] = Tag{Tag: v.Tag}
	}

	newArticle := Article{
		AuthorID:  article.AuthorID,
		Title:     article.Title,
		Subtitle:  article.Subtitle,
		Content:   article.Content,
		Thumbnail: article.Thumbnail,
		Tags:      tags,
	}

	err := ar.db.Create(&newArticle).Error
	return toArticleCore(&newArticle), err
}

func (ar *articleRepository) UpdateArticle(article articles.ArticleCore) (articles.ArticleCore, error) {
	updatedTags := make([]Tag, len(article.Tags))

	for i, tag := range article.Tags {
		updatedTags[i] = Tag{Tag: tag.Tag, ArticleID: article.ID}
	}

	updatedArticle := Article{
		ID:        article.ID,
		AuthorID:  article.AuthorID,
		Title:     article.Title,
		Subtitle:  article.Subtitle,
		Content:   article.Content,
		Thumbnail: article.Thumbnail,
		Nsfw:      article.Nsfw,
		Tags:      updatedTags,
		CreatedAt: article.CreatedAt,
	}

	err := ar.db.Where("article_id = ?", article.ID).Delete(Tag{}).Error
	if err != nil {
		return articles.ArticleCore{}, err
	}

	err = ar.db.Save(&updatedArticle).Error
	if err != nil {
		return articles.ArticleCore{}, err
	}

	return toArticleCore(&updatedArticle), nil
}
