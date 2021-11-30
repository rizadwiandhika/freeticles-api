package business

import (
	"errors"
	"fmt"
	"net/http"
	"sync"

	"github.com/rizadwiandhika/miniproject-backend-alterra/features/articles"
	"github.com/rizadwiandhika/miniproject-backend-alterra/features/bookmarks"
	"github.com/rizadwiandhika/miniproject-backend-alterra/features/reactions"
	"github.com/rizadwiandhika/miniproject-backend-alterra/features/users"
	"golang.org/x/crypto/bcrypt"
)

type userBusiness struct {
	userData         users.IData
	articleBusiness  articles.IBusiness
	reactionBusiness reactions.IBusiness
	bookmarkBusiness bookmarks.IBusiness
}

func NewBusiness(data users.IData, ab articles.IBusiness, rb reactions.IBusiness, bb bookmarks.IBusiness) *userBusiness {
	return &userBusiness{
		userData:         data,
		articleBusiness:  ab,
		reactionBusiness: rb,
		bookmarkBusiness: bb,
	}
}

func (ub *userBusiness) FindUserById(id uint) (users.UserCore, error, int) {
	fetchedUser, err := ub.userData.SelectUserById(id)
	if err != nil {
		return users.UserCore{}, err, http.StatusInternalServerError
	}
	if fetchedUser.IsNotFound() {
		return users.UserCore{}, errors.New("User not found"), http.StatusNotFound
	}
	return fetchedUser, nil, http.StatusOK
}

func (ub *userBusiness) FindUsersByIds(ids []uint) ([]users.UserCore, error, int) {
	fetchedUsers, err := ub.userData.SelectUsersByIds(ids)
	if err != nil {
		return nil, err, http.StatusInternalServerError
	}
	return fetchedUsers, nil, http.StatusOK
}

func (ub *userBusiness) FindUsers() ([]users.UserCore, error, int) {
	fetchedUsers, err := ub.userData.SelectUsers()
	if err != nil {
		return nil, err, http.StatusInternalServerError
	}
	return fetchedUsers, nil, http.StatusOK
}

func (ub *userBusiness) FindUserFollowers(username string) (users.UserCore, error, int) {
	user, err := ub.userData.SelectUserByUsername(username)
	if err != nil {
		return user, err, http.StatusInternalServerError
	}
	if user.IsNotFound() {
		return user, errors.New("User not found"), http.StatusNotFound
	}

	followers, err := ub.userData.SelectUserFollowers(user.ID)
	if err != nil {
		return user, err, http.StatusInternalServerError
	}

	followersUserIds := make([]uint, len(followers))
	for i, following := range followers {
		followersUserIds[i] = following.FollowerID
	}

	followerUsers, err := ub.userData.SelectUsersByIds(followersUserIds)
	if err != nil {
		return user, err, http.StatusInternalServerError
	}

	followersMap := make(map[uint]users.UserCore)
	for _, followingUser := range followerUsers {
		followersMap[followingUser.ID] = followingUser
	}

	for i := range followers {
		FollowingUserID := followers[i].FollowerID
		followers[i].FollowerUsername = followersMap[FollowingUserID].Username
		followers[i].FollowerEmail = followersMap[FollowingUserID].Email
		followers[i].FollowerName = followersMap[FollowingUserID].Name
	}

	user.Followers = followers
	return user, nil, http.StatusOK
}

func (ub *userBusiness) FindUserFollowings(username string) (users.UserCore, error, int) {
	user, err := ub.userData.SelectUserByUsername(username)
	if err != nil {
		return user, err, http.StatusInternalServerError
	}
	if user.IsNotFound() {
		return user, errors.New("User not found"), http.StatusNotFound
	}

	followings, err := ub.userData.SelectUserFollowings(user.ID)
	if err != nil {
		return user, err, http.StatusInternalServerError
	}
	fmt.Printf("%+v\n", followings)

	followingIds := make([]uint, len(followings))
	for i, following := range followings {
		followingIds[i] = following.UserID
	}

	followingUsers, err := ub.userData.SelectUsersByIds(followingIds)
	if err != nil {
		return user, err, http.StatusInternalServerError
	}

	followingsMap := make(map[uint]users.UserCore)
	for _, followingUser := range followingUsers {
		followingsMap[followingUser.ID] = followingUser
	}

	for i := range followings {
		FollowingUserID := followings[i].UserID
		followings[i].FollowerUsername = followingsMap[FollowingUserID].Username
		followings[i].FollowerEmail = followingsMap[FollowingUserID].Email
		followings[i].FollowerName = followingsMap[FollowingUserID].Name
	}

	user.Followers = followings
	return user, nil, http.StatusOK
}

func (ub *userBusiness) FindUserByUsername(username string) (users.UserCore, error, int) {
	fetchedUser, err := ub.userData.SelectUserByUsername(username)
	if err != nil {
		return users.UserCore{}, err, http.StatusInternalServerError
	}
	if fetchedUser.IsNotFound() {
		return users.UserCore{}, errors.New("User not found"), http.StatusNotFound
	}
	return fetchedUser, nil, http.StatusOK
}

func (ub *userBusiness) FindUserByEmail(email string) (users.UserCore, error, int) {
	fetchedUser, err := ub.userData.SelectUserByEmail(email)
	if err != nil {
		return users.UserCore{}, err, http.StatusInternalServerError
	}
	if fetchedUser.IsNotFound() {
		return users.UserCore{}, errors.New("User not found"), http.StatusNotFound
	}
	return fetchedUser, nil, http.StatusOK
}

func (ub *userBusiness) CreateUser(user users.UserCore) (users.UserCore, error, int) {
	existingUser, err := ub.userData.SelectUserByUsername(user.Username)
	if err != nil {
		return user, err, http.StatusInternalServerError
	}
	if !existingUser.IsNotFound() {
		return user, errors.New("User already exists"), http.StatusConflict
	}

	const COST = 14
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), COST)
	if err != nil {
		return user, err, http.StatusInternalServerError
	}

	user.Password = string(hashedPassword)
	newUser, err := ub.userData.InsertUser(user)
	if err != nil {
		return user, err, http.StatusInternalServerError
	}

	return newUser, nil, http.StatusCreated
}

func (ub *userBusiness) CreateFollower(username string, followerUsername string) (error, int) {
	user, err := ub.userData.SelectUserByUsername(username)
	if err != nil {
		return err, http.StatusInternalServerError
	}
	if user.IsNotFound() {
		return errors.New("User not found"), http.StatusNotFound
	}

	followerUser, err := ub.userData.SelectUserByUsername(followerUsername)
	if err != nil {
		return err, http.StatusInternalServerError
	}
	if followerUser.IsNotFound() {
		return errors.New("User not found"), http.StatusNotFound
	}

	follower := users.FollowerCore{
		UserID:     user.ID,
		FollowerID: followerUser.ID,
	}
	err = ub.userData.InsertFollower(follower)
	if err != nil {
		return err, http.StatusInternalServerError
	}
	return nil, http.StatusCreated
}

func (ub *userBusiness) EditUser(user users.UserCore) (users.UserCore, error, int) {
	existingUser, err := ub.userData.SelectUserByUsername(user.Username)
	if err != nil {
		return user, err, http.StatusInternalServerError
	}
	if existingUser.IsNotFound() {
		return users.UserCore{}, errors.New("User not found"), http.StatusNotFound
	}

	const COST = 14
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), COST)
	if err != nil {
		return user, err, http.StatusInternalServerError
	}

	existingUser.Username = user.Username
	existingUser.Email = user.Email
	existingUser.Name = user.Name
	existingUser.Password = string(hashedPassword)

	updatedUser, err := ub.userData.UpdateUser(existingUser)
	if err != nil {
		return user, err, http.StatusInternalServerError
	}

	return updatedUser, nil, http.StatusOK
}

func (ub *userBusiness) RemoveUser(username string) (error, int) {
	existingUser, err := ub.userData.SelectUserByUsername(username)
	if err != nil {
		return err, http.StatusInternalServerError
	}
	if existingUser.IsNotFound() {
		return errors.New("User not found"), http.StatusNotFound
	}

	var wg sync.WaitGroup
	wg.Add(4)

	go func() {
		_ = ub.bookmarkBusiness.DeleteBookmarksByUserId(existingUser.ID)
		wg.Done()
	}()
	go func() {
		_ = ub.reactionBusiness.RemoveCommentsByUserId(existingUser.ID)
		wg.Done()
	}()
	go func() {
		_ = ub.reactionBusiness.RemoveLikesByUserId(existingUser.ID)
		wg.Done()
	}()
	go func() {
		_ = ub.reactionBusiness.RemoveReportsByUserId(existingUser.ID)
		wg.Done()
	}()

	wg.Wait()
	err, _ = ub.articleBusiness.RemoveUserArticles(existingUser.ID)
	if err != nil {
		return err, http.StatusInternalServerError
	}

	err = ub.userData.DeleteAllUserFollow(existingUser.ID)
	if err != nil {
		return err, http.StatusInternalServerError
	}

	err = ub.userData.DeleteUser(username)
	if err != nil {
		return err, http.StatusInternalServerError
	}

	return nil, http.StatusNoContent
}

func (ub *userBusiness) RemoveFollowing(username string, followingUsername string) (error, int) {
	user, err := ub.userData.SelectUserByUsername(username)
	if err != nil {
		return err, http.StatusInternalServerError
	}
	if user.IsNotFound() {
		return errors.New("User following not found"), http.StatusNotFound
	}

	followingUser, err := ub.userData.SelectUserByUsername(followingUsername)
	if err != nil {
		return err, http.StatusInternalServerError
	}
	if followingUser.IsNotFound() {
		return errors.New("User following not found"), http.StatusNotFound
	}

	following := users.FollowerCore{
		UserID:     followingUser.ID,
		FollowerID: user.ID,
	}
	err = ub.userData.DeleteFollowing(following)
	if err != nil {
		return err, http.StatusInternalServerError
	}
	return nil, http.StatusOK
}
