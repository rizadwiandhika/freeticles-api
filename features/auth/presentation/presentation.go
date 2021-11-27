package presentation

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rizadwiandhika/miniproject-backend-alterra/features/auth"
	"github.com/rizadwiandhika/miniproject-backend-alterra/features/auth/presentation/request"
	"github.com/rizadwiandhika/miniproject-backend-alterra/features/auth/presentation/response"
)

type any interface{}
type json map[string]any

type AuthPresentation struct {
	ab auth.IBusiness
}

func NewPresentation(ab auth.IBusiness) *AuthPresentation {
	return &AuthPresentation{
		ab: ab,
	}
}

func (ap *AuthPresentation) PostLogin(c echo.Context) error {
	user := request.Login{}
	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, json{
			"message": "Something went wrong...",
			"error":   err.Error(),
		})
	}

	userCore := user.ToUserCore()
	jwtToken, err, status := ap.ab.Authenticate(userCore)
	if err != nil {
		return c.JSON(status, json{
			"message": "Authentication failed",
			"error":   err.Error(),
			"user":    response.ToUserLoginFailed(userCore),
		})
	}

	return c.JSON(status, json{
		"message": "Successfully login",
		"token":   jwtToken,
	})
}

func (ap *AuthPresentation) PostRegister(c echo.Context) error {
	newUser := request.Register{}
	err := c.Bind(&newUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, json{
			"message": "Something went wrong...",
			"error":   err.Error(),
		})
	}

	userCore := newUser.ToUserCore()
	createdUser, err, status := ap.ab.Register(userCore)
	if err != nil {
		return c.JSON(status, json{
			"message": "Failed creating user",
			"error":   err.Error(),
			"user":    response.ToUserRegisterFailed(userCore),
		})
	}

	return c.JSON(status, json{
		"message": "User created",
		"user":    response.ToUserRegisterSuccess(createdUser),
	})
}
