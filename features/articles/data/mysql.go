package data

import (
	"github.com/rizadwiandhika/miniproject-backend-alterra/features/articles"
	"gorm.io/gorm"
)

type articleRepository struct {
	db *gorm.DB
}

func NewMySQLRepository(db *gorm.DB) *articleRepository {
	return &articleRepository{db}
}

func (ar *articleRepository) SelectArticleById(id uint) (articles.ArticleCore, error) {
	article := Article{}

	err := ar.db.Preload("Tags").First(&article, id).Error
	if err != nil {
		return articles.ArticleCore{}, err
	}

	return toArticleCore(&article), nil
}
