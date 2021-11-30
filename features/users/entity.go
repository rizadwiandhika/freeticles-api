package users

import "time"

type UserCore struct {
	ID        uint
	Username  string
	Email     string
	Followers []FollowerCore
	Role      string
	Name      string
	Password  string
	UpdatedAt time.Time
	CreatedAt time.Time
}

type FollowerCore struct {
	UserID           uint
	FollowerID       uint
	FollowerUsername string
	FollowerEmail    string
	FollowerName     string
}

type IBusiness interface {
	FindUsers() ([]UserCore, error, int)
	FindUserFollowers(username string) (UserCore, error, int)
	FindUserFollowings(username string) (UserCore, error, int)
	FindUsersByIds(ids []uint) ([]UserCore, error, int)
	FindUserById(id uint) (UserCore, error, int)
	FindUserByUsername(username string) (UserCore, error, int)
	FindUserByEmail(email string) (UserCore, error, int)
	CreateUser(user UserCore) (UserCore, error, int)
	CreateFollower(username string, followerUsername string) (error, int)
	EditUser(user UserCore) (UserCore, error, int)
	RemoveUser(username string) (error, int)
	RemoveFollowing(username string, followingUsername string) (error, int)
}

type IData interface {
	SelectUsers() ([]UserCore, error)
	SelectUserFollowers(userID uint) ([]FollowerCore, error)
	SelectUserFollowings(userID uint) ([]FollowerCore, error)
	SelectUserById(id uint) (UserCore, error)
	SelectUsersByIds(ids []uint) ([]UserCore, error)
	SelectUserByUsername(username string) (UserCore, error)
	SelectUserByEmail(email string) (UserCore, error)
	InsertUser(user UserCore) (UserCore, error)
	InsertFollower(follower FollowerCore) error
	UpdateUser(user UserCore) (UserCore, error)
	DeleteUser(username string) error
	DeleteFollowing(following FollowerCore) error
	DeleteAllUserFollow(userID uint) error
}

func (u *UserCore) IsNotFound() bool {
	return u.ID == 0
}
