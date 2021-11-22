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
