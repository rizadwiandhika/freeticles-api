package presentation

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rizadwiandhika/miniproject-backend-alterra/features/users"
)

type any interface{}
type json map[string]any
type list []any

type UserPresentation struct {
	userBusiness users.IBusiness
}

func NewPresentation(articleBusiness users.IBusiness) *UserPresentation {
	return &UserPresentation{articleBusiness}
}
func (up *UserPresentation) GetDetailUser(c echo.Context) error {
	var username string
	echo.PathParamsBinder(c).String("id", &username)

	user, err := up.userBusiness.FindUserByUsername(username)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, json{
			"message": "Could not get user",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, json{"user": user})
}
