package data

import (
	"time"

	"github.com/rizadwiandhika/miniproject-backend-alterra/features/users"
)

type User struct {
	ID        uint
	Username  string
	Email     string
	Name      string
	Password  string
	UpdatedAt time.Time
	CreatedAt time.Time
}

func toUserCore(u *User) users.UserCore {
	return users.UserCore{
		ID:        u.ID,
		Username:  u.Username,
		Email:     u.Email,
		Name:      u.Name,
		UpdatedAt: u.UpdatedAt,
		CreatedAt: u.CreatedAt,
	}
}

func toSliceUserCore(u []User) []users.UserCore {
	users := make([]users.UserCore, len(u))

	for i, v := range u {
		users[i] = toUserCore(&v)
	}

	return users
}
