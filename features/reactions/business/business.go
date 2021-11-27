package business

import (
	"errors"
	"net/http"

	"github.com/rizadwiandhika/miniproject-backend-alterra/features/articles"
	"github.com/rizadwiandhika/miniproject-backend-alterra/features/reactions"
	"github.com/rizadwiandhika/miniproject-backend-alterra/features/users"
)

type reactionBusiness struct {
	reactionData    reactions.IData
	userBusniess    users.IBusiness
	articleBusiness articles.IBusiness
}

func NewBusiness(rd reactions.IData, ub users.IBusiness, ab articles.IBusiness) *reactionBusiness {
	return &reactionBusiness{
		reactionData:    rd,
		userBusniess:    ub,
		articleBusiness: ab,
	}
}

func (rb *reactionBusiness) FindCommentsByArticleId(articleId uint) ([]reactions.CommentCore, error, int) {
	comments, err := rb.reactionData.SelectCommentsByArticleId(articleId)
	if err != nil {
		return nil, err, http.StatusInternalServerError
	}

	userIds := []uint{}
	for _, c := range comments {
		userIds = append(userIds, c.UserID)
	}

	users, err, _ := rb.userBusniess.FindUsersByIds(userIds)
	if err != nil {
		return nil, err, http.StatusInternalServerError
	}

	usersMap := make(map[uint]reactions.UserCore)
	for _, user := range users {
		usersMap[user.ID] = reactions.UserCore{
			Username: user.Username,
			Name:     user.Name,
		}
	}

	for i := range comments {
		userId := comments[i].UserID
		comments[i].User = usersMap[userId]
	}

	return comments, nil, http.StatusOK
}

func (rb *reactionBusiness) PostLike(username string, articleId uint) (error, int) {
	user, article := rb.findUserAndArticle(username, articleId)
	if user == nil || article == nil {
		return errors.New("User or article not found"), http.StatusNotFound
	}

	like := reactions.LikeCore{
		UserID:    user.ID,
		ArticleID: article.ID,
	}
	fetchedLike, err := rb.reactionData.SelectLike(like)
	if err != nil {
		return err, http.StatusInternalServerError
	}
	if fetchedLike.ID != 0 {
		return errors.New("Article already liked"), http.StatusUnprocessableEntity
	}

	newLike := reactions.LikeCore{
		UserID:    user.ID,
		ArticleID: article.ID,
	}

	err = rb.reactionData.InsertLike(newLike)
	if err != nil {
		return err, http.StatusInternalServerError
	}

	return nil, http.StatusCreated
}

func (rb *reactionBusiness) PostComment(username string, articleId uint, commentar string) (error, int) {
	user, article := rb.findUserAndArticle(username, articleId)
	if user == nil || article == nil {
		return errors.New("User or article not found"), http.StatusNotFound
	}

	newComment := reactions.CommentCore{
		UserID:    user.ID,
		ArticleID: article.ID,
		Commentar: commentar,
	}
	err := rb.reactionData.InsertComment(newComment)
	if err != nil {
		return err, http.StatusInternalServerError
	}

	return nil, http.StatusCreated
}

func (rb *reactionBusiness) Dislike(username string, articleId uint) (error, int) {
	user, article := rb.findUserAndArticle(username, articleId)

	like := reactions.LikeCore{
		UserID:    user.ID,
		ArticleID: article.ID,
	}
	fetchedLike, err := rb.reactionData.SelectLike(like)
	if err != nil {
		return err, http.StatusInternalServerError
	}
	if fetchedLike.ID == 0 {
		return errors.New("Article is not liked"), http.StatusUnprocessableEntity
	}

	err = rb.reactionData.DeleteLikeById(fetchedLike.ID)
	if err != nil {
		return err, http.StatusInternalServerError
	}

	return nil, http.StatusOK
}

func (rb *reactionBusiness) ReportArticle(username string, articleId uint, reportTypeId uint) (error, int) {
	user, article := rb.findUserAndArticle(username, articleId)
	if user == nil || article == nil {
		return errors.New("User or article not found"), http.StatusNotFound
	}

	report := reactions.ReportCore{
		UserID:       user.ID,
		ArticleID:    article.ID,
		ReportTypeID: reportTypeId,
	}
	existingReport, err := rb.reactionData.SelectUserReport(report.UserID, report.ArticleID)
	if err != nil {
		return err, http.StatusInternalServerError
	}
	if existingReport.ID != 0 {
		return errors.New("Article already reported"), http.StatusUnprocessableEntity
	}

	err = rb.reactionData.InsertReport(report)
	if err != nil {
		return err, http.StatusInternalServerError
	}

	return nil, http.StatusCreated
}

func (rb *reactionBusiness) findUserAndArticle(username string, articleId uint) (*reactions.UserCore, *reactions.ArticleCore) {
	userChannnel := make(chan *reactions.UserCore)
	articleChannnel := make(chan *reactions.ArticleCore)

	go func() {
		user, err, _ := rb.userBusniess.FindUserByUsername(username)
		if err != nil {
			userChannnel <- nil
			return
		}
		userChannnel <- &reactions.UserCore{
			ID:       user.ID,
			Username: user.Username,
			Name:     user.Name,
		}
	}()

	go func() {
		article, err, _ := rb.articleBusiness.FindArticleById(articleId)
		if err != nil {
			articleChannnel <- nil
			return
		}
		articleChannnel <- &reactions.ArticleCore{
			ID:    article.ID,
			Title: article.Title,
		}
	}()

	user := <-userChannnel
	article := <-articleChannnel

	return user, article
}
