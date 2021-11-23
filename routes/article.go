package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/rizadwiandhika/miniproject-backend-alterra/factory"
)

func SetupArticleRoutes(e *echo.Echo, presenter *factory.Presenter) {
	routes := e.Group("/articles")

	routes.GET("", presenter.ArticlePresentation.GetArticles)
	routes.GET("/:id", presenter.ArticlePresentation.GetDetailArticle)
	routes.DELETE("/:id", presenter.ArticlePresentation.DeleteArticle)
}
