package data

import (
	"time"

	"github.com/rizadwiandhika/miniproject-backend-alterra/features/users"
)

type User struct {
	ID        uint   `gorm:"primariKey"`
	Username  string `gorm:"size:64;unique;not null"`
	Email     string `gorm:"size:64;unique;not null"`
	Name      string `gorm:"size:64;not null"`
	Password  string `gorm:"size:512;not null"`
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
