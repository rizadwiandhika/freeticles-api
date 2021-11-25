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
	FindUsers() ([]UserCore, error)
	FindUsersByIds(ids []uint) ([]UserCore, error)
	FindUserById(id uint) (UserCore, error)
	FindUserByUsername(username string) (UserCore, error)
	FindUserByEmail(email string) (UserCore, error)
	CreateUser(user UserCore) (UserCore, error)
	EditUser(user UserCore) (UserCore, error)
	RemoveUser(username string) error
}

type IData interface {
	SelectUsers() ([]UserCore, error)
	SelectUserById(id uint) (UserCore, error)
	SelectUsersByIds(ids []uint) ([]UserCore, error)
	SelectUserByUsername(username string) (UserCore, error)
	SelectUserByEmail(email string) (UserCore, error)
	InsertUser(user UserCore) (UserCore, error)
	UpdateUser(user UserCore) (UserCore, error)
	DeleteUser(username string) error
}
