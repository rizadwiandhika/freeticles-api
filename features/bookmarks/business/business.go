package business

import (
	"errors"

	"github.com/rizadwiandhika/miniproject-backend-alterra/features/articles"
	"github.com/rizadwiandhika/miniproject-backend-alterra/features/bookmarks"
	"github.com/rizadwiandhika/miniproject-backend-alterra/features/users"
)

type bookmarkBusiness struct {
	bookmarkData    bookmarks.IData
	articleBusiness articles.IBusiness
	userBusiness    users.IBusiness
}

func NewBusiness(bd bookmarks.IData, ub users.IBusiness, ab articles.IBusiness) *bookmarkBusiness {
	return &bookmarkBusiness{
		bookmarkData:    bd,
		articleBusiness: ab,
		userBusiness:    ub,
	}
}

func (bb *bookmarkBusiness) CreateBookmark(username string, articleID uint) error {
	user, err, _ := bb.userBusiness.FindUserByUsername(username)
	if err != nil {
		return err
	}

	article, err, _ := bb.articleBusiness.FindArticleById(articleID)
	if err != nil {
		return err
	}

	existingBookmark, err := bb.bookmarkData.SelectDetailUserBookmark(user.ID, article.ID)
	if err != nil {
		return err
	}
	if existingBookmark.ID != 0 {
		return errors.New("Bookmark already exists")
	}

	newBookmark := bookmarks.BookmarkCore{
		UserID:    user.ID,
		ArticleID: articleID,
	}
	return bb.bookmarkData.InsertBookmark(newBookmark)
}

func (bb *bookmarkBusiness) FindUserBookmarks(username string) ([]bookmarks.BookmarkCore, error) {
	user, err, _ := bb.userBusiness.FindUserByUsername(username)
	if err != nil {
		return []bookmarks.BookmarkCore{}, err
	}

	userBookmarks, err := bb.bookmarkData.SelectUserBookmarks(user.ID)
	if err != nil {
		return []bookmarks.BookmarkCore{}, err
	}

	for i := range userBookmarks {
		article, err, _ := bb.articleBusiness.FindArticleById(userBookmarks[i].ArticleID)
		if err != nil {
			return []bookmarks.BookmarkCore{}, err
		}
		userBookmarks[i].Article.Title = article.Title
		userBookmarks[i].Article.Subtitle = article.Subtitle
		userBookmarks[i].Article.Thumbnail = article.Thumbnail
	}

	return userBookmarks, nil
}

func (bb *bookmarkBusiness) DeleteBookmark(username string, articleID uint) error {
	user, err, _ := bb.userBusiness.FindUserByUsername(username)
	if err != nil {
		return err
	}

	bookmark := bookmarks.BookmarkCore{
		UserID:    user.ID,
		ArticleID: articleID,
	}
	return bb.bookmarkData.DeleteBookmark(bookmark)
}
