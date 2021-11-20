package factory

import (
	"github.com/rizadwiandhika/miniproject-backend-alterra/config"
	articlesBusiness "github.com/rizadwiandhika/miniproject-backend-alterra/features/articles/business"
	articlesData "github.com/rizadwiandhika/miniproject-backend-alterra/features/articles/data"
	articlesPresentation "github.com/rizadwiandhika/miniproject-backend-alterra/features/articles/presentation"
	usersBusiness "github.com/rizadwiandhika/miniproject-backend-alterra/features/users/business"
	usersData "github.com/rizadwiandhika/miniproject-backend-alterra/features/users/data"
)

type Presenter struct {
	ArticlePresentation *articlesPresentation.ArticlePresentation
}

func New() *Presenter {
	userData := usersData.NewMySQLRepository(config.DB)
	articleData := articlesData.NewMySQLRepository(config.DB)

	userBusiness := usersBusiness.NewBusiness(userData)
	articleBusiness := articlesBusiness.NewBusiness(articleData, userBusiness)

	articlePresentation := articlesPresentation.NewPresentation(articleBusiness)

	return &Presenter{
		ArticlePresentation: articlePresentation,
	}
}
