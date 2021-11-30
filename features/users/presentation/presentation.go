package presentation

import (
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/rizadwiandhika/miniproject-backend-alterra/features/users"
	"github.com/rizadwiandhika/miniproject-backend-alterra/features/users/presentation/request"
	"github.com/rizadwiandhika/miniproject-backend-alterra/features/users/presentation/response"
)

type any interface{}
type json map[string]any

type UserPresentation struct {
	userBusiness users.IBusiness
}

func NewPresentation(articleBusiness users.IBusiness) *UserPresentation {
	return &UserPresentation{articleBusiness}
}

func (up *UserPresentation) GetUsers(c echo.Context) error {
	users, err, status := up.userBusiness.FindUsers()

	if err != nil {
		return c.JSON(status, json{
			"message": "Could not get users",
			"error":   err.Error(),
		})
	}

	return c.JSON(status, json{
		"users": response.FromSliceUserCore(users),
	})
}

func (up *UserPresentation) GetDetailUser(c echo.Context) error {
	var username string
	echo.PathParamsBinder(c).String("username", &username)

	user, err, status := up.userBusiness.FindUserByUsername(username)
	if err != nil {
		return c.JSON(status, json{
			"username": username,
			"message":  "Failed retrieving user",
			"error":    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, json{
		"message": "Success retrieving user",
		"user":    response.FromUserCore(&user),
	})
}

func (up *UserPresentation) PutEditUser(c echo.Context) error {
	var username string
	echo.PathParamsBinder(c).String("username", &username)

	issuer := c.Get("user").(jwt.MapClaims)

	if issuer["username"] != username && issuer["role"] != "admin" {
		return c.JSON(http.StatusForbidden, json{
			"message": "Unauthorized user!",
		})
	}

	var user request.User
	err := c.Bind(&user)
	if err != nil {
		core := request.ToUserCore(user)
		return c.JSON(http.StatusInternalServerError, json{
			"message": "Failed updating user",
			"error":   err.Error(),
			"user":    response.FromUserCore(&core),
		})
	}

	userCore := request.ToUserCore(user)
	editedUser, err, status := up.userBusiness.EditUser(userCore)
	if err != nil {
		return c.JSON(status, json{
			"message": "Failed updating user",
			"error":   err.Error(),
			"user":    response.FromUserCore(&userCore),
		})
	}

	return c.JSON(status, json{
		"message":     "Success updating user",
		"updatedUser": response.FromUserCore(&editedUser),
	})
}

func (up *UserPresentation) DeleteUser(c echo.Context) error {
	var username string
	echo.PathParamsBinder(c).String("username", &username)

	issuer := c.Get("user").(jwt.MapClaims)

	if issuer["username"] != username && issuer["role"] != "admin" {
		return c.JSON(http.StatusForbidden, json{
			"message": "Unauthorized user!",
		})
	}

	err, status := up.userBusiness.RemoveUser(username)
	if err != nil {
		return c.JSON(status, json{
			"message": "Failed deleting user",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, json{"message": "Delete user success"})
}

func (up *UserPresentation) GetUserFollowers(c echo.Context) error {
	var username string
	echo.PathParamsBinder(c).String("username", &username)

	user, err, status := up.userBusiness.FindUserFollowers(username)
	if err != nil {
		return c.JSON(status, json{
			"message": "Failed retrieving followers",
			"error":   err.Error(),
		})
	}

	return c.JSON(status, json{
		"message":   "Success retrieving followers",
		"username":  username,
		"followers": response.FromSliceFollowerCore(user.Followers),
	})
}

func (up *UserPresentation) GetUserFollowings(c echo.Context) error {
	var username string
	echo.PathParamsBinder(c).String("username", &username)

	user, err, status := up.userBusiness.FindUserFollowings(username)
	if err != nil {
		return c.JSON(status, json{
			"message": "Failed retrieving followers",
			"error":   err.Error(),
		})
	}

	return c.JSON(status, json{
		"message":   "Success retrieving followers",
		"username":  username,
		"followers": response.FromSliceFollowerCore(user.Followers),
	})
}

func (up *UserPresentation) PostUserFollowing(c echo.Context) error {
	var username string
	var followingsUsername string

	echo.PathParamsBinder(c).String("username", &username)
	echo.PathParamsBinder(c).String("followingusername", &followingsUsername)

	issuer := c.Get("user").(jwt.MapClaims)

	if issuer["username"] != username && issuer["role"] != "admin" {
		return c.JSON(http.StatusForbidden, json{
			"message": "Unauthorized user!",
		})
	}

	err, status := up.userBusiness.CreateFollower(followingsUsername, username)
	if err != nil {
		return c.JSON(status, json{
			"message": "Failed following user",
			"error":   err.Error(),
		})
	}

	return c.JSON(status, json{
		"message": "Success following user",
	})
}

func (up *UserPresentation) DeleteUserFollowing(c echo.Context) error {
	var username string
	var followingUsername string

	echo.PathParamsBinder(c).String("username", &username)
	echo.PathParamsBinder(c).String("followingusername", &followingUsername)

	issuer := c.Get("user").(jwt.MapClaims)

	if issuer["username"] != username && issuer["role"] != "admin" {
		return c.JSON(http.StatusForbidden, json{
			"message": "Unauthorized user!",
		})
	}

	err, status := up.userBusiness.RemoveFollowing(username, followingUsername)
	if err != nil {
		return c.JSON(status, json{
			"message": "Failed unfollowing user",
			"error":   err.Error(),
		})
	}

	return c.JSON(status, json{"message": "Success unfollowing user"})
}
