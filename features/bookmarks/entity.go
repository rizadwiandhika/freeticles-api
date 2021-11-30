package bookmarks

import "time"

type BookmarkCore struct {
	ID        uint
	UserID    uint
	ArticleID uint
	Article   ArticleCore
	CreatedAt time.Time
}

type ArticleCore struct {
	Title     string
	Subtitle  string
	Thumbnail string
}

type IBusiness interface {
	FindUserBookmarks(username string) ([]BookmarkCore, error)
	CreateBookmark(username string, articleID uint) error
	DeleteBookmark(username string, articleID uint) error
	DeleteBookmarksByArticleId(articleID uint) error
	DeleteBookmarksByUserId(userID uint) error
}

type IData interface {
	SelectDetailUserBookmark(userID uint, articleID uint) (BookmarkCore, error)
	SelectUserBookmarks(userID uint) ([]BookmarkCore, error)
	InsertBookmark(bookmark BookmarkCore) error
	DeleteBookmark(bookmark BookmarkCore) error
	DeleteBookmarksByArticleId(articleID uint) error
	DeleteBookmarksByUserId(userID uint) error
}
