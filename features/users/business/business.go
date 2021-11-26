package business

import (
	"errors"
	"net/http"

	"github.com/rizadwiandhika/miniproject-backend-alterra/features/users"
	"golang.org/x/crypto/bcrypt"
)

type userBusiness struct {
	userData users.IData
}

func NewBusiness(data users.IData) *userBusiness {
	return &userBusiness{
		userData: data,
	}
}

func (ub *userBusiness) FindUserById(id uint) (users.UserCore, error, int) {
	fetchedUser, err := ub.userData.SelectUserById(id)
	if err != nil {
		return users.UserCore{}, err, http.StatusInternalServerError
	}
	if fetchedUser.IsNotFound() {
		return users.UserCore{}, errors.New("User not found"), http.StatusNotFound
	}
	return fetchedUser, nil, http.StatusOK
}

func (ub *userBusiness) FindUsersByIds(ids []uint) ([]users.UserCore, error, int) {
	fetchedUsers, err := ub.userData.SelectUsersByIds(ids)
	if err != nil {
		return nil, err, http.StatusInternalServerError
	}
	return fetchedUsers, nil, http.StatusOK
}

func (ub *userBusiness) FindUsers() ([]users.UserCore, error, int) {
	fetchedUsers, err := ub.userData.SelectUsers()
	if err != nil {
		return nil, err, http.StatusInternalServerError
	}
	return fetchedUsers, nil, http.StatusOK
}

func (ub *userBusiness) FindUserByUsername(username string) (users.UserCore, error, int) {
	fetchedUser, err := ub.userData.SelectUserByUsername(username)
	if err != nil {
		return users.UserCore{}, err, http.StatusInternalServerError
	}
	if fetchedUser.IsNotFound() {
		return users.UserCore{}, errors.New("User not found"), http.StatusNotFound
	}
	return fetchedUser, nil, http.StatusOK
}

func (ub *userBusiness) FindUserByEmail(email string) (users.UserCore, error, int) {
	fetchedUser, err := ub.userData.SelectUserByEmail(email)
	if err != nil {
		return users.UserCore{}, err, http.StatusInternalServerError
	}
	if fetchedUser.IsNotFound() {
		return users.UserCore{}, errors.New("User not found"), http.StatusNotFound
	}
	return fetchedUser, nil, http.StatusOK
}

func (ub *userBusiness) CreateUser(user users.UserCore) (users.UserCore, error, int) {
	existingUser, err := ub.userData.SelectUserByUsername(user.Username)
	if err != nil {
		return user, err, http.StatusInternalServerError
	}
	if !existingUser.IsNotFound() {
		return user, errors.New("User already exists"), http.StatusConflict
	}

	const COST = 14
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), COST)
	if err != nil {
		return user, err, http.StatusInternalServerError
	}

	user.Password = string(hashedPassword)
	newUser, err := ub.userData.InsertUser(user)
	if err != nil {
		return user, err, http.StatusInternalServerError
	}

	return newUser, nil, http.StatusCreated
}

func (ub *userBusiness) EditUser(user users.UserCore) (users.UserCore, error, int) {
	existingUser, err := ub.userData.SelectUserByUsername(user.Username)
	if err != nil {
		return user, err, http.StatusInternalServerError
	}
	if existingUser.IsNotFound() {
		return users.UserCore{}, errors.New("User not found"), http.StatusNotFound
	}

	const COST = 14
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), COST)
	if err != nil {
		return user, err, http.StatusInternalServerError
	}

	existingUser.Username = user.Username
	existingUser.Email = user.Email
	existingUser.Name = user.Name
	existingUser.Password = string(hashedPassword)

	updatedUser, err := ub.userData.UpdateUser(existingUser)
	if err != nil {
		return user, err, http.StatusInternalServerError
	}

	return updatedUser, nil, http.StatusOK
}

func (ub *userBusiness) RemoveUser(username string) (error, int) {
	existingUser, err := ub.userData.SelectUserByUsername(username)
	if err != nil {
		return err, http.StatusInternalServerError
	}
	if existingUser.IsNotFound() {
		return errors.New("User not found"), http.StatusNotFound
	}

	err = ub.userData.DeleteUser(username)
	if err != nil {
		return err, http.StatusInternalServerError
	}
	return nil, http.StatusNoContent
}
