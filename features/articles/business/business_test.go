package business_test

import (
	"errors"
	"net/http"
	"os"
	"testing"

	"github.com/rizadwiandhika/miniproject-backend-alterra/features/articles"
	"github.com/rizadwiandhika/miniproject-backend-alterra/features/users"

	ab "github.com/rizadwiandhika/miniproject-backend-alterra/features/articles/business"
	amock "github.com/rizadwiandhika/miniproject-backend-alterra/features/articles/mocks"
	bmock "github.com/rizadwiandhika/miniproject-backend-alterra/features/bookmarks/mocks"
	rmock "github.com/rizadwiandhika/miniproject-backend-alterra/features/reactions/mocks"
	umock "github.com/rizadwiandhika/miniproject-backend-alterra/features/users/mocks"
	tmock "github.com/rizadwiandhika/miniproject-backend-alterra/third-parties/translate/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	articleData amock.IData

	articleBusiness   articles.IBusiness
	bookmarksBusiness bmock.IBusiness
	userBusiness      umock.IBusiness
	reactionBusiness  rmock.IBusiness
	translateBusiness tmock.IBusiness

	articleValue articles.ArticleCore
	userValue    users.UserCore
)

// bikin adapternya
func TestMain(m *testing.M) {
	articleBusiness = ab.NewBusiness(&articleData, &translateBusiness, &userBusiness, &reactionBusiness, &bookmarksBusiness)

	userValue = users.UserCore{
		Username: "rizadwi",
		Email:    "rizadwi@mail.com",
		Name:     "Rizal Dwi Andhika",
	}

	articleValue = articles.ArticleCore{
		ID:        1,
		AuthorID:  1,
		Title:     "mock title",
		Subtitle:  "mock subtitle",
		Content:   "mock content",
		Thumbnail: "mock thumbnail",
	}

	os.Exit(m.Run())
}

func TestFindArticles(t *testing.T) {
	t.Run("valid - find articles", func(t *testing.T) {
		var mc *mock.Call

		mc = articleData.On("SelectArticles", mock.Anything)
		mc = mc.Return([]articles.ArticleCore{articleValue}, nil)
		mc.Once()

		mc = userBusiness.On("FindUserById", mock.AnythingOfType("uint"))
		mc = mc.Return(userValue, nil, 200)
		mc.Once()

		mc = reactionBusiness.On("CountTotalArticleLikes", mock.AnythingOfType("uint"))
		mc = mc.Return(1, nil)
		mc.Once()

		data, _, _ := articleBusiness.FindArticles(articles.QueryParams{})

		assert.Equal(t, 1, len(data))
	})

	t.Run("valid - return error when SelectArticles failed", func(t *testing.T) {
		var mc *mock.Call

		mc = articleData.On("SelectArticles", mock.Anything)
		mc = mc.Return(nil, errors.New("error"))
		mc.Once()

		_, err, _ := articleBusiness.FindArticles(articles.QueryParams{})

		assert.NotNil(t, err)

	})

	t.Run("valid - return error when FindUserById failed", func(t *testing.T) {
		var mc *mock.Call

		mc = articleData.On("SelectArticles", mock.Anything)
		mc = mc.Return([]articles.ArticleCore{articleValue}, nil)
		mc.Once()

		mc = userBusiness.On("FindUserById", mock.AnythingOfType("uint"))
		mc = mc.Return(userValue, nil, 200)
		mc.Once()

		mc = reactionBusiness.On("CountTotalArticleLikes", mock.AnythingOfType("uint"))
		mc = mc.Return(0, errors.New("error"))
		mc.Once()

		_, err, _ := articleBusiness.FindArticles(articles.QueryParams{})

		assert.NotNil(t, err)
	})

	t.Run("valid - return error when CountTotalArticleLikes failed", func(t *testing.T) {
		var mc *mock.Call

		mc = articleData.On("SelectArticles", mock.Anything)
		mc = mc.Return([]articles.ArticleCore{articleValue}, nil)
		mc.Once()

		mc = userBusiness.On("FindUserById", mock.AnythingOfType("uint"))
		mc = mc.Return(users.UserCore{}, errors.New("error"), 500)
		mc.Once()

		_, err, _ := articleBusiness.FindArticles(articles.QueryParams{})

		assert.NotNil(t, err)
	})
}

func TestFindArticleById(t *testing.T) {
	t.Run("valid - find detail article", func(t *testing.T) {
		articleData.On("SelectArticleById", mock.AnythingOfType("uint")).Return(articleValue, nil).Once()
		userBusiness.On("FindUserById", mock.AnythingOfType("uint")).Return(userValue, nil, 200).Once()

		reactionBusiness.On("CountTotalArticleLikes", mock.AnythingOfType("uint")).Return(1, nil).Once()

		data, _, _ := articleBusiness.FindArticleById(1)

		assert.Equal(t, articleValue.ID, data.ID)
	})

	t.Run("valid - error when SelectArticleById", func(t *testing.T) {
		articleData.On("SelectArticleById", mock.AnythingOfType("uint")).Return(articles.ArticleCore{}, errors.New("error")).Once()

		_, err, _ := articleBusiness.FindArticleById(1)

		assert.NotNil(t, err)
	})

	t.Run("valid - error when articleData is not found", func(t *testing.T) {
		articleData.On("SelectArticleById", mock.AnythingOfType("uint")).Return(articles.ArticleCore{}, nil).Once()

		_, err, _ := articleBusiness.FindArticleById(1)

		assert.NotNil(t, err)
	})

	t.Run("valid - error when FindUserById failed", func(t *testing.T) {
		articleData.On("SelectArticleById", mock.AnythingOfType("uint")).Return(articleValue, nil).Once()
		userBusiness.On("FindUserById", mock.AnythingOfType("uint")).Return(users.UserCore{}, errors.New("error"), 500).Once()

		_, err, _ := articleBusiness.FindArticleById(1)

		assert.NotNil(t, err)
	})

	t.Run("valid - find detail article", func(t *testing.T) {
		articleData.On("SelectArticleById", mock.AnythingOfType("uint")).Return(articleValue, nil).Once()
		userBusiness.On("FindUserById", mock.AnythingOfType("uint")).Return(userValue, nil, 200).Once()

		reactionBusiness.On("CountTotalArticleLikes", mock.AnythingOfType("uint")).Return(0, errors.New("abc")).Once()

		_, err, _ := articleBusiness.FindArticleById(1)

		assert.NotNil(t, err)
	})
}

func TestFindTranslatedArticleById(t *testing.T) {
	t.Run("valid - find translated detail article", func(t *testing.T) {
		var mc *mock.Call

		mc = articleData.On("SelectArticleById", mock.AnythingOfType("uint"))
		mc = mc.Return(articleValue, nil)
		mc.Once()

		mc = userBusiness.On("FindUserById", mock.AnythingOfType("uint"))
		mc = mc.Return(userValue, nil, 200)
		mc.Once()

		mc = translateBusiness.On("Translate", mock.Anything)
		mc = mc.Return("translated title", nil)
		mc.Once()

		mc = translateBusiness.On("Translate", mock.Anything)
		mc = mc.Return("translated subtitle", nil)
		mc.Once()

		data, _, status := articleBusiness.FindTranslatedArticleById(1, "en")

		assert.Equal(t, http.StatusOK, status)
		assert.Equal(t, "translated title", data.Title)
		assert.Equal(t, "translated subtitle", data.Subtitle)
	})

	t.Run("valid - error when SelectArticleById failed", func(t *testing.T) {
		var mc *mock.Call

		mc = articleData.On("SelectArticleById", mock.AnythingOfType("uint"))
		mc = mc.Return(articles.ArticleCore{}, errors.New("error"))
		mc.Once()

		_, err, status := articleBusiness.FindTranslatedArticleById(1, "en")

		assert.Equal(t, http.StatusInternalServerError, status)
		assert.NotNil(t, err)
	})

	t.Run("valid - error when SelectArticleById return 0 records", func(t *testing.T) {
		var mc *mock.Call

		mc = articleData.On("SelectArticleById", mock.AnythingOfType("uint"))
		mc = mc.Return(articles.ArticleCore{}, nil)
		mc.Once()

		_, err, status := articleBusiness.FindTranslatedArticleById(1, "en")

		assert.Equal(t, http.StatusNotFound, status)
		assert.NotNil(t, err)
	})

	t.Run("valid - error when FindUserById failed", func(t *testing.T) {
		var mc *mock.Call

		mc = articleData.On("SelectArticleById", mock.AnythingOfType("uint"))
		mc = mc.Return(articleValue, nil)
		mc.Once()

		mc = userBusiness.On("FindUserById", mock.AnythingOfType("uint"))
		mc = mc.Return(users.UserCore{}, errors.New("error"), 500)
		mc.Once()

		_, err, status := articleBusiness.FindTranslatedArticleById(1, "en")

		assert.Equal(t, http.StatusInternalServerError, status)
		assert.NotNil(t, err)
	})

	t.Run("valid - error when translate title", func(t *testing.T) {
		var mc *mock.Call

		mc = articleData.On("SelectArticleById", mock.AnythingOfType("uint"))
		mc = mc.Return(articleValue, nil)
		mc.Once()

		mc = userBusiness.On("FindUserById", mock.AnythingOfType("uint"))
		mc = mc.Return(userValue, nil, 200)
		mc.Once()

		mc = translateBusiness.On("Translate", mock.Anything)
		mc = mc.Return("", errors.New("error"))
		mc.Once()

		_, err, status := articleBusiness.FindTranslatedArticleById(1, "en")

		assert.Equal(t, http.StatusInternalServerError, status)
		assert.NotNil(t, err)
	})

	t.Run("valid - error when translate title", func(t *testing.T) {
		var mc *mock.Call

		mc = articleData.On("SelectArticleById", mock.AnythingOfType("uint"))
		mc = mc.Return(articleValue, nil)
		mc.Once()

		mc = userBusiness.On("FindUserById", mock.AnythingOfType("uint"))
		mc = mc.Return(userValue, nil, 200)
		mc.Once()

		mc = translateBusiness.On("Translate", mock.Anything)
		mc = mc.Return("translate title", nil)
		mc.Once()

		mc = translateBusiness.On("Translate", mock.Anything)
		mc = mc.Return("", errors.New("error"))
		mc.Once()

		_, err, status := articleBusiness.FindTranslatedArticleById(1, "en")

		assert.Equal(t, http.StatusInternalServerError, status)
		assert.NotNil(t, err)
	})
}

func TestFindUserArticles(t *testing.T) {
	t.Run("valid - FindUserArticles", func(t *testing.T) {
		var mc *mock.Call

		mc = userBusiness.On("FindUserByUsername", mock.AnythingOfType("string"))
		mc = mc.Return(userValue, nil, 200)
		mc.Once()

		mc = articleData.On("SelectArticlesByAuthorId", mock.AnythingOfType("uint"))
		mc = mc.Return([]articles.ArticleCore{articleValue}, nil)
		mc.Once()

		reactionBusiness.On("CountTotalArticleLikes", mock.AnythingOfType("uint")).Return(1, nil).Once()

		userArticles, _, _ := articleBusiness.FindUserArticles("riza.dwii")

		assert.Equal(t, 1, len(userArticles))
	})

	t.Run("valid - error when SelectArticlesByAuthorId failed", func(t *testing.T) {
		var mc *mock.Call

		mc = userBusiness.On("FindUserByUsername", mock.AnythingOfType("string"))
		mc = mc.Return(users.UserCore{}, errors.New("error"), 500)
		mc.Once()

		_, err, status := articleBusiness.FindUserArticles("owo")

		assert.NotNil(t, err)
		assert.Equal(t, http.StatusInternalServerError, status)
	})

	t.Run("valid - when failed SelectArticlesByAuthorId", func(t *testing.T) {
		var mc *mock.Call

		mc = userBusiness.On("FindUserByUsername", mock.AnythingOfType("string"))
		mc = mc.Return(userValue, nil, 200)
		mc.Once()

		mc = articleData.On("SelectArticlesByAuthorId", mock.AnythingOfType("uint"))
		mc = mc.Return(nil, errors.New("error"))
		mc.Once()

		_, err, status := articleBusiness.FindUserArticles("owo")

		assert.NotNil(t, err)
		assert.Equal(t, http.StatusInternalServerError, status)
	})

	t.Run("valid - when CountTotalArticleLikes failed", func(t *testing.T) {
		var mc *mock.Call

		mc = userBusiness.On("FindUserByUsername", mock.AnythingOfType("string"))
		mc = mc.Return(userValue, nil, 200)
		mc.Once()

		mc = articleData.On("SelectArticlesByAuthorId", mock.AnythingOfType("uint"))
		mc = mc.Return([]articles.ArticleCore{articleValue}, nil)
		mc.Once()

		reactionBusiness.On("CountTotalArticleLikes", mock.AnythingOfType("uint")).Return(0, errors.New("err")).Once()

		_, err, status := articleBusiness.FindUserArticles("riza.dwii")

		assert.NotNil(t, err)
		assert.Equal(t, http.StatusInternalServerError, status)
	})

}

func TestRemoveArticleById(t *testing.T) {
	t.Run("valid - RemoveArticleById", func(t *testing.T) {
		var mc *mock.Call

		mc = articleData.On("DeleteArticleTags", mock.AnythingOfType("uint"))
		mc = mc.Return(nil)
		mc.Once()

		mc = articleData.On("DeleteArticleById", mock.AnythingOfType("uint"))
		mc = mc.Return(nil)
		mc.Once()

		mc = reactionBusiness.On("RemoveCommentsByArticleId", mock.AnythingOfType("uint"))
		mc = mc.Return(nil)
		mc.Once()

		mc = reactionBusiness.On("RemoveLikesByArticleId", mock.AnythingOfType("uint"))
		mc = mc.Return(nil)
		mc.Once()

		mc = reactionBusiness.On("RemoveReportsByArticleId", mock.AnythingOfType("uint"))
		mc = mc.Return(nil)
		mc.Once()

		err, status := articleBusiness.RemoveArticleById(1)

		assert.Nil(t, err)
		assert.Equal(t, http.StatusAccepted, status)
	})

	t.Run("valid - RemoveArticleById", func(t *testing.T) {
		var mc *mock.Call

		mc = articleData.On("DeleteArticleTags", mock.AnythingOfType("uint"))
		mc = mc.Return(errors.New("error"))
		mc.Once()

		err, status := articleBusiness.RemoveArticleById(1)

		assert.NotNil(t, err)
		assert.Equal(t, http.StatusInternalServerError, status)
	})

	t.Run("valid - RemoveArticleById", func(t *testing.T) {
		var mc *mock.Call

		mc = articleData.On("DeleteArticleTags", mock.AnythingOfType("uint"))
		mc = mc.Return(nil)
		mc.Once()

		mc = articleData.On("DeleteArticleById", mock.AnythingOfType("uint"))
		mc = mc.Return(errors.New("error"))
		mc.Once()

		err, status := articleBusiness.RemoveArticleById(1)

		assert.NotNil(t, err)
		assert.Equal(t, http.StatusInternalServerError, status)
	})

}

func TestRemoveUserArticles(t *testing.T) {
	t.Run("valid - RemoveUserArticles", func(t *testing.T) {
		var mc *mock.Call

		mc = articleData.On("SelectArticlesByAuthorId", mock.AnythingOfType("uint"))
		mc = mc.Return([]articles.ArticleCore{articleValue}, nil)
		mc.Once()

		mc = articleData.On("DeleteTagByArticleIds", mock.AnythingOfType("[]uint"))
		mc = mc.Return(nil)
		mc.Once()

		mc = articleData.On("DeleteArticlesByUserId", mock.AnythingOfType("uint"))
		mc = mc.Return(nil)
		mc.Once()

		err, status := articleBusiness.RemoveUserArticles(1)

		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, status)
	})

	t.Run("valid - error when SelectArticlesByAuthorId failed", func(t *testing.T) {
		var mc *mock.Call

		mc = articleData.On("SelectArticlesByAuthorId", mock.AnythingOfType("uint"))
		mc = mc.Return(nil, errors.New("error"))
		mc.Once()

		err, status := articleBusiness.RemoveUserArticles(1)

		assert.NotNil(t, err)
		assert.Equal(t, http.StatusInternalServerError, status)
	})

	t.Run("valid - DeleteTagByArticleIds failed", func(t *testing.T) {
		var mc *mock.Call

		mc = articleData.On("SelectArticlesByAuthorId", mock.AnythingOfType("uint"))
		mc = mc.Return([]articles.ArticleCore{articleValue}, nil)
		mc.Once()

		mc = articleData.On("DeleteTagByArticleIds", mock.AnythingOfType("[]uint"))
		mc = mc.Return(errors.New("error"))
		mc.Once()

		err, status := articleBusiness.RemoveUserArticles(1)

		assert.Error(t, err)
		assert.Equal(t, http.StatusInternalServerError, status)
	})

	t.Run("valid - error when DeleteArticlesByUserId", func(t *testing.T) {
		var mc *mock.Call

		mc = articleData.On("SelectArticlesByAuthorId", mock.AnythingOfType("uint"))
		mc = mc.Return([]articles.ArticleCore{articleValue}, nil)
		mc.Once()

		mc = articleData.On("DeleteTagByArticleIds", mock.AnythingOfType("[]uint"))
		mc = mc.Return(nil)
		mc.Once()

		mc = articleData.On("DeleteArticlesByUserId", mock.AnythingOfType("uint"))
		mc = mc.Return(errors.New("error"))
		mc.Once()

		err, status := articleBusiness.RemoveUserArticles(1)

		assert.Error(t, err)
		assert.Equal(t, http.StatusInternalServerError, status)
	})
}

func TestCreateArticle(t *testing.T) {
	t.Run("valid - CreateArticle", func(t *testing.T) {
		var mc *mock.Call

		mc = articleData.On("InsertArticle", mock.Anything)
		mc = mc.Return(articleValue, nil)
		mc.Once()

		article, err, status := articleBusiness.CreateArticle(articleValue)

		assert.Equal(t, articleValue, article)
		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, status)
	})

	t.Run("valid - CreateArticle", func(t *testing.T) {
		var mc *mock.Call

		mc = articleData.On("InsertArticle", mock.Anything)
		mc = mc.Return(articles.ArticleCore{}, errors.New("error"))
		mc.Once()

		_, err, status := articleBusiness.CreateArticle(articleValue)

		assert.NotNil(t, err)
		assert.Equal(t, http.StatusInternalServerError, status)
	})
}
