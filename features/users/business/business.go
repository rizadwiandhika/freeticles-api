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
