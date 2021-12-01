package presentation_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/rizadwiandhika/miniproject-backend-alterra/features/articles"
	"github.com/rizadwiandhika/miniproject-backend-alterra/features/articles/presentation"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	amock "github.com/rizadwiandhika/miniproject-backend-alterra/features/articles/mocks"
	ap "github.com/rizadwiandhika/miniproject-backend-alterra/features/articles/presentation"

	"github.com/rizadwiandhika/miniproject-backend-alterra/routes"
)

var (
	e *echo.Echo

	articleBusiness     amock.IBusiness
	articlePresentation *presentation.ArticlePresentation

	articleValue articles.ArticleCore
)

func TestMain(m *testing.M) {
	e = routes.Setup()
	articlePresentation = ap.NewPresentation(&articleBusiness)

	articleValue = articles.ArticleCore{
		ID:       1,
		AuthorID: 1,
		Title:    "Title",
		Subtitle: "Subtitle",
		Content:  "Content",
	}

	os.Exit(m.Run())
}

func TestGetArticles(t *testing.T) {
	mockHttpRequest := func() (*http.Request, *httptest.ResponseRecorder, echo.Context) {
		req, _ := http.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		return req, rec, c
		// c.SetPath("/articles")

		// res := rec.Result()
		// defer res.Body.Close()
	}

	t.Run("valid - GetArticles", func(t *testing.T) {
		// mock incomming request
		_, rec, c := mockHttpRequest()
		defer rec.Result().Body.Close()

		var mc *mock.Call

		mc = articleBusiness.On("FindArticles", articles.QueryParams{})
		mc = mc.Return([]articles.ArticleCore{articleValue}, nil, 200)
		mc.Once()

		if assert.NoError(t, articlePresentation.GetArticles(c)) {
			assert.NotEqual(t, "", rec.Body.String())
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})

	t.Run("valid - error when FindArticles failed", func(t *testing.T) {
		// mock incomming request
		_, rec, c := mockHttpRequest()
		defer rec.Result().Body.Close()

		var mc *mock.Call

		mc = articleBusiness.On("FindArticles", articles.QueryParams{})
		mc = mc.Return(nil, errors.New("error"), 500)
		mc.Once()

		if assert.NoError(t, articlePresentation.GetArticles(c)) {
			assert.NotEqual(t, "", rec.Body.String())
			assert.Equal(t, http.StatusInternalServerError, rec.Code)
		}
	})
}

func TestGetUserArticles(t *testing.T) {
	mockHttpRequest := func() (*http.Request, *httptest.ResponseRecorder, echo.Context) {
		req, _ := http.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		return req, rec, c
	}

	t.Run("valid - GetUserArticles", func(t *testing.T) {
		_, rec, c := mockHttpRequest()
		defer rec.Result().Body.Close()

		c.SetParamNames("username")
		c.SetParamValues("riza.dwi")

		var mc *mock.Call

		mc = articleBusiness.On("FindUserArticles", "riza.dwi")
		mc = mc.Return([]articles.ArticleCore{articleValue}, nil, 200)
		mc.Once()

		if assert.NoError(t, articlePresentation.GetUserArticles(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.NotEqual(t, "", rec.Body.String())
		}
	})

	t.Run("valid - error when FindUserArticles failed", func(t *testing.T) {
		_, rec, c := mockHttpRequest()
		defer rec.Result().Body.Close()

		c.SetParamNames("username")
		c.SetParamValues("riza.dwi")

		var mc *mock.Call

		mc = articleBusiness.On("FindUserArticles", "riza.dwi")
		mc = mc.Return(nil, errors.New("error"), 500)
		mc.Once()

		if assert.NoError(t, articlePresentation.GetUserArticles(c)) {
			assert.Equal(t, http.StatusInternalServerError, rec.Code)
			assert.NotEqual(t, "", rec.Body.String())
		}
	})
}

func TestGetDetailArticle(t *testing.T) {
	mockHttpRequest := func() (*http.Request, *httptest.ResponseRecorder, echo.Context) {
		req, _ := http.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		return req, rec, c
	}

	t.Run("valid - GetUserArticles", func(t *testing.T) {
		_, rec, c := mockHttpRequest()
		defer rec.Result().Body.Close()

		c.SetPath("/articles/:id")
		c.SetParamNames("id")
		c.SetParamValues("1")

		var mc *mock.Call

		mc = articleBusiness.On("FindArticleById", mock.AnythingOfType("uint"))
		mc = mc.Return(articleValue, nil, 200)
		mc.Once()

		if assert.NoError(t, articlePresentation.GetDetailArticle(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.NotEqual(t, "", rec.Body.String())
		}
	})

	t.Run("valid - error when FindArticleById failed", func(t *testing.T) {
		_, rec, c := mockHttpRequest()
		defer rec.Result().Body.Close()

		c.SetPath("/articles/:id")
		c.SetParamNames("username")
		c.SetParamValues("riza.dwi")

		var mc *mock.Call

		mc = articleBusiness.On("FindArticleById", mock.AnythingOfType("uint"))
		mc = mc.Return(articles.ArticleCore{}, errors.New("error"), 500)
		mc.Once()

		if assert.NoError(t, articlePresentation.GetDetailArticle(c)) {
			assert.Equal(t, http.StatusInternalServerError, rec.Code)
			assert.NotEqual(t, "", rec.Body.String())
		}
	})
}

func TestGetTranslatedDetailArticle(t *testing.T) {
	mockHttpRequest := func() (*http.Request, *httptest.ResponseRecorder, echo.Context) {
		req, _ := http.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		return req, rec, c
	}

	t.Run("valid - GetTranslatedDetailArticle", func(t *testing.T) {
		_, rec, c := mockHttpRequest()
		defer rec.Result().Body.Close()

		c.SetPath("/:en/articles/:id/translate")
		c.SetParamNames("lang", "id")
		c.SetParamValues("en", "1")

		var mc *mock.Call

		mc = articleBusiness.On("FindTranslatedArticleById", uint(1), "en")
		mc = mc.Return(articleValue, nil, 200)
		mc.Once()

		if assert.NoError(t, articlePresentation.GetTranslatedDetailArticle(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.NotEqual(t, "", rec.Body.String())
		}
	})

	t.Run("valid - error when FindTranslatedArticleById failed", func(t *testing.T) {
		_, rec, c := mockHttpRequest()
		defer rec.Result().Body.Close()

		c.SetPath("/:en/articles/:id/translate")
		c.SetParamNames("lang", "id")
		c.SetParamValues("en", "1")

		var mc *mock.Call

		mc = articleBusiness.On("FindTranslatedArticleById", uint(1), "en")
		mc = mc.Return(articles.ArticleCore{}, errors.New("error"), 500)
		mc.Once()

		if assert.NoError(t, articlePresentation.GetTranslatedDetailArticle(c)) {
			assert.Equal(t, http.StatusInternalServerError, rec.Code)
			assert.NotEqual(t, "", rec.Body.String())
		}
	})

}
