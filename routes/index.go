package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rizadwiandhika/miniproject-backend-alterra/factory"
)

func Setup() *echo.Echo {
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())

	presenter := factory.New()

	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}` + "\n",
	}))

	SetupArticleRoutes(e, presenter)
	SetupUserRoutes(e, presenter)
	SetupAuthRoutes(e, presenter)

	return e
}
