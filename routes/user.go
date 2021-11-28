package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/rizadwiandhika/miniproject-backend-alterra/factory"
	"github.com/rizadwiandhika/miniproject-backend-alterra/middleware"
)

func SetupUserRoutes(e *echo.Echo, presenter *factory.Presenter) {
	routes := e.Group("/users")

	routes.DELETE("/:username/followings/:followingusername", presenter.UserPresentation.DeleteUserFollowing, middleware.IsAuth())
	routes.DELETE("/:username/bookmarks/:articleid", presenter.BookmarkPresentation.DeleteBookmark, middleware.IsAuth())
	routes.DELETE("/:username", presenter.UserPresentation.DeleteUser, middleware.IsAuth())

	routes.POST("/:username/followings/:followingusername", presenter.UserPresentation.PostUserFollowing, middleware.IsAuth())
	routes.POST("/:username/bookmarks/:articleid", presenter.BookmarkPresentation.PostBookmark, middleware.IsAuth())

	routes.PUT("/:username", presenter.UserPresentation.PutEditUser, middleware.IsAuth())

	routes.GET("/:username/articles", presenter.ArticlePresentation.GetUserArticles)
	routes.GET("/:username/bookmarks", presenter.BookmarkPresentation.GetUserBookmarks, middleware.IsAuth())
	routes.GET("/:username/followers", presenter.UserPresentation.GetUserFollowers)
	routes.GET("/:username/followings", presenter.UserPresentation.GetUserFollowings)
	routes.GET("/:username", presenter.UserPresentation.GetDetailUser)
	routes.GET("", presenter.UserPresentation.GetUsers)
}
