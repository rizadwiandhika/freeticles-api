package business_test

import (
	"errors"
	"net/http"
	"os"
	"testing"

	"github.com/rizadwiandhika/miniproject-backend-alterra/features/articles"
	reactions "github.com/rizadwiandhika/miniproject-backend-alterra/features/reactions"
	rb "github.com/rizadwiandhika/miniproject-backend-alterra/features/reactions/business"
	"github.com/rizadwiandhika/miniproject-backend-alterra/features/users"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	amocks "github.com/rizadwiandhika/miniproject-backend-alterra/features/articles/mocks"
	rmocks "github.com/rizadwiandhika/miniproject-backend-alterra/features/reactions/mocks"
	umocks "github.com/rizadwiandhika/miniproject-backend-alterra/features/users/mocks"
	iamocks "github.com/rizadwiandhika/miniproject-backend-alterra/third-parties/image/mocks"
)

var (
	reactionBusiness reactions.IBusiness

	reactionData    rmocks.IData
	userBusiness    umocks.IBusiness
	articleBusiness amocks.IBusiness
	imageAnalyzer   iamocks.IBusiness

	u1 users.UserCore
	u2 users.UserCore

	a1 articles.ArticleCore

	c1 reactions.CommentCore
	c2 reactions.CommentCore
)

func TestMain(m *testing.M) {
	reactionBusiness = rb.NewBusiness(&reactionData, &imageAnalyzer, &userBusiness, &articleBusiness)

	u1 = users.UserCore{
		ID:       1,
		Username: "riza.dwii",
	}
	u2 = users.UserCore{
		ID:       2,
		Username: "hernowo",
	}

	a1 = articles.ArticleCore{
		ID: 1,
	}

	c1 = reactions.CommentCore{
		ID:        1,
		UserID:    1,
		ArticleID: 1,
		Commentar: "wow",
	}
	c2 = reactions.CommentCore{
		ID:        2,
		UserID:    2,
		ArticleID: 1,
		Commentar: "wow",
	}

	os.Exit(m.Run())
}

func TestFindCommentsByArticleId(t *testing.T) {
	t.Run("valid - FindCommentsByArticleId", func(t *testing.T) {
		reactionData.On("SelectCommentsByArticleId", mock.AnythingOfType("uint")).Return([]reactions.CommentCore{c1, c2}, nil).Once()

		userBusiness.On("FindUsersByIds", mock.AnythingOfType("[]uint")).Return([]users.UserCore{u1, u2}, nil, 123).Once()

		comments, err, status := reactionBusiness.FindCommentsByArticleId(1)

		assert.GreaterOrEqual(t, len(comments), 0)
		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, status)
	})

	t.Run("valid - when SelectCommentsByArticleId failed", func(t *testing.T) {
		reactionData.On("SelectCommentsByArticleId", mock.AnythingOfType("uint")).Return(nil, errors.New("any error")).Once()

		comments, err, status := reactionBusiness.FindCommentsByArticleId(1)

		assert.Nil(t, comments)
		assert.Error(t, err)
		assert.Equal(t, http.StatusInternalServerError, status)
	})

	t.Run("valid - when FindUsersByIds failed", func(t *testing.T) {
		reactionData.On("SelectCommentsByArticleId", mock.AnythingOfType("uint")).Return([]reactions.CommentCore{c1, c2}, nil).Once()

		userBusiness.On("FindUsersByIds", mock.AnythingOfType("[]uint")).Return(nil, errors.New("abc"), 123).Once()

		comments, err, status := reactionBusiness.FindCommentsByArticleId(1)

		assert.Nil(t, comments)
		assert.Error(t, err)
		assert.Equal(t, http.StatusInternalServerError, status)
	})
}

func TestCountTotalArticleLikes(t *testing.T) {
	t.Run("valid - CountTotalArticleLikes", func(t *testing.T) {
		reactionData.On("SelectCountLikes", mock.AnythingOfType("uint")).Return(1, nil).Once()

		count, err := reactionBusiness.CountTotalArticleLikes(1)

		assert.Equal(t, 1, count)
		assert.Nil(t, err)
	})

	t.Run("valid - when CountTotalArticleLikes failed", func(t *testing.T) {
		reactionData.On("SelectCountLikes", mock.AnythingOfType("uint")).Return(0, errors.New("any error")).Once()

		count, err := reactionBusiness.CountTotalArticleLikes(1)

		assert.Equal(t, 0, count)
		assert.Error(t, err)
	})
}

func TestPostLike(t *testing.T) {
	t.Run("valid - PostLike", func(t *testing.T) {
		userBusiness.On("FindUserByUsername", mock.AnythingOfType("string")).Return(u1, nil, 123).Once()
		articleBusiness.On("FindArticleById", mock.AnythingOfType("uint")).Return(a1, nil, 123).Once()

		reactionData.On("SelectLike", mock.Anything).Return(reactions.LikeCore{}, nil).Once()

		reactionData.On("InsertLike", mock.Anything).Return(nil).Once()

		err, status := reactionBusiness.PostLike("riza.dwii", 1)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusCreated, status)
	})

	t.Run("valid - when user or article not found", func(t *testing.T) {
		userBusiness.On("FindUserByUsername", mock.AnythingOfType("string")).Return(users.UserCore{}, errors.New("abc"), 123).Once()
		articleBusiness.On("FindArticleById", mock.AnythingOfType("uint")).Return(articles.ArticleCore{}, errors.New("abc"), 123).Once()

		err, status := reactionBusiness.PostLike("riza.dwii", 1)

		assert.Error(t, err)
		assert.Equal(t, http.StatusNotFound, status)
	})

	t.Run("valid - when SelectLike failed", func(t *testing.T) {
		userBusiness.On("FindUserByUsername", mock.AnythingOfType("string")).Return(u1, nil, 123).Once()
		articleBusiness.On("FindArticleById", mock.AnythingOfType("uint")).Return(a1, nil, 123).Once()

		reactionData.On("SelectLike", mock.AnythingOfType("reactions.LikeCore")).Return(reactions.LikeCore{}, errors.New("abc")).Once()

		err, status := reactionBusiness.PostLike(u1.Username, a1.ID)

		assert.Error(t, err)
		assert.Equal(t, http.StatusInternalServerError, status)
	})

	t.Run("valid - when user already like the article", func(t *testing.T) {
		userBusiness.On("FindUserByUsername", mock.AnythingOfType("string")).Return(u1, nil, 123).Once()
		articleBusiness.On("FindArticleById", mock.AnythingOfType("uint")).Return(a1, nil, 123).Once()

		reactionData.On("SelectLike", mock.AnythingOfType("reactions.LikeCore")).Return(reactions.LikeCore{ID: 1}, nil).Once()

		err, status := reactionBusiness.PostLike(u1.Username, a1.ID)

		assert.Error(t, err)
		assert.Equal(t, http.StatusUnprocessableEntity, status)
	})

	t.Run("valid - when InsertLike failed", func(t *testing.T) {
		userBusiness.On("FindUserByUsername", mock.AnythingOfType("string")).Return(u1, nil, 123).Once()
		articleBusiness.On("FindArticleById", mock.AnythingOfType("uint")).Return(a1, nil, 123).Once()

		reactionData.On("SelectLike", mock.Anything).Return(reactions.LikeCore{}, nil).Once()

		reactionData.On("InsertLike", mock.Anything).Return(errors.New("err")).Once()

		err, status := reactionBusiness.PostLike("riza.dwii", 1)

		assert.Error(t, err)
		assert.Equal(t, http.StatusInternalServerError, status)
	})
}

func TestPostComment(t *testing.T) {
	t.Run("valid - PostComment", func(t *testing.T) {
		userBusiness.On("FindUserByUsername", mock.AnythingOfType("string")).Return(u1, nil, 123).Once()
		articleBusiness.On("FindArticleById", mock.AnythingOfType("uint")).Return(a1, nil, 123).Once()

		reactionData.On("InsertComment", mock.Anything).Return(nil).Once()

		err, status := reactionBusiness.PostComment("riza.dwii", 1, "comment")

		assert.NoError(t, err)
		assert.Equal(t, http.StatusCreated, status)
	})

	t.Run("valid - when user or article not found", func(t *testing.T) {
		userBusiness.On("FindUserByUsername", mock.AnythingOfType("string")).Return(users.UserCore{}, errors.New("abc"), 123).Once()
		articleBusiness.On("FindArticleById", mock.AnythingOfType("uint")).Return(articles.ArticleCore{}, errors.New("abc"), 123).Once()

		err, status := reactionBusiness.PostComment("riza.dwii", 1, "comment")

		assert.Error(t, err)
		assert.Equal(t, http.StatusNotFound, status)
	})

	t.Run("valid - when InsertComment failed", func(t *testing.T) {
		userBusiness.On("FindUserByUsername", mock.AnythingOfType("string")).Return(u1, nil, 123).Once()
		articleBusiness.On("FindArticleById", mock.AnythingOfType("uint")).Return(a1, nil, 123).Once()

		reactionData.On("InsertComment", mock.AnythingOfType("reactions.CommentCore")).Return(errors.New("abc")).Once()

		err, status := reactionBusiness.PostComment("riza.dwii", 1, "comment")
		assert.Error(t, err)
		assert.Equal(t, http.StatusInternalServerError, status)
	})
}

func TestDislike(t *testing.T) {
	t.Run("valid - Dislike", func(t *testing.T) {
		userBusiness.On("FindUserByUsername", mock.AnythingOfType("string")).Return(u1, nil, 123).Once()
		articleBusiness.On("FindArticleById", mock.AnythingOfType("uint")).Return(a1, nil, 123).Once()

		reactionData.On("SelectLike", mock.Anything).Return(reactions.LikeCore{ID: 1}, nil).Once()

		reactionData.On("DeleteLikeById", mock.AnythingOfType("uint")).Return(nil).Once()

		err, status := reactionBusiness.Dislike("riza.dwii", 1)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, status)
	})

	t.Run("valid - when SelectLike failed", func(t *testing.T) {
		userBusiness.On("FindUserByUsername", mock.AnythingOfType("string")).Return(u1, nil, 123).Once()
		articleBusiness.On("FindArticleById", mock.AnythingOfType("uint")).Return(a1, nil, 123).Once()

		reactionData.On("SelectLike", mock.AnythingOfType("reactions.LikeCore")).Return(reactions.LikeCore{}, errors.New("abc")).Once()

		err, status := reactionBusiness.Dislike("riza.dwii", 1)

		assert.Error(t, err)
		assert.Equal(t, http.StatusInternalServerError, status)
	})

	t.Run("valid - when indeed user didn't like the article", func(t *testing.T) {
		userBusiness.On("FindUserByUsername", mock.AnythingOfType("string")).Return(u1, nil, 123).Once()
		articleBusiness.On("FindArticleById", mock.AnythingOfType("uint")).Return(a1, nil, 123).Once()

		reactionData.On("SelectLike", mock.AnythingOfType("reactions.LikeCore")).Return(reactions.LikeCore{ID: 0}, nil).Once()

		err, status := reactionBusiness.Dislike("riza.dwii", 1)

		assert.Error(t, err)
		assert.Equal(t, http.StatusUnprocessableEntity, status)
	})

	t.Run("valid - when DeleteLikeById failed", func(t *testing.T) {
		userBusiness.On("FindUserByUsername", mock.AnythingOfType("string")).Return(u1, nil, 123).Once()
		articleBusiness.On("FindArticleById", mock.AnythingOfType("uint")).Return(a1, nil, 123).Once()

		reactionData.On("SelectLike", mock.AnythingOfType("reactions.LikeCore")).Return(reactions.LikeCore{ID: 1}, nil).Once()

		reactionData.On("DeleteLikeById", mock.AnythingOfType("uint")).Return(errors.New("err")).Once()

		err, status := reactionBusiness.Dislike("riza.dwii", 1)

		assert.Error(t, err)
		assert.Equal(t, http.StatusInternalServerError, status)
	})
}

func TestReportArticle(t *testing.T) {
	t.Run("valid - ReportArticle (not sexual content)", func(t *testing.T) {
		userBusiness.On("FindUserByUsername", mock.AnythingOfType("string")).Return(u1, nil, 123).Once()
		articleBusiness.On("FindArticleById", mock.AnythingOfType("uint")).Return(a1, nil, 123).Once()

		reactionData.On("SelectUserReport", mock.AnythingOfType("uint"), mock.AnythingOfType("uint")).Return(reactions.ReportCore{}, nil).Once()

		reactionData.On("InsertReport", mock.AnythingOfType("reactions.ReportCore")).Return(nil).Once()

		err, status := reactionBusiness.ReportArticle("riza.dwii", 1, 1)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusCreated, status)
	})

	t.Run("valid - ReportArticle sexsual content", func(t *testing.T) {
		userBusiness.On("FindUserByUsername", mock.AnythingOfType("string")).Return(u1, nil, 123).Once()
		articleBusiness.On("FindArticleById", mock.AnythingOfType("uint")).Return(a1, nil, 123).Once()

		reactionData.On("SelectUserReport", mock.AnythingOfType("uint"), mock.AnythingOfType("uint")).Return(reactions.ReportCore{}, nil).Once()

		reactionData.On("InsertReport", mock.AnythingOfType("reactions.ReportCore")).Return(nil).Once()

		err, status := reactionBusiness.ReportArticle("riza.dwii", 1, 2)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusCreated, status)
	})

	t.Run("valid - when user or article not found", func(t *testing.T) {
		userBusiness.On("FindUserByUsername", mock.AnythingOfType("string")).Return(users.UserCore{}, errors.New("abc"), 123).Once()
		articleBusiness.On("FindArticleById", mock.AnythingOfType("uint")).Return(articles.ArticleCore{}, errors.New("abc"), 123).Once()

		err, status := reactionBusiness.ReportArticle("riza.dwii", 1, 1)

		assert.Error(t, err)
		assert.Equal(t, http.StatusNotFound, status)
	})

	t.Run("valid - when SelectUserReport failed", func(t *testing.T) {
		userBusiness.On("FindUserByUsername", mock.AnythingOfType("string")).Return(u1, nil, 123).Once()
		articleBusiness.On("FindArticleById", mock.AnythingOfType("uint")).Return(a1, nil, 123).Once()

		reactionData.On("SelectUserReport", mock.AnythingOfType("uint"), mock.AnythingOfType("uint")).Return(reactions.ReportCore{}, errors.New("abc")).Once()

		err, status := reactionBusiness.ReportArticle("riza.dwii", 1, 1)

		assert.Error(t, err)
		assert.Equal(t, http.StatusInternalServerError, status)
	})

	t.Run("valid - when article is already reported by that user", func(t *testing.T) {
		userBusiness.On("FindUserByUsername", mock.AnythingOfType("string")).Return(u1, nil, 123).Once()
		articleBusiness.On("FindArticleById", mock.AnythingOfType("uint")).Return(a1, nil, 123).Once()

		reactionData.On("SelectUserReport", mock.AnythingOfType("uint"), mock.AnythingOfType("uint")).Return(reactions.ReportCore{ID: 1}, nil).Once()

		err, status := reactionBusiness.ReportArticle("riza.dwii", 1, 1)

		assert.Error(t, err)
		assert.Equal(t, http.StatusUnprocessableEntity, status)
	})

	t.Run("valid - when InsertReport failed", func(t *testing.T) {
		userBusiness.On("FindUserByUsername", mock.AnythingOfType("string")).Return(u1, nil, 123).Once()
		articleBusiness.On("FindArticleById", mock.AnythingOfType("uint")).Return(a1, nil, 123).Once()

		reactionData.On("SelectUserReport", mock.AnythingOfType("uint"), mock.AnythingOfType("uint")).Return(reactions.ReportCore{}, nil).Once()

		reactionData.On("InsertReport", mock.AnythingOfType("reactions.ReportCore")).Return(errors.New("err")).Once()

		err, status := reactionBusiness.ReportArticle("riza.dwii", 1, 1)

		assert.Error(t, err)
		assert.Equal(t, http.StatusInternalServerError, status)
	})
}
