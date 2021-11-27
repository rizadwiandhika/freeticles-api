package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/rizadwiandhika/miniproject-backend-alterra/factory"
	"github.com/rizadwiandhika/miniproject-backend-alterra/middleware"
)

func SetupArticleRoutes(e *echo.Echo, presenter *factory.Presenter) {
	routes := e.Group("/articles")

	routes.DELETE("/:id/likes", presenter.ReactionPresentation.DeleteLike, middleware.IsAuth())
	routes.DELETE("/:id", presenter.ArticlePresentation.DeleteArticle)

	routes.POST(
		"/:id/likes",
		presenter.ReactionPresentation.PostLike,
		middleware.IsAuth(),
	)
	routes.POST("/:id/comments", presenter.ReactionPresentation.PostComment, middleware.IsAuth())
	routes.POST("/:id/reports", presenter.ReactionPresentation.PostReport, middleware.IsAuth())

	routes.GET("/:id", presenter.ArticlePresentation.GetDetailArticle)
	routes.GET("", presenter.ArticlePresentation.GetArticles)
}
