package presentation

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rizadwiandhika/miniproject-backend-alterra/features/articles"
	"github.com/rizadwiandhika/miniproject-backend-alterra/features/articles/presentation/request"
)

type any interface{}
type json map[string]any
type list []any

type ArticlePresentation struct {
	articleBusiness articles.IBusiness
}

func NewPresentation(articleBusiness articles.IBusiness) *ArticlePresentation {
	return &ArticlePresentation{articleBusiness}
}

func (ap *ArticlePresentation) GetArticles(c echo.Context) error {
	query := request.QueryParams{}

	err := c.Bind(&query)
	if err != nil {
		return c.JSON(http.StatusBadRequest, json{
			"message": "Could not get articles",
			"error":   err.Error(),
		})
	}

	queryCore := request.ToQueryParamsCore(&query)
	articles, err, status := ap.articleBusiness.FindArticles(queryCore)
	if err != nil {
		return c.JSON(status, json{
			"message": "Could not get articles",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, json{"articles": articles})
}

func (ap *ArticlePresentation) GetUserArticles(c echo.Context) error {
	var username string
	echo.PathParamsBinder(c).String("username", &username)

	articles, err, status := ap.articleBusiness.FindUserArticles(username)
	if err != nil {
		return c.JSON(status, json{
			"message": "Could not get user articles",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, json{"articles": articles})
}

func (ap *ArticlePresentation) GetDetailArticle(c echo.Context) error {
	var id uint
	echo.PathParamsBinder(c).Uint("id", &id)

	articles, err, status := ap.articleBusiness.FindArticleById(id)
	if err != nil {
		return c.JSON(status, json{
			"message": "Could not get article",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, json{"articles": articles})
}

func (ap *ArticlePresentation) DeleteArticle(c echo.Context) error {
	var id uint
	echo.PathParamsBinder(c).Uint("id", &id)

	if id < 1 {
		return c.JSON(http.StatusBadRequest, json{
			"message": "Failed deleting article",
			"error":   "id must be greater than 0",
			"id":      id,
		})
	}

	err, status := ap.articleBusiness.RemoveArticleById(id)
	if err != nil {
		return c.JSON(status, json{
			"message": "Could not delete article",
			"error":   err.Error(),
			"id":      id,
		})
	}

	return c.JSON(http.StatusOK, json{"message": "Article has been deleted"})
}
