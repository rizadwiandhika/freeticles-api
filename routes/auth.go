package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/rizadwiandhika/miniproject-backend-alterra/factory"
)

func SetupAuthRoutes(e *echo.Echo, presenter *factory.Presenter) {
	routes := e.Group("/auth")

	routes.POST("/register", presenter.AuthPresentation.PostRegister)
	routes.POST("/login", presenter.AuthPresentation.PostLogin)

}
