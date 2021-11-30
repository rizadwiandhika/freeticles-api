package factory

import (
	"github.com/rizadwiandhika/miniproject-backend-alterra/config"
	"github.com/rizadwiandhika/miniproject-backend-alterra/third-parties/image/cloudmersive"
	"github.com/rizadwiandhika/miniproject-backend-alterra/third-parties/translate/libre"

	articlesBusiness "github.com/rizadwiandhika/miniproject-backend-alterra/features/articles/business"
	articlesData "github.com/rizadwiandhika/miniproject-backend-alterra/features/articles/data"
	articlesPresentation "github.com/rizadwiandhika/miniproject-backend-alterra/features/articles/presentation"

	authBusiness "github.com/rizadwiandhika/miniproject-backend-alterra/features/auth/business"
	authPresentation "github.com/rizadwiandhika/miniproject-backend-alterra/features/auth/presentation"

	bookmarksBusiness "github.com/rizadwiandhika/miniproject-backend-alterra/features/bookmarks/business"
	bookmarksData "github.com/rizadwiandhika/miniproject-backend-alterra/features/bookmarks/data"
	bookmarksPresentation "github.com/rizadwiandhika/miniproject-backend-alterra/features/bookmarks/presentation"

	reactionsBusiness "github.com/rizadwiandhika/miniproject-backend-alterra/features/reactions/business"
	reactionsData "github.com/rizadwiandhika/miniproject-backend-alterra/features/reactions/data"
	reactionsPresentation "github.com/rizadwiandhika/miniproject-backend-alterra/features/reactions/presentation"

	usersBusiness "github.com/rizadwiandhika/miniproject-backend-alterra/features/users/business"
	usersData "github.com/rizadwiandhika/miniproject-backend-alterra/features/users/data"
	usersPresentation "github.com/rizadwiandhika/miniproject-backend-alterra/features/users/presentation"
)

type Presenter struct {
	ArticlePresentation  *articlesPresentation.ArticlePresentation
	UserPresentation     *usersPresentation.UserPresentation
	AuthPresentation     *authPresentation.AuthPresentation
	ReactionPresentation *reactionsPresentation.ReactionPresentation
	BookmarkPresentation *bookmarksPresentation.BookmarkPresentation
}

func New() *Presenter {
	libreTranslate := libre.NewTranslate()
	cloudmersiveImageAnalyzer := cloudmersive.NewImageAnalyzer()

	userData := usersData.NewMySQLRepository(config.DB)
	articleData := articlesData.NewMySQLRepository(config.DB)
	reactionData := reactionsData.NewMySQLRepository(config.DB)
	bookmarkData := bookmarksData.NewMySQLRepository(config.DB)

	userBusiness := usersBusiness.NewBusiness(
		userData,
		articlesBusiness.NewBusiness(articleData, nil, nil, nil, nil),
		reactionsBusiness.NewBusiness(reactionData, nil, nil, nil),
		bookmarksBusiness.NewBusiness(bookmarkData, nil, nil),
	)
	articleBusiness := articlesBusiness.NewBusiness(
		articleData,
		libreTranslate,
		userBusiness,
		reactionsBusiness.NewBusiness(reactionData, nil, nil, nil),
		bookmarksBusiness.NewBusiness(bookmarkData, nil, nil),
	)
	authBusiness := authBusiness.NewBusniness(userBusiness)
	reactionBusiness := reactionsBusiness.NewBusiness(reactionData, cloudmersiveImageAnalyzer, userBusiness, articleBusiness)
	bookmarkBusiness := bookmarksBusiness.NewBusiness(bookmarkData, userBusiness, articleBusiness)

	userPresentation := usersPresentation.NewPresentation(userBusiness)
	articlePresentation := articlesPresentation.NewPresentation(articleBusiness)
	authPresentation := authPresentation.NewPresentation(authBusiness)
	reactionPresentation := reactionsPresentation.NewPresentation(reactionBusiness)
	bookmarkPresentation := bookmarksPresentation.NewPresentation(bookmarkBusiness)

	return &Presenter{
		ArticlePresentation:  articlePresentation,
		UserPresentation:     userPresentation,
		AuthPresentation:     authPresentation,
		ReactionPresentation: reactionPresentation,
		BookmarkPresentation: bookmarkPresentation,
	}
}
