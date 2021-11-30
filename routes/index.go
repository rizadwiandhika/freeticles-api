package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rizadwiandhika/miniproject-backend-alterra/factory"
)

func Setup() *echo.Echo {
	e := echo.New()

	presenter := factory.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())

	SetupArticleRoutes(e, presenter)
	SetupUserRoutes(e, presenter)
	SetupAuthRoutes(e, presenter)

	return e
}
