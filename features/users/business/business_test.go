package business_test

import (
	"errors"
	"net/http"
	"os"
	"testing"

	amock "github.com/rizadwiandhika/miniproject-backend-alterra/features/articles/mocks"
	bmock "github.com/rizadwiandhika/miniproject-backend-alterra/features/bookmarks/mocks"
	rmock "github.com/rizadwiandhika/miniproject-backend-alterra/features/reactions/mocks"
	"github.com/rizadwiandhika/miniproject-backend-alterra/features/users"
	ub "github.com/rizadwiandhika/miniproject-backend-alterra/features/users/business"
	umock "github.com/rizadwiandhika/miniproject-backend-alterra/features/users/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	userBusiness users.IBusiness

	userData         umock.IData
	articleBusiness  amock.IBusiness
	reactionBusiness rmock.IBusiness
	bookmarkBusiness bmock.IBusiness

	riza    users.UserCore
	hernowo users.UserCore
	jamie   users.UserCore

	rizaFollower1 users.FollowerCore
	rizaFollower2 users.FollowerCore
	// rizaFollower3 users.FollowerCore
)

func TestMain(m *testing.M) {
	userBusiness = ub.NewBusiness(&userData, &articleBusiness, &reactionBusiness, &bookmarkBusiness)

	riza = users.UserCore{
		ID:       1,
		Username: "riza",
		Email:    "riza@mail.com",
		Name:     "Riza Dwi Andhika",
		Role:     "admin",
	}
	hernowo = users.UserCore{
		ID:       2,
		Username: "hernowo",
		Email:    "hernowo@mail.com",
		Name:     "Hernowo Ari",
	}
	jamie = users.UserCore{
		ID:       3,
		Username: "jamie",
		Email:    "jamie@mail.com",
		Name:     "Jamie Saviola",
	}

	rizaFollower1 = users.FollowerCore{
		UserID:     1,
		FollowerID: 2,
	}
	rizaFollower2 = users.FollowerCore{
		UserID:     1,
		FollowerID: 3,
	}

	os.Exit(m.Run())
}

func TestFindUserById(t *testing.T) {
	t.Run("valid - FindUserById", func(t *testing.T) {
		userData.On("SelectUserById", mock.Anything).Return(riza, nil).Once()

		user, err, status := userBusiness.FindUserById(uint(1))

		assert.Equal(t, user.ID, riza.ID)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, status)
	})

	t.Run("valid - when SelectUserById failed", func(t *testing.T) {
		userData.On("SelectUserById", mock.Anything).Return(users.UserCore{}, errors.New("err")).Once()

		user, err, status := userBusiness.FindUserById(uint(1))

		assert.NotEqual(t, uint(1), user.ID)
		assert.Error(t, err)
		assert.Equal(t, http.StatusInternalServerError, status)
	})

	t.Run("valid - when user not found", func(t *testing.T) {
		userData.On("SelectUserById", mock.Anything).Return(users.UserCore{}, nil).Once()

		user, err, status := userBusiness.FindUserById(1)

		assert.NotEqual(t, uint(1), user.ID)
		assert.Error(t, err)
		assert.Equal(t, http.StatusNotFound, status)
	})
}

func TestFindUsersByIds(t *testing.T) {
	t.Run("valid FindUsersByIds", func(t *testing.T) {
		userData.On("SelectUsersByIds", mock.Anything).Return([]users.UserCore{riza}, nil).Once()

		users, err, status := userBusiness.FindUsersByIds([]uint{1})

		assert.Equal(t, users[0].ID, riza.ID)
		assert.Greater(t, len(users), 0)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, status)
	})

	t.Run("valid when FindUsersByIds failed", func(t *testing.T) {
		userData.On("SelectUsersByIds", mock.Anything).Return(nil, errors.New("err")).Once()

		users, err, status := userBusiness.FindUsersByIds([]uint{1})

		assert.Nil(t, users)
		assert.Error(t, err)
		assert.Equal(t, http.StatusInternalServerError, status)
	})
}

func TestFindUsers(t *testing.T) {
	t.Run("valid - FindUsers", func(t *testing.T) {
		userData.On("SelectUsers").Return([]users.UserCore{riza}, nil).Once()

		users, err, status := userBusiness.FindUsers()

		assert.Equal(t, users[0].ID, riza.ID)
		assert.Greater(t, len(users), 0)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, status)
	})

	t.Run("valid - when SelectUsers failed", func(t *testing.T) {
		userData.On("SelectUsers").Return(nil, errors.New("err")).Once()

		users, err, status := userBusiness.FindUsers()

		assert.Nil(t, users)
		assert.Error(t, err)
		assert.Equal(t, http.StatusInternalServerError, status)
	})
}

func TestFindUserFollowers(t *testing.T) {
	t.Run("valid - FindUserFollowers", func(t *testing.T) {
		userData.On("SelectUserByUsername", mock.AnythingOfType("string")).Return(riza, nil).Once()

		userData.On("SelectUserFollowers", mock.AnythingOfType("uint")).Return([]users.FollowerCore{rizaFollower1, rizaFollower2}, nil).Once()

		userData.On("SelectUsersByIds", mock.AnythingOfType("[]uint")).Return([]users.UserCore{hernowo, jamie}, nil).Once()

		dataRiza, err, status := userBusiness.FindUserFollowers("riza")

		assert.Equal(t, 2, len(dataRiza.Followers))
		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, status)
	})

	t.Run("valid - when SelectUserByUsername failed", func(t *testing.T) {
		userData.On("SelectUserByUsername", mock.AnythingOfType("string")).Return(users.UserCore{}, errors.New("abc")).Once()

		_, err, status := userBusiness.FindUserFollowers("riza")

		assert.Error(t, err)
		assert.Equal(t, http.StatusInternalServerError, status)
	})

	t.Run("valid - when user not found", func(t *testing.T) {
		userData.On("SelectUserByUsername", mock.AnythingOfType("string")).Return(users.UserCore{}, nil).Once()

		_, err, status := userBusiness.FindUserFollowers("riza")

		assert.Error(t, err)
		assert.Equal(t, http.StatusNotFound, status)
	})

	t.Run("valid - when SelectUserFollowers failed", func(t *testing.T) {
		userData.On("SelectUserByUsername", mock.AnythingOfType("string")).Return(riza, nil).Once()

		userData.On("SelectUserFollowers", mock.AnythingOfType("uint")).Return(nil, errors.New("err")).Once()

		_, err, status := userBusiness.FindUserFollowers("riza")

		assert.Error(t, err)
		assert.Equal(t, http.StatusInternalServerError, status)
	})

	t.Run("valid - when SelectUsersByIds failed", func(t *testing.T) {
		userData.On("SelectUserByUsername", mock.AnythingOfType("string")).Return(riza, nil).Once()

		userData.On("SelectUserFollowers", mock.AnythingOfType("uint")).Return([]users.FollowerCore{rizaFollower1, rizaFollower2}, nil).Once()

		userData.On("SelectUsersByIds", mock.AnythingOfType("[]uint")).Return(nil, errors.New("err")).Once()

		_, err, status := userBusiness.FindUserFollowers("riza")

		assert.Error(t, err)
		assert.Equal(t, http.StatusInternalServerError, status)
	})
}

func TestUserFollowings(t *testing.T) {
	t.Run("valid - FindUserFollowings", func(t *testing.T) {
		userData.On("SelectUserByUsername", mock.AnythingOfType("string")).Return(riza, nil).Once()

		userData.On("SelectUserFollowings", mock.AnythingOfType("uint")).Return([]users.FollowerCore{rizaFollower1, rizaFollower2}, nil).Once()

		userData.On("SelectUsersByIds", mock.AnythingOfType("[]uint")).Return([]users.UserCore{hernowo, jamie}, nil).Once()

		dataRiza, err, status := userBusiness.FindUserFollowings("riza")

		assert.Equal(t, 2, len(dataRiza.Followers))
		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, status)
	})

	t.Run("valid - when SelectUserByUsername failed", func(t *testing.T) {
		userData.On("SelectUserByUsername", mock.AnythingOfType("string")).Return(users.UserCore{}, errors.New("abc")).Once()

		_, err, status := userBusiness.FindUserFollowings("riza")

		assert.Error(t, err)
		assert.Equal(t, http.StatusInternalServerError, status)
	})

	t.Run("valid - when user not found", func(t *testing.T) {
		userData.On("SelectUserByUsername", mock.AnythingOfType("string")).Return(users.UserCore{}, nil).Once()

		_, err, status := userBusiness.FindUserFollowings("riza")

		assert.Error(t, err)
		assert.Equal(t, http.StatusNotFound, status)
	})

	t.Run("valid - when SelectUserFollowers failed", func(t *testing.T) {
		userData.On("SelectUserByUsername", mock.AnythingOfType("string")).Return(riza, nil).Once()

		userData.On("SelectUserFollowings", mock.AnythingOfType("uint")).Return(nil, errors.New("err")).Once()

		_, err, status := userBusiness.FindUserFollowings("riza")

		assert.Error(t, err)
		assert.Equal(t, http.StatusInternalServerError, status)
	})

	t.Run("valid - when SelectUsersByIds failed", func(t *testing.T) {
		userData.On("SelectUserByUsername", mock.AnythingOfType("string")).Return(riza, nil).Once()

		userData.On("SelectUserFollowings", mock.AnythingOfType("uint")).Return([]users.FollowerCore{rizaFollower1, rizaFollower2}, nil).Once()

		userData.On("SelectUsersByIds", mock.AnythingOfType("[]uint")).Return(nil, errors.New("err")).Once()

		_, err, status := userBusiness.FindUserFollowings("riza")

		assert.Error(t, err)
		assert.Equal(t, http.StatusInternalServerError, status)
	})
}

func TestFindUserByUsername(t *testing.T) {
	t.Run("valid - FindUserByUsername", func(t *testing.T) {
		userData.On("SelectUserByUsername", mock.Anything).Return(riza, nil).Once()

		user, err, status := userBusiness.FindUserByUsername(string(mock.AnythingOfType("string")))

		assert.Equal(t, user.ID, riza.ID)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, status)
	})

	t.Run("valid - when SelectUserByUsername failed", func(t *testing.T) {
		userData.On("SelectUserByUsername", mock.Anything).Return(users.UserCore{}, errors.New("err")).Once()

		user, err, status := userBusiness.FindUserByUsername(string(mock.AnythingOfType("string")))

		assert.NotEqual(t, uint(1), user.ID)
		assert.Error(t, err)
		assert.Equal(t, http.StatusInternalServerError, status)
	})

	t.Run("valid - when user not found", func(t *testing.T) {
		userData.On("SelectUserByUsername", mock.Anything).Return(users.UserCore{}, nil).Once()

		user, err, status := userBusiness.FindUserByUsername("riza.dwi")

		assert.NotEqual(t, uint(1), user.ID)
		assert.Error(t, err)
		assert.Equal(t, http.StatusNotFound, status)
	})
}

func TestFindUserByEmail(t *testing.T) {
	t.Run("valid - FindUserByEmail", func(t *testing.T) {
		userData.On("SelectUserByEmail", mock.Anything).Return(riza, nil).Once()

		user, err, status := userBusiness.FindUserByEmail(string(mock.AnythingOfType("string")))

		assert.Equal(t, user.ID, riza.ID)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, status)
	})

	t.Run("valid - when SelectUserByEmail failed", func(t *testing.T) {
		userData.On("SelectUserByEmail", mock.Anything).Return(users.UserCore{}, errors.New("err")).Once()

		user, err, status := userBusiness.FindUserByEmail(string(mock.AnythingOfType("string")))

		assert.NotEqual(t, uint(1), user.ID)
		assert.Error(t, err)
		assert.Equal(t, http.StatusInternalServerError, status)
	})

	t.Run("valid - when user not found", func(t *testing.T) {
		userData.On("SelectUserByEmail", mock.Anything).Return(users.UserCore{}, nil).Once()

		user, err, status := userBusiness.FindUserByEmail("riza@mail.com")

		assert.NotEqual(t, uint(1), user.ID)
		assert.Error(t, err)
		assert.Equal(t, http.StatusNotFound, status)
	})
}
