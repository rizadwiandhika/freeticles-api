package response

import "github.com/rizadwiandhika/miniproject-backend-alterra/features/auth"

type UserRegisterSuccess struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Email    string `json:"email"`
}

type UserRegisterFailed struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLoginFailed struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func ToUserRegisterSuccess(u auth.UserCore) UserRegisterSuccess {
	return UserRegisterSuccess{
		ID:       u.ID,
		Username: u.Username,
		Name:     u.Name,
		Email:    u.Email,
	}
}

func ToUserRegisterFailed(u auth.UserCore) UserRegisterFailed {
	return UserRegisterFailed{
		Username: u.Username,
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
	}
}

func ToUserLoginFailed(u auth.UserCore) UserLoginFailed {
	return UserLoginFailed{
		Username: u.Username,
		Email:    u.Email,
		Password: u.Password,
	}
}
