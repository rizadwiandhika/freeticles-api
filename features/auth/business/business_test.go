package business_test

import (
	"errors"
	"net/http"
	"os"
	"testing"

	"github.com/rizadwiandhika/miniproject-backend-alterra/features/auth"
	authb "github.com/rizadwiandhika/miniproject-backend-alterra/features/auth/business"
	"github.com/rizadwiandhika/miniproject-backend-alterra/features/users"
	umock "github.com/rizadwiandhika/miniproject-backend-alterra/features/users/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

var (
	userBusiness umock.IBusiness
	authBusniess auth.IBusiness

	userValue users.UserCore
)

func TestMain(m *testing.M) {
	authBusniess = authb.NewBusniness(&userBusiness)

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("riza123"), 12)
	userValue = users.UserCore{
		Username: "riza.dwii",
		Email:    "riza@mail.com",
		Password: string(hashedPassword),
	}

	os.Exit(m.Run())
}

func TestAuthenticate(t *testing.T) {
	t.Run("valid - when using username", func(t *testing.T) {
		var mc *mock.Call

		mc = userBusiness.On("FindUserByUsername", mock.Anything)
		mc = mc.Return(userValue, nil, 200)
		mc.Once()

		loginUser := auth.UserCore{
			Username: "riza.dwii",
			Password: "riza123",
		}
		token, _, _ := authBusniess.Authenticate(loginUser)

		assert.NotEmpty(t, token)
	})

	t.Run("valid - when using email", func(t *testing.T) {
		var mc *mock.Call

		mc = userBusiness.On("FindUserByEmail", mock.Anything)
		mc = mc.Return(userValue, nil, 200)
		mc.Once()

		loginUser := auth.UserCore{
			Email:    "riza@mail.com",
			Password: "riza123",
		}
		token, _, _ := authBusniess.Authenticate(loginUser)

		assert.NotEmpty(t, token)
	})

	t.Run("valid - error when username & password is empty", func(t *testing.T) {
		token, err, _ := authBusniess.Authenticate(auth.UserCore{})

		assert.Empty(t, token)
		assert.NotNil(t, err)
	})

	t.Run("valid - error when user not found", func(t *testing.T) {
		var mc *mock.Call

		mc = userBusiness.On("FindUserByUsername", mock.Anything)
		mc = mc.Return(users.UserCore{}, errors.New("abc"), 404)
		mc.Once()

		loginUser := auth.UserCore{
			Username: "riza.dwii",
			Password: "riza123",
		}
		token, err, status := authBusniess.Authenticate(loginUser)

		assert.Empty(t, token)
		assert.NotNil(t, err)
		assert.Equal(t, http.StatusInternalServerError, status)
	})

	t.Run("valid - error invalid password", func(t *testing.T) {
		var mc *mock.Call

		mc = userBusiness.On("FindUserByUsername", mock.Anything)
		mc = mc.Return(userValue, nil, 200)
		mc.Once()

		loginUser := auth.UserCore{
			Username: "riza.dwii",
			Password: "riza",
		}
		token, err, status := authBusniess.Authenticate(loginUser)

		assert.Empty(t, token)
		assert.NotNil(t, err)
		assert.Equal(t, http.StatusUnauthorized, status)
	})
}

func TestRegister(t *testing.T) {
	t.Run("valid - Register", func(t *testing.T) {
		var mc *mock.Call

		mc = userBusiness.On("CreateUser", mock.AnythingOfType("users.UserCore"))
		mc = mc.Return(userValue, nil, 200)
		mc.Once()

		registerUser := auth.UserCore{
			Username: "riza.dwii",
			Email:    "riza@mail.com",
			Password: "riza123",
		}

		_, err, status := authBusniess.Register(registerUser)

		assert.Nil(t, err)
		assert.Equal(t, http.StatusCreated, status)
	})

	t.Run("valid - error when CreateUser failed", func(t *testing.T) {
		var mc *mock.Call

		mc = userBusiness.On("CreateUser", mock.AnythingOfType("users.UserCore"))
		mc = mc.Return(userValue, errors.New("abc"), 123)
		mc.Once()

		registerUser := auth.UserCore{
			Username: "riza.dwii",
			Email:    "riza@mail.com",
			Password: "riza123",
		}

		_, err, status := authBusniess.Register(registerUser)

		assert.NotNil(t, err)
		assert.Equal(t, http.StatusInternalServerError, status)
	})
}
