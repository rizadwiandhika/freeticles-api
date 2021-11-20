package business

import "github.com/rizadwiandhika/miniproject-backend-alterra/features/users"

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
	return ub.userData.InsertUser(user)
}

func (ub *userBusiness) EditUser(user users.UserCore) (users.UserCore, error) {
	return ub.userData.UpdateUser(user)
}

func (ub *userBusiness) RemoveUser(username string) error {
	return ub.userData.DeleteUser(username)
}
