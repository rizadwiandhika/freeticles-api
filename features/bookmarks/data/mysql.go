package data

import (
	"github.com/rizadwiandhika/miniproject-backend-alterra/features/bookmarks"
	"gorm.io/gorm"
)

type bookmarkRepository struct {
	db *gorm.DB
}

func NewMySQLRepository(db *gorm.DB) *bookmarkRepository {
	return &bookmarkRepository{db: db}
}

func (br *bookmarkRepository) InsertBookmark(bookmark bookmarks.BookmarkCore) error {
	newBookmark := Bookmark{
		UserID:    bookmark.UserID,
		ArticleID: bookmark.ArticleID,
	}
	return br.db.Create(&newBookmark).Error
}

func (br *bookmarkRepository) SelectUserBookmarks(userID uint) ([]bookmarks.BookmarkCore, error) {
	var bookmarks []Bookmark
	err := br.db.Where("user_id = ?", userID).Find(&bookmarks).Error
	if err != nil {
		return nil, err
	}

	return ToSliceBookmarkCore(bookmarks), nil
}

func (br *bookmarkRepository) SelectDetailUserBookmark(userID uint, articleID uint) (bookmarks.BookmarkCore, error) {
	var bookmark Bookmark
	err := br.db.Where("user_id = ? AND article_id = ?", userID, articleID).First(&bookmark).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return bookmarks.BookmarkCore{}, err
	}

	return bookmark.ToBookmarkCore(), nil
}

func (br *bookmarkRepository) UpdateBookmark(bookmark bookmarks.BookmarkCore) error {
	return br.db.Model(&Bookmark{}).Where("user_id = ? AND article_id = ?", bookmark.UserID, bookmark.ArticleID).Update("user_id", bookmark.UserID).Error
}

func (br *bookmarkRepository) DeleteBookmark(bookmark bookmarks.BookmarkCore) error {
	userID := bookmark.UserID
	articleID := bookmark.ArticleID
	return br.db.Where("user_id = ? AND article_id = ?", userID, articleID).Delete(&Bookmark{}).Error
}

func (br *bookmarkRepository) DeleteBookmarksByArticleId(articleID uint) error {
	return br.db.Where("article_id = ?", articleID).Delete(&Bookmark{}).Error
}
func (br *bookmarkRepository) DeleteBookmarksByUserId(userID uint) error {
	return br.db.Where("user_id = ?", userID).Delete(&Bookmark{}).Error
}
