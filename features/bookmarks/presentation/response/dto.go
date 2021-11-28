package response

import "github.com/rizadwiandhika/miniproject-backend-alterra/features/bookmarks"

type Bookmark struct {
	ID        uint    `json:"id"`
	UserID    uint    `json:"userId"`
	ArticleID uint    `json:"articleId"`
	Article   Article `json:"article"`
}

type Article struct {
	Title     string `json:"title"`
	Subtitle  string `json:"subtitle"`
	Thumbnail string `json:"thumbnail"`
}

func FromBookmarkCore(b *bookmarks.BookmarkCore) Bookmark {
	return Bookmark{
		ID:        b.ID,
		UserID:    b.UserID,
		ArticleID: b.ArticleID,
		Article: Article{
			Title:     b.Article.Title,
			Subtitle:  b.Article.Subtitle,
			Thumbnail: b.Article.Thumbnail,
		},
	}
}

func FromSliceBookmarkCore(b []bookmarks.BookmarkCore) []Bookmark {
	bookmarksResponse := make([]Bookmark, len(b))
	for i := range b {
		bookmarksResponse[i] = FromBookmarkCore(&b[i])
	}
	return bookmarksResponse
}
