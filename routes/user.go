package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/rizadwiandhika/miniproject-backend-alterra/factory"
)

func SetupUserRoutes(e *echo.Echo, presenter *factory.Presenter) {
	routes := e.Group("/users")

	routes.GET("/:username/articles", presenter.ArticlePresentation.GetUserArticles)
	routes.GET("/:username", presenter.UserPresentation.GetDetailUser)
	routes.DELETE("/:username", presenter.UserPresentation.DeleteUser)
	routes.PUT("/:username", presenter.UserPresentation.PutEditUser)

	routes.GET("", presenter.UserPresentation.GetUsers)
}
