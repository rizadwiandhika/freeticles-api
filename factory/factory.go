package factory

import (
	"github.com/rizadwiandhika/miniproject-backend-alterra/config"
	articlesBusiness "github.com/rizadwiandhika/miniproject-backend-alterra/features/articles/business"
	articlesData "github.com/rizadwiandhika/miniproject-backend-alterra/features/articles/data"
	articlesPresentation "github.com/rizadwiandhika/miniproject-backend-alterra/features/articles/presentation"
	usersBusiness "github.com/rizadwiandhika/miniproject-backend-alterra/features/users/business"
	usersData "github.com/rizadwiandhika/miniproject-backend-alterra/features/users/data"
	usersPresentation "github.com/rizadwiandhika/miniproject-backend-alterra/features/users/presentation"
)

type Presenter struct {
	ArticlePresentation *articlesPresentation.ArticlePresentation
	UserPresentation    *usersPresentation.UserPresentation
}

func New() *Presenter {
	userData := usersData.NewMySQLRepository(config.DB)
	articleData := articlesData.NewMySQLRepository(config.DB)

	userBusiness := usersBusiness.NewBusiness(userData)
	articleBusiness := articlesBusiness.NewBusiness(articleData, userBusiness)

	userPresentation := usersPresentation.NewPresentation(userBusiness)
	articlePresentation := articlesPresentation.NewPresentation(articleBusiness)

	return &Presenter{
		ArticlePresentation: articlePresentation,
		UserPresentation:    userPresentation,
	}
}
