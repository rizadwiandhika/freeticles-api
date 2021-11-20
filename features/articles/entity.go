package articles

import "time"

type ArticleCore struct {
	ID        uint
	AuthorID  uint
	Author    AuthorCore
	Tags      []TagCore
	Title     string
	Subtitle  string
	Content   string
	Thumbnail string
	Nsfw      bool
	UpdatedAt time.Time
	CreatedAt time.Time
}

type TagCore struct {
	Tag string
}

type AuthorCore struct {
	Username string
	Email    string
	Name     string
}

/* Interface for: presentation <-> bussiness layer */
type IBusiness interface {
	FindArticleById(id uint) (ArticleCore, error)
}

/* Interface for: bussiness <-> data layer  */
type IData interface {
	SelectArticleById(id uint) (ArticleCore, error)
}
