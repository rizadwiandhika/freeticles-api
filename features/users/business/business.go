package business

import (
	"errors"

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

func (ub *userBusiness) FindUserById(id uint) (users.UserCore, error) {
	return ub.userData.SelectUserById(id)
}

func (ub *userBusiness) FindUsersByIds(ids []uint) ([]users.UserCore, error) {
	return ub.userData.SelectUsersByIds(ids)
}

func (ub *userBusiness) FindUsers() ([]users.UserCore, error) {
	return ub.userData.SelectUsers()
}

func (ub *userBusiness) FindUserByUsername(username string) (users.UserCore, error) {
	return ub.userData.SelectUserByUsername(username)
}

func (ub *userBusiness) FindUserByEmail(email string) (users.UserCore, error) {
	return ub.userData.SelectUserByEmail(email)
}

func (ub *userBusiness) CreateUser(user users.UserCore) (users.UserCore, error) {
	const COST = 14
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), COST)
	if err != nil {
		return users.UserCore{}, err
	}

	user.Password = string(hashedPassword)
	newUser, err := ub.userData.InsertUser(user)
	if err != nil {
		return users.UserCore{}, err
	}

	return newUser, nil
}

func (ub *userBusiness) EditUser(user users.UserCore) (users.UserCore, error) {
	existingUser, err := ub.userData.SelectUserByUsername(user.Username)
	if err != nil {
		return users.UserCore{}, err
	}

	if existingUser.IsNotFound() {
		return users.UserCore{}, errors.New("User not found")
	}

	const COST = 14
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), COST)
	if err != nil {
		return users.UserCore{}, err
	}

	existingUser.Username = user.Username
	existingUser.Email = user.Email
	existingUser.Name = user.Name
	existingUser.Password = string(hashedPassword)

	return ub.userData.UpdateUser(existingUser)
}

func (ub *userBusiness) RemoveUser(username string) error {
	existingUser, err := ub.userData.SelectUserByUsername(username)
	if err != nil {
		return err
	}

	if existingUser.IsNotFound() {
		return errors.New("User not found")
	}

	return ub.userData.DeleteUser(username)
}
