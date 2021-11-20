package users

import "time"

type UserCore struct {
	ID        uint
	Username  string
	Email     string
	Name      string
	Password  string
	UpdatedAt time.Time
	CreatedAt time.Time
}

type IBusiness interface {
	FindUserById(id uint) (UserCore, error)
}

type IData interface {
	SelectUserById(id uint) (UserCore, error)
}
