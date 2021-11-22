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

func (ar *articleRepository) SelectArticles(q articles.QueryParams) ([]articles.ArticleCore, error) {
	return nil, nil
}

func (ar *articleRepository) SelectArticleById(id uint) (articles.ArticleCore, error) {
	article := Article{}

	err := ar.db.Preload("Tags").First(&article, id).Error
	if err != nil {
		return articles.ArticleCore{}, err
	}

	return toArticleCore(&article), nil
}

func (ar *articleRepository) SelectArticleByKeyword(keyword string) ([]articles.ArticleCore, error) {
	return nil, nil
}

func (ar *articleRepository) DeleteArticleById(id int) (articles.ArticleCore, error) {
	return articles.ArticleCore{}, nil
}

func (ar *articleRepository) InsertArticle(article articles.ArticleCore) (articles.ArticleCore, error) {
	return articles.ArticleCore{}, nil
}

func (ar *articleRepository) UpdateArticle(article articles.ArticleCore) error {
	return nil
}
