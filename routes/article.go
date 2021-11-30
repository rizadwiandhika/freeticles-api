package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/rizadwiandhika/miniproject-backend-alterra/factory"
	"github.com/rizadwiandhika/miniproject-backend-alterra/middleware"
)

func SetupArticleRoutes(e *echo.Echo, presenter *factory.Presenter) {
	e.GET("/:lang/articles/:id", presenter.ArticlePresentation.GetTranslatedDetailArticle)

	routes := e.Group("/articles")

	routes.DELETE("/:id/likes", presenter.ReactionPresentation.DeleteLike, middleware.IsAuth())
	routes.DELETE("/:id", presenter.ArticlePresentation.DeleteArticle, middleware.IsAuth())

	routes.PUT("", presenter.ArticlePresentation.PutEditArticle, middleware.IsAuth())

	routes.POST("/:id/likes", presenter.ReactionPresentation.PostLike, middleware.IsAuth())
	routes.POST("/:id/comments", presenter.ReactionPresentation.PostComment, middleware.IsAuth())
	routes.POST("/:id/reports", presenter.ReactionPresentation.PostReport, middleware.IsAuth())
	routes.POST("", presenter.ArticlePresentation.PostArticle, middleware.IsAuth())

	routes.GET("/:id/comments", presenter.ReactionPresentation.GetArticleComments)
	routes.GET("/:id", presenter.ArticlePresentation.GetDetailArticle)
	routes.GET("", presenter.ArticlePresentation.GetArticles)
}
