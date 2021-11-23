package response

import (
	"time"

	"github.com/rizadwiandhika/miniproject-backend-alterra/features/articles"
)

type Article struct {
	ID        uint      `json:"id"`
	AuthorID  uint      `json:"authorId"`
	Author    User      `json:"author"`
	Tags      []Tag     `json:"tags"`
	Title     string    `json:"title"`
	Subtitle  string    `json:"subtitle"`
	Content   string    `json:"content"`
	Thumbnail string    `json:"thumbnail"`
	Nsfw      bool      `json:"nsfw"`
	Likes     int       `json:"likes"`
	Comments  []Comment `json:"comments"`
	UpdatedAt time.Time `json:"updatedAt"`
	CreatedAt time.Time `json:"createdAt"`
}

type Tag struct {
	Tag string `json:"tag"`
}

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Name     string `json:"name"`
}

type Comment struct {
	ID        uint      `json:"id"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"createdAt"`
	User      User      `json:"user"`
}

func FromArticleCore(a *articles.ArticleCore) Article {
	return Article{
		ID:        a.ID,
		AuthorID:  a.AuthorID,
		Author:    FromUserCore(&a.Author),
		Tags:      FromSliceTagCore(a.Tags),
		Title:     a.Title,
		Subtitle:  a.Subtitle,
		Content:   a.Content,
		Thumbnail: a.Thumbnail,
		Nsfw:      a.Nsfw,
		Likes:     a.Likes,
		Comments:  FromSliceCommentCore(a.Comments),
		UpdatedAt: a.UpdatedAt,
		CreatedAt: a.CreatedAt,
	}
}

func FromTagCore(tag *articles.TagCore) Tag {
	return Tag{
		Tag: tag.Tag,
	}
}

func FromUserCore(u *articles.UserCore) User {
	return User{
		Username: u.Username,
		Email:    u.Email,
		Name:     u.Name,
	}
}

func FromCommentCore(c *articles.CommentCore) Comment {
	return Comment{
		ID:        c.ID,
		Comment:   c.Comment,
		CreatedAt: c.CreatedAt,
		User:      FromUserCore(&c.User),
	}
}

func FromSliceTagCore(t []articles.TagCore) []Tag {
	tags := make([]Tag, len(t))
	for i, v := range t {
		tags[i] = FromTagCore(&v)
	}
	return tags
}

func FromSliceUserCore(u []articles.UserCore) []User {
	users := make([]User, len(u))
	for i, v := range u {
		users[i] = FromUserCore(&v)
	}
	return users
}

func FromSliceCommentCore(u []articles.CommentCore) []Comment {
	comments := make([]Comment, len(u))
	for i, v := range u {
		comments[i] = FromCommentCore(&v)
	}
	return comments
}
