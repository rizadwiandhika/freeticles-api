package presentation

import (
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/rizadwiandhika/miniproject-backend-alterra/features/reactions"
	"github.com/rizadwiandhika/miniproject-backend-alterra/features/reactions/presentation/request"
	"github.com/rizadwiandhika/miniproject-backend-alterra/features/reactions/presentation/response"
)

type any interface{}
type json map[string]any

type ReactionPresentation struct {
	reactionBusiness reactions.IBusiness
}

func NewPresentation(rb reactions.IBusiness) *ReactionPresentation {
	return &ReactionPresentation{
		reactionBusiness: rb,
	}
}

func (rp *ReactionPresentation) GetArticleComments(c echo.Context) error {
	var articleID uint
	echo.PathParamsBinder(c).Uint("id", &articleID)

	comments, err, status := rp.reactionBusiness.FindCommentsByArticleId(articleID)
	if err != nil {
		return c.JSON(status, json{
			"message": "Failed getting comments",
			"error":   err.Error(),
		})
	}

	return c.JSON(status, json{
		"message":  "success fetching comments",
		"comments": response.FromSliceCommentCore(comments),
	})
}

func (rp *ReactionPresentation) PostLike(c echo.Context) error {
	user := c.Get("user").(jwt.MapClaims)
	username := user["username"].(string)
	body := request.Request{}

	err := c.Bind(&body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Something went wrong")
	}

	err, status := rp.reactionBusiness.PostLike(username, body.ArticleID)
	if err != nil {
		return c.JSON(status, json{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, json{
		"message": "article liked!",
	})
}

func (rp *ReactionPresentation) DeleteLike(c echo.Context) error {
	user := c.Get("user").(jwt.MapClaims)
	username := user["username"].(string)
	body := request.Request{}

	err := c.Bind(&body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Something went wrong")
	}

	err, status := rp.reactionBusiness.Dislike(username, body.ArticleID)
	if err != nil {
		return c.JSON(status, json{
			"message": err.Error(),
		})
	}

	return c.JSON(status, json{
		"message": "article unliked!",
	})
}

func (rp *ReactionPresentation) PostComment(c echo.Context) error {
	user := c.Get("user").(jwt.MapClaims)
	username := user["username"].(string)
	body := request.Request{}

	err := c.Bind(&body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Something went wrong")
	}

	err, status := rp.reactionBusiness.PostComment(username, body.ArticleID, body.Commentar)
	if err != nil {
		return c.JSON(status, json{
			"message": err.Error(),
		})
	}

	return c.JSON(status, json{
		"message": "comment posted!",
	})
}

func (rp *ReactionPresentation) PostReport(c echo.Context) error {
	user := c.Get("user").(jwt.MapClaims)
	username := user["username"].(string)
	body := request.Request{}

	err := c.Bind(&body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Something went wrong")
	}

	err, status := rp.reactionBusiness.ReportArticle(username, body.ArticleID, body.ReportTypeID)
	if err != nil {
		return c.JSON(status, json{
			"message": "Failed reporting articles",
			"error":   err.Error(),
		})
	}

	return c.JSON(status, json{
		"message": "article reported!",
	})
}
