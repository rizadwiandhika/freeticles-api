package business

import "github.com/rizadwiandhika/miniproject-backend-alterra/features/articles"

type articleBusiness struct {
	articleData articles.IData
}

func NewBusiness(data articles.IData) *articleBusiness {
	return &articleBusiness{articleData: data}
}

func (ab *articleBusiness) FindArticleById(id uint) (articles.ArticleCore, error) {
	articleData, err := ab.articleData.SelectArticleById(id)
	if err != nil {
		return articles.ArticleCore{}, err
	}

	return articleData, nil
}
