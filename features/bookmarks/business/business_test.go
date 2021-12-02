package business_test

import (
	"errors"
	"os"
	"testing"

	"github.com/rizadwiandhika/miniproject-backend-alterra/features/articles"
	amock "github.com/rizadwiandhika/miniproject-backend-alterra/features/articles/mocks"
	"github.com/rizadwiandhika/miniproject-backend-alterra/features/bookmarks"
	bb "github.com/rizadwiandhika/miniproject-backend-alterra/features/bookmarks/business"
	bmock "github.com/rizadwiandhika/miniproject-backend-alterra/features/bookmarks/mocks"
	"github.com/rizadwiandhika/miniproject-backend-alterra/features/users"
	umock "github.com/rizadwiandhika/miniproject-backend-alterra/features/users/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	bookmarkBusiness bookmarks.IBusiness
	bookmarkData     bmock.IData
	userBusiness     umock.IBusiness
	articleBusiness  amock.IBusiness

	userValue     users.UserCore
	articleValue  articles.ArticleCore
	bookmarkValue bookmarks.BookmarkCore
)

func TestMain(m *testing.M) {
	bookmarkBusiness = bb.NewBusiness(&bookmarkData, &userBusiness, &articleBusiness)

	userValue = users.UserCore{
		ID:       1,
		Username: "abc",
		Name:     "K Haidar",
	}

	articleValue = articles.ArticleCore{
		ID:    1,
		Title: "Title",
	}

	bookmarkValue = bookmarks.BookmarkCore{
		ID:        1,
		UserID:    1,
		ArticleID: 1,
	}

	os.Exit(m.Run())
}

func TestCreateBookmark(t *testing.T) {
	t.Run("valid - CreateBookmark", func(t *testing.T) {
		var mc *mock.Call

		mc = userBusiness.On("FindUserByUsername", mock.Anything)
		mc = mc.Return(userValue, nil, 123)
		mc.Once()

		mc = articleBusiness.On("FindArticleById", mock.Anything)
		mc = mc.Return(articleValue, nil, 123)
		mc.Once()

		mc = bookmarkData.On("SelectDetailUserBookmark", mock.Anything, mock.Anything)
		mc = mc.Return(bookmarks.BookmarkCore{}, nil)
		mc.Once()

		mc = bookmarkData.On("InsertBookmark", mock.Anything)
		mc = mc.Return(nil)
		mc.Once()

		err := bookmarkBusiness.CreateBookmark(userValue.Username, articleValue.ID)

		assert.NoError(t, err)
	})

	t.Run("valid - when FindUserByUsername failed", func(t *testing.T) {
		var mc *mock.Call

		mc = userBusiness.On("FindUserByUsername", mock.Anything)
		mc = mc.Return(users.UserCore{}, errors.New("err"), 123)
		mc.Once()

		err := bookmarkBusiness.CreateBookmark(userValue.Username, articleValue.ID)

		assert.Error(t, err)
	})

	t.Run("valid - when FindArticleById failed", func(t *testing.T) {
		var mc *mock.Call

		mc = userBusiness.On("FindUserByUsername", mock.Anything)
		mc = mc.Return(userValue, nil, 123)
		mc.Once()

		mc = articleBusiness.On("FindArticleById", mock.Anything)
		mc = mc.Return(articles.ArticleCore{}, errors.New("err"), 123)
		mc.Once()

		err := bookmarkBusiness.CreateBookmark(userValue.Username, articleValue.ID)

		assert.Error(t, err)
	})

	t.Run("valid - when SelectDetailUserBookmark failed", func(t *testing.T) {
		var mc *mock.Call

		mc = userBusiness.On("FindUserByUsername", mock.Anything)
		mc = mc.Return(userValue, nil, 123)
		mc.Once()

		mc = articleBusiness.On("FindArticleById", mock.Anything)
		mc = mc.Return(articleValue, nil, 123)
		mc.Once()

		mc = bookmarkData.On("SelectDetailUserBookmark", mock.Anything, mock.Anything)
		mc = mc.Return(bookmarks.BookmarkCore{}, errors.New("err"))
		mc.Once()

		err := bookmarkBusiness.CreateBookmark(userValue.Username, articleValue.ID)

		assert.Error(t, err)
	})

	t.Run("valid - when bookmark already exists", func(t *testing.T) {
		var mc *mock.Call

		mc = userBusiness.On("FindUserByUsername", mock.Anything)
		mc = mc.Return(userValue, nil, 123)
		mc.Once()

		mc = articleBusiness.On("FindArticleById", mock.Anything)
		mc = mc.Return(articleValue, nil, 123)
		mc.Once()

		mc = bookmarkData.On("SelectDetailUserBookmark", mock.Anything, mock.Anything)
		mc = mc.Return(bookmarkValue, nil)
		mc.Once()

		err := bookmarkBusiness.CreateBookmark(userValue.Username, articleValue.ID)

		assert.Error(t, err)
	})

}

func TestFindUserBookmarks(t *testing.T) {
	t.Run("valid - FindUserBookmarks", func(t *testing.T) {
		var mc *mock.Call

		mc = userBusiness.On("FindUserByUsername", mock.Anything)
		mc = mc.Return(userValue, nil, 123)
		mc.Once()

		mc = bookmarkData.On("SelectUserBookmarks", mock.Anything)
		mc = mc.Return([]bookmarks.BookmarkCore{bookmarkValue}, nil)
		mc.Once()

		mc = articleBusiness.On("FindArticleById", mock.Anything)
		mc = mc.Return(articleValue, nil, 123)
		mc.Once()

		books, err := bookmarkBusiness.FindUserBookmarks(userValue.Username)

		assert.Greater(t, len(books), 0)
		assert.NoError(t, err)
	})

	t.Run("valid - when FindUserByUsername failed", func(t *testing.T) {
		var mc *mock.Call

		mc = userBusiness.On("FindUserByUsername", mock.Anything)
		mc = mc.Return(users.UserCore{}, errors.New("err"), 123)
		mc.Once()

		books, err := bookmarkBusiness.FindUserBookmarks(userValue.Username)

		assert.Equal(t, 0, len(books))
		assert.Error(t, err)
	})

	t.Run("valid - when SelectDetailUserBookmarks failed", func(t *testing.T) {
		var mc *mock.Call

		mc = userBusiness.On("FindUserByUsername", mock.Anything)
		mc = mc.Return(userValue, nil, 123)
		mc.Once()

		mc = bookmarkData.On("SelectUserBookmarks", mock.Anything)
		mc = mc.Return([]bookmarks.BookmarkCore{}, errors.New("err"))
		mc.Once()

		books, err := bookmarkBusiness.FindUserBookmarks(userValue.Username)

		assert.Equal(t, 0, len(books))
		assert.Error(t, err)
	})

	t.Run("valid - FindUserBookmarks", func(t *testing.T) {
		var mc *mock.Call

		mc = userBusiness.On("FindUserByUsername", mock.Anything)
		mc = mc.Return(userValue, nil, 123)
		mc.Once()

		mc = bookmarkData.On("SelectUserBookmarks", mock.Anything)
		mc = mc.Return([]bookmarks.BookmarkCore{bookmarkValue}, nil)
		mc.Once()

		mc = articleBusiness.On("FindArticleById", mock.Anything)
		mc = mc.Return(articles.ArticleCore{}, errors.New("err"), 123)
		mc.Once()

		books, err := bookmarkBusiness.FindUserBookmarks(userValue.Username)

		assert.Equal(t, len(books), 0)
		assert.Error(t, err)
	})
}
