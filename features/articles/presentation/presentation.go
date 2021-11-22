package presentation

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rizadwiandhika/miniproject-backend-alterra/features/articles"
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

func (ap ArticlePresentation) GetDetailArticle(c echo.Context) error {
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
