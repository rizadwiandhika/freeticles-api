package articles

import "time"

type ArticleCore struct {
	ID        uint
	AuthorID  uint
	Author    UserCore
	Tags      []TagCore
	Title     string
	Subtitle  string
	Content   string
	Thumbnail string
	Nsfw      bool
	Likes     int
	Comments  []CommentCore
	UpdatedAt time.Time
	CreatedAt time.Time
}

type TagCore struct {
	Tag string
}

type UserCore struct {
	Username string
	Email    string
	Name     string
}

type CommentCore struct {
	ID        uint
	Comment   string
	CreatedAt time.Time
	User      UserCore
}

type QueryParams struct {
	Keyword string
	Today   bool
	Limit   int
	Offset  int
}

/* Interface for: presentation <-> bussiness layer */
type IBusiness interface {
	FindArticles(params QueryParams) ([]ArticleCore, error, int)
	FindArticleById(id uint) (ArticleCore, error, int)
	FindUserArticles(username string) ([]ArticleCore, error, int)
	RemoveArticleById(id uint) (error, int)
	CreateArticle(article ArticleCore) (ArticleCore, error, int)
	EditArticle(article ArticleCore) (ArticleCore, error, int)
}

/* Interface for: bussiness <-> data layer  */
type IData interface {
	SelectArticles(params QueryParams) ([]ArticleCore, error)
	SelectArticleById(id uint) (ArticleCore, error)
	SelectArticlesByAuthorId(id uint) ([]ArticleCore, error)
	DeleteArticleById(id uint) error
	InsertArticle(article ArticleCore) (ArticleCore, error)
	UpdateArticle(article ArticleCore) (ArticleCore, error)
}

func (a *ArticleCore) IsNotFound() bool {
	return a.ID == 0
}
