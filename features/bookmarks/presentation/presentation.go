package presentation

import (
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/rizadwiandhika/miniproject-backend-alterra/features/bookmarks"
	"github.com/rizadwiandhika/miniproject-backend-alterra/features/bookmarks/presentation/response"
)

type any interface{}
type json map[string]any

type BookmarkPresentation struct {
	bookmarkBusiness bookmarks.IBusiness
}

func NewPresentation(bb bookmarks.IBusiness) *BookmarkPresentation {
	return &BookmarkPresentation{
		bookmarkBusiness: bb,
	}
}

func (bp *BookmarkPresentation) GetUserBookmarks(c echo.Context) error {
	var username string
	echo.PathParamsBinder(c).String("username", &username)

	issuer := c.Get("user").(jwt.MapClaims)
	if issuer["username"] != username && issuer["role"] != "admin" {
		return c.JSON(http.StatusForbidden, json{
			"message": "User is not authorized",
		})
	}

	bookmarks, err := bp.bookmarkBusiness.FindUserBookmarks(username)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, json{
			"message": "Failed getting user bookmarks",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, json{
		"message":   "Success retrieving user bookmarks",
		"bookmarks": response.FromSliceBookmarkCore(bookmarks),
	})
}

func (bp *BookmarkPresentation) PostBookmark(c echo.Context) error {
	var username string
	var articleID uint

	echo.PathParamsBinder(c).String("username", &username)
	echo.PathParamsBinder(c).Uint("articleid", &articleID)

	issuer := c.Get("user").(jwt.MapClaims)
	if issuer["username"] != username {
		return c.JSON(http.StatusBadRequest, json{
			"message": "Invalid username!",
		})
	}

	err := bp.bookmarkBusiness.CreateBookmark(username, articleID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, json{
			"message": "Failed creating bookmark",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, json{
		"message": "Bookmark created",
	})
}

func (bp *BookmarkPresentation) DeleteBookmark(c echo.Context) error {
	var username string
	var articleID uint

	echo.PathParamsBinder(c).String("username", &username)
	echo.PathParamsBinder(c).Uint("articleid", &articleID)

	issuer := c.Get("user").(jwt.MapClaims)
	if issuer["username"] != username {
		return c.JSON(http.StatusBadRequest, json{
			"message": "Invalid username!",
		})
	}

	err := bp.bookmarkBusiness.DeleteBookmark(username, articleID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, json{
			"message": "Failed deleting bookmark",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, json{
		"message": "Bookmark deleted",
	})
}
