package request

import "github.com/rizadwiandhika/miniproject-backend-alterra/features/users"

type User struct {
	Username string `param:"username"`
	Name     string `json:"name"`
	Email    string `json:"email"`
}

func ToUserCore(user User) users.UserCore {
	return users.UserCore{
		Username: user.Username,
		Name:     user.Name,
		Email:    user.Email,
	}
}
