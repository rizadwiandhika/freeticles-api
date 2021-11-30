package presentation

import (
	jsonPackage "encoding/json"
	"net/http"
	"path"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/rizadwiandhika/miniproject-backend-alterra/config"
	"github.com/rizadwiandhika/miniproject-backend-alterra/features/articles"
	"github.com/rizadwiandhika/miniproject-backend-alterra/features/articles/presentation/request"
	"github.com/rizadwiandhika/miniproject-backend-alterra/features/articles/presentation/response"
	"github.com/rizadwiandhika/miniproject-backend-alterra/helpers"
)

type any interface{}
type json map[string]any

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

	return c.JSON(http.StatusOK, json{"articles": response.FromSliceArticleCore(articles)})
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

	return c.JSON(http.StatusOK, json{"articles": response.FromSliceArticleCore(articles)})
}

func (ap *ArticlePresentation) GetDetailArticle(c echo.Context) error {
	var id uint
	echo.PathParamsBinder(c).Uint("id", &id)

	article, err, status := ap.articleBusiness.FindArticleById(id)
	if err != nil {
		return c.JSON(status, json{
			"message": "Could not get article",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, json{"articles": response.FromArticleCore(&article)})
}

func (ap *ArticlePresentation) GetTranslatedDetailArticle(c echo.Context) error {
	var id uint
	var lang string
	echo.PathParamsBinder(c).Uint("id", &id)
	echo.PathParamsBinder(c).String("lang", &lang)

	article, err, status := ap.articleBusiness.FindTranslatedArticleById(id, lang)
	if err != nil {
		return c.JSON(status, json{
			"message": "Could not get article",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, json{"articles": response.FromArticleCore(&article)})
}

func (ap *ArticlePresentation) PostArticle(c echo.Context) error {
	issuer := c.Get("user").(jwt.MapClaims)
	username := issuer["username"].(string)
	userID := uint(issuer["userId"].(float64))

	var article request.Article
	var pathDestination string

	err := c.Bind(&article)
	if err != nil {
		return c.JSON(http.StatusBadRequest, json{
			"message": "Failed posting article",
			"error":   err.Error(),
		})
	}

	tags := []string{}
	err = jsonPackage.Unmarshal([]byte(article.Tags), &tags)
	if err != nil {
		return c.JSON(http.StatusBadRequest, json{
			"message": "Failed posting article. Invalid tags JSON",
			"error":   err.Error(),
		})
	}

	thumbnail, err := c.FormFile("thumbnail")
	if err == nil {
		// file was provided
		currentTime := time.Now().Format("2006-01-02T15-04-05")
		filename := currentTime + "_" + thumbnail.Filename
		pathDestination = path.Join(config.WORKING_DIR, "thumbnails", username, filename)

		err = helpers.SaveFile(thumbnail, pathDestination)
		if err != nil {
			return c.JSON(http.StatusBadRequest, json{
				"message": "Failed uploading thumbnail",
				"error":   err.Error(),
			})
		}
	}

	tagCores := make([]articles.TagCore, len(tags))
	for i, tag := range tags {
		tagCores[i] = articles.TagCore{Tag: tag}
	}

	articleCore := articles.ArticleCore{
		AuthorID:  userID,
		Tags:      tagCores,
		Title:     article.Title,
		Subtitle:  article.Subtitle,
		Content:   article.Content,
		Thumbnail: pathDestination,
	}

	createdArticle, err, status := ap.articleBusiness.CreateArticle(articleCore)
	if err != nil {
		return c.JSON(status, json{
			"message": "Failed creating article",
			"error":   err.Error(),
		})
	}

	return c.JSON(status, json{
		"message": "Article has been created",
		"article": response.FromArticleCoreToModifiedArticle(&createdArticle),
	})
}

func (ap *ArticlePresentation) PutEditArticle(c echo.Context) error {
	issuer := c.Get("user").(jwt.MapClaims)
	username := issuer["username"].(string)
	userID := uint(issuer["userId"].(float64))

	var article request.Article
	var pathDestination string

	err := c.Bind(&article)
	if err != nil {
		return c.JSON(http.StatusBadRequest, json{
			"message": "Failed updating article",
			"error":   err.Error(),
		})
	}

	tags := []string{}
	err = jsonPackage.Unmarshal([]byte(article.Tags), &tags)
	if err != nil {
		return c.JSON(http.StatusBadRequest, json{
			"message": "Failed posting article. Invalid tags JSON",
			"error":   err.Error(),
		})
	}

	existingArticle, err, status := ap.articleBusiness.FindArticleById(article.ID)
	if err != nil {
		return c.JSON(status, json{
			"message": "Something went wrong...",
			"error":   err.Error(),
		})
	}
	if issuer["role"] != "admin" && existingArticle.AuthorID != userID {
		return c.JSON(http.StatusForbidden, json{
			"message": "You are not authorized to edit this article",
		})
	}

	thumbnail, err := c.FormFile("thumbnail")
	if err == nil {
		// file was provided
		currentTime := time.Now().Format("2006-01-02T15-04-05")
		filename := currentTime + "_" + thumbnail.Filename
		pathDestination = path.Join(config.WORKING_DIR, "thumbnails", username, filename)

		err = helpers.SaveFile(thumbnail, pathDestination)
		if err != nil {
			return c.JSON(http.StatusBadRequest, json{
				"message": "Failed uploading thumbnail",
				"error":   err.Error(),
			})
		}
	}

	tagCores := make([]articles.TagCore, len(tags))
	for i, tag := range tags {
		tagCores[i] = articles.TagCore{Tag: tag}
	}

	updatedArticleCore := articles.ArticleCore{
		ID:        article.ID,
		Tags:      tagCores,
		Title:     article.Title,
		Subtitle:  article.Subtitle,
		Content:   article.Content,
		Thumbnail: pathDestination,
	}

	updatedArticle, err, status := ap.articleBusiness.EditArticle(updatedArticleCore)
	if err != nil {
		return c.JSON(status, json{
			"message": "Failed updating article",
			"error":   err.Error(),
		})
	}

	return c.JSON(status, json{
		"message": "Article has been updated!",
		"article": response.FromArticleCoreToModifiedArticle(&updatedArticle),
	})
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

	issuer := c.Get("user").(jwt.MapClaims)
	userID := uint(issuer["userId"].(float64))

	existingArticle, err, _ := ap.articleBusiness.FindArticleById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, json{
			"message": "Something went wrong...",
			"error":   err.Error(),
		})
	}

	if issuer["role"] != "admin" && existingArticle.AuthorID != userID {
		return c.JSON(http.StatusForbidden, json{
			"message": "You are not authorized to delete this article",
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
