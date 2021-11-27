package reactions

import "time"

type CommentCore struct {
	ID        uint
	UserID    uint
	ArticleID uint
	Commentar string
	User      UserCore
	Article   ArticleCore
	CreatedAt time.Time
}

type LikeCore struct {
	ID        uint
	UserID    uint
	ArticleID uint
	CreatedAt time.Time
}

type ReportCore struct {
	ID           uint
	UserID       uint
	ArticleID    uint
	ReportTypeID uint
	ReportType   ReportTypeCore
	User         UserCore
	Article      ArticleCore
	CreatedAt    time.Time
}

type ReportTypeCore struct {
	ID          uint
	Name        string
	Description string
}

type UserCore struct {
	ID       uint
	Username string
	Name     string
}

type ArticleCore struct {
	ID    uint
	Title string
}

type IBusiness interface {
	FindCommentsByArticleId(articleId uint) ([]CommentCore, error, int)
	PostLike(username string, articleId uint) (error, int)
	PostComment(username string, articleId uint, commentar string) (error, int)
	Dislike(username string, articleId uint) (error, int)
	ReportArticle(username string, articleId uint, reportTypeId uint) (error, int)
}

type IData interface {
	SelectLike(like LikeCore) (LikeCore, error)
	SelectCommentsByArticleId(articleId uint) ([]CommentCore, error)
	SelectUserReport(userId uint, articleId uint) (ReportCore, error)
	InsertLike(like LikeCore) error
	InsertComment(comment CommentCore) error
	InsertReport(report ReportCore) error
	DeleteLikeById(likeId uint) error
}
