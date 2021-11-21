package request

import "github.com/rizadwiandhika/miniproject-backend-alterra/features/auth"

type Register struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (r *Register) ToUserCore() auth.UserCore {
	return auth.UserCore{
		Username: r.Username,
		Email:    r.Email,
		Name:     r.Name,
		Password: r.Password,
	}
}

func (l *Login) ToUserCore() auth.UserCore {
	return auth.UserCore{
		Username: l.Username,
		Password: l.Password,
	}
}
