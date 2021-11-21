package response

import "github.com/rizadwiandhika/miniproject-backend-alterra/features/users"

type User struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Name     string `json:"name"`
}

func FromUserCore(u *users.UserCore) User {
	return User{
		ID:       u.ID,
		Username: u.Username,
		Email:    u.Email,
		Name:     u.Name,
	}
}

func FromSliceUserCore(u []users.UserCore) []User {
	users := make([]User, len(u))
	for i, v := range u {
		users[i] = FromUserCore(&v)
	}
	return users
}
