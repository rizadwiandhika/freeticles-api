package data

import (
	"github.com/rizadwiandhika/miniproject-backend-alterra/features/users"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewMySQLRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

func (ur *userRepository) SelectUserById(id uint) (users.UserCore, error) {
	user := User{}

	err := ur.db.First(&user, id).Error
	if err != nil {
		return users.UserCore{}, err
	}

	return toUserCore(&user), nil
}
