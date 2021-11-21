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

func (ur *userRepository) SelectUsers() ([]users.UserCore, error) {
	users := []User{}

	err := ur.db.Find(&users).Error
	if err != nil {
		return nil, err
	}

	return toSliceUserCore(users), nil
}

func (ur *userRepository) SelectUserByUsername(username string) (users.UserCore, error) {
	user := User{}

	err := ur.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return users.UserCore{}, err
	}

	return toUserCore(&user), nil
}

func (ur *userRepository) SelectUserByEmail(email string) (users.UserCore, error) {
	user := User{}

	err := ur.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return users.UserCore{}, err
	}

	return toUserCore(&user), nil
}

func (ur *userRepository) InsertUser(user users.UserCore) (users.UserCore, error) {
	newUser := User{
		Username: user.Username,
		Email:    user.Email,
		Name:     user.Name,
		Password: user.Password,
	}

	err := ur.db.Create(&newUser).Error
	if err != nil {
		return users.UserCore{}, err
	}

	return toUserCore(&newUser), nil
}

func (ur *userRepository) UpdateUser(user users.UserCore) (users.UserCore, error) {
	updatedUser := User{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Name:      user.Name,
		Password:  user.Password,
		UpdatedAt: user.UpdatedAt,
		CreatedAt: user.CreatedAt,
	}

	err := ur.db.Save(&updatedUser).Error
	if err != nil {
		return users.UserCore{}, err
	}

	return toUserCore(&updatedUser), nil
}

func (ur *userRepository) DeleteUser(username string) error {
	err := ur.db.Where("username = ?", username).Delete(User{}).Error
	return err
}
