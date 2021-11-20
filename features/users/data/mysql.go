package data

import (
	"github.com/rizadwiandhika/miniproject-backend-alterra/features/users"
	"golang.org/x/crypto/bcrypt"
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
	const COST = 14

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), COST)
	if err != nil {
		return users.UserCore{}, err
	}

	newUser := User{
		Username: user.Username,
		Email:    user.Email,
		Name:     user.Name,
		Password: string(hashedPassword),
	}

	err = ur.db.Create(&newUser).Error
	if err != nil {
		return users.UserCore{}, err
	}

	userCore := toUserCore(&newUser)
	userCore.Password = user.Password // Give back user the raw password. Not the hashed one.

	return userCore, nil
}

func (ur *userRepository) UpdateUser(user users.UserCore) (users.UserCore, error) {
	const COST = 14

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), COST)
	if err != nil {
		return users.UserCore{}, err
	}

	updatedUser := User{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Name:     user.Name,
		Password: string(hashedPassword),
	}

	err = ur.db.Save(&updatedUser).Error
	if err != nil {
		return users.UserCore{}, err
	}

	userCore := toUserCore(&updatedUser)
	userCore.Password = user.Password // Give back user the raw password. Not the hashed one.

	return userCore, nil
}

func (ur *userRepository) DeleteUser(username string) error {
	err := ur.db.Where("username = ?", username).Delete(User{}).Error
	return err
}
