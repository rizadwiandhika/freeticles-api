package business

import (
	"errors"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/rizadwiandhika/miniproject-backend-alterra/config"
	"github.com/rizadwiandhika/miniproject-backend-alterra/features/auth"
	"github.com/rizadwiandhika/miniproject-backend-alterra/features/users"
	"golang.org/x/crypto/bcrypt"
)

type authBusniess struct {
	ub users.IBusiness
}

func NewBusniness(ub users.IBusiness) *authBusniess {
	return &authBusniess{
		ub: ub,
	}
}

func (ab *authBusniess) Authenticate(user auth.UserCore) (string, error, int) {
	if user.Username == "" && user.Password == "" {
		return "", errors.New("username or password is empty"), http.StatusUnprocessableEntity
	}

	var fetchedUser users.UserCore
	var err error

	if user.Username != "" {
		fetchedUser, err = ab.ub.FindUserByUsername(user.Username)
	} else {
		fetchedUser, err = ab.ub.FindUserByEmail(user.Email)
	}
	if err != nil {
		return "", err, http.StatusInternalServerError
	}

	err = bcrypt.CompareHashAndPassword([]byte(fetchedUser.Password), []byte(user.Password))
	if err != nil {
		// err = errors.New("wrong username or password")
		return "", err, http.StatusNotAcceptable
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":   fetchedUser.ID,
		"username": fetchedUser.Username,
		"exp":      time.Now().Add(3 * time.Hour).Unix(),
	})

	signedToken, err := token.SignedString([]byte(config.ENV.JWT_SECRET))
	if err != nil {
		return "", err, http.StatusInternalServerError
	}

	return signedToken, nil, http.StatusAccepted
}

func (ab *authBusniess) Register(user auth.UserCore) (auth.UserCore, error, int) {
	newUser := users.UserCore{
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
		Name:     user.Name,
	}

	result, err := ab.ub.CreateUser(newUser)
	if err != nil {
		return auth.UserCore{}, err, http.StatusInternalServerError
	}

	createdUser := auth.UserCore{
		ID:       result.ID,
		Username: result.Username,
		Email:    result.Email,
		Name:     result.Name,
		Password: result.Password,
	}

	return createdUser, err, http.StatusCreated
}
