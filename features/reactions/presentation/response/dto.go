package response

import (
	"time"

	"github.com/rizadwiandhika/miniproject-backend-alterra/features/reactions"
)

type LikeResponse struct {
	Message string `json:"message"`
}

type Comment struct {
	Commentar string    `json:"commentar"`
	User      User      `json:"user"`
	CreatedAt time.Time `json:"createdAt"`
}

type User struct {
	ID       uint   `json:"userId"`
	Username string `json:"username"`
	Name     string `json:"name"`
}

func FromCommentCore(c *reactions.CommentCore) Comment {
	return Comment{
		Commentar: c.Commentar,
		CreatedAt: c.CreatedAt,
		User: User{
			ID:       c.User.ID,
			Username: c.User.Username,
			Name:     c.User.Name,
		},
	}
}

func FromSliceCommentCore(c []reactions.CommentCore) []Comment {
	var comments []Comment
	for _, v := range c {
		comments = append(comments, FromCommentCore(&v))
	}
	return comments
}
