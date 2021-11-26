package middleware

import (
	"errors"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rizadwiandhika/miniproject-backend-alterra/config"
)

func IsAuth() echo.MiddlewareFunc {
	keyFunction := func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(config.ENV.JWT_SECRET), nil
	}

	parseToken := func(auth string, c echo.Context) (interface{}, error) {
		token, err := jwt.Parse(auth, keyFunction)
		if err != nil {
			return nil, err
		}
		if !token.Valid {
			return nil, errors.New("Invalid token")
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return nil, errors.New("Invalid claims")
		}

		return claims, nil
	}

	errorHandlerWithContext := func(err error, c echo.Context) error {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "Unauthenticated user!",
			"error":   err.Error(),
		})
	}

	jwtConfig := middleware.JWTConfig{
		ParseTokenFunc:          parseToken,
		ErrorHandlerWithContext: errorHandlerWithContext,
	}

	return middleware.JWTWithConfig(jwtConfig)
}
