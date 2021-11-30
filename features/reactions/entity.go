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
	CountTotalArticleLikes(articleId uint) (int, error)
	PostLike(username string, articleId uint) (error, int)
	PostComment(username string, articleId uint, commentar string) (error, int)
	Dislike(username string, articleId uint) (error, int)
	ReportArticle(username string, articleId uint, reportTypeId uint) (error, int)
	RemoveCommentsByArticleId(articleId uint) error
	RemoveLikesByArticleId(articleId uint) error
	RemoveReportsByArticleId(articleId uint) error
	RemoveCommentsByUserId(userId uint) error
	RemoveLikesByUserId(userId uint) error
	RemoveReportsByUserId(userId uint) error
}

type IData interface {
	SelectLike(like LikeCore) (LikeCore, error)
	SelectCountLikes(articleId uint) (int, error)
	SelectCommentsByArticleId(articleId uint) ([]CommentCore, error)
	SelectUserReport(userId uint, articleId uint) (ReportCore, error)
	InsertLike(like LikeCore) error
	InsertComment(comment CommentCore) error
	InsertReport(report ReportCore) error
	DeleteLikeById(likeId uint) error
	DeleteCommentsByArticleId(articleId uint) error
	DeleteLikesByArticleId(articleId uint) error
	DeleteReportsByArticleId(articleId uint) error
	DeleteCommentsByUserId(userId uint) error
	DeleteLikesByUserId(userId uint) error
	DeleteReportsByUserId(userId uint) error
}
