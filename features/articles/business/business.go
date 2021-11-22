package business

import (
	"net/http"

	"github.com/rizadwiandhika/miniproject-backend-alterra/features/articles"
	"github.com/rizadwiandhika/miniproject-backend-alterra/features/users"
)

type articleBusiness struct {
	articleData  articles.IData
	userBusiness users.IBusiness
}

func NewBusiness(data articles.IData, userBusiness users.IBusiness) *articleBusiness {
	return &articleBusiness{
		articleData:  data,
		userBusiness: userBusiness,
	}
}

func (ab *articleBusiness) FindArticles(params articles.QueryParams) ([]articles.ArticleCore, error, int) {
	articles, err := ab.articleData.SelectArticles(params)
	if err != nil {
		return nil, err, http.StatusInternalServerError
	}

	return articles, nil, http.StatusOK
}

func (ab *articleBusiness) FindArticleById(id uint) (articles.ArticleCore, error, int) {
	articleData, err := ab.articleData.SelectArticleById(id)
	if err != nil {
		return articles.ArticleCore{}, err, http.StatusInternalServerError
	}

	userData, err := ab.userBusiness.FindUserById(articleData.AuthorID)
	if err != nil {
		return articles.ArticleCore{}, err, http.StatusInternalServerError
	}

	articleData.Author.Username = userData.Username
	articleData.Author.Email = userData.Email
	articleData.Author.Name = userData.Name

	return articleData, nil, http.StatusOK
}

func (ab *articleBusiness) RemoveArticleById(id uint) (error, int) {
	err := ab.articleData.DeleteArticleById(id)
	if err != nil {
		return err, http.StatusInternalServerError
	}

	return nil, http.StatusAccepted
}

func (ab *articleBusiness) CreateArticle(article articles.ArticleCore) (articles.ArticleCore, error, int) {
	createdArticle, err := ab.articleData.InsertArticle(article)
	if err != nil {
		return article, err, http.StatusInternalServerError
	}

	return createdArticle, nil, http.StatusOK
}

func (ab *articleBusiness) EditArticle(article articles.ArticleCore) (articles.ArticleCore, error, int) {
	editedArticle, err := ab.articleData.UpdateArticle(article)
	if err != nil {
		return article, err, http.StatusInternalServerError
	}

	return editedArticle, nil, http.StatusOK
}
