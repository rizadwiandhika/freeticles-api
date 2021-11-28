package data

import (
	"time"

	"github.com/rizadwiandhika/miniproject-backend-alterra/features/users"
)

type User struct {
	ID        uint    `gorm:"primariKey"`
	Username  string  `gorm:"size:64;unique;not null"`
	Role      string  `gorm:"size:32;not null;default:user"`
	Email     string  `gorm:"size:64;unique;not null"`
	Followers []*User `gorm:"many2many:followers;"`
	Name      string  `gorm:"size:64;not null"`
	Password  string  `gorm:"size:512;not null"`
	UpdatedAt time.Time
	CreatedAt time.Time
}

type Follower struct {
	UserID     uint `gorm:"primaryKey"`
	FollowerID uint `gorm:"primaryKey"`
	CreatedAt  time.Time
}

func toUserCore(u *User) users.UserCore {
	return users.UserCore{
		ID:        u.ID,
		Username:  u.Username,
		Email:     u.Email,
		Role:      u.Role,
		Name:      u.Name,
		Password:  u.Password,
		UpdatedAt: u.UpdatedAt,
		CreatedAt: u.CreatedAt,
	}
}

func toFollowerCore(f *Follower) users.FollowerCore {
	return users.FollowerCore{
		UserID:     f.UserID,
		FollowerID: f.FollowerID,
	}
}

func toSliceUserCore(u []User) []users.UserCore {
	users := make([]users.UserCore, len(u))

	for i, v := range u {
		users[i] = toUserCore(&v)
	}

	return users
}

func toSliceFollowerCore(f []Follower) []users.FollowerCore {
	follows := make([]users.FollowerCore, len(f))

	for i, v := range f {
		follows[i] = toFollowerCore(&v)
	}

	return follows
}
