package business

import (
	"errors"
	"net/http"

	"github.com/rizadwiandhika/miniproject-backend-alterra/features/articles"
	"github.com/rizadwiandhika/miniproject-backend-alterra/features/reactions"
	"github.com/rizadwiandhika/miniproject-backend-alterra/features/users"
	images "github.com/rizadwiandhika/miniproject-backend-alterra/third-parties/image"
)

type reactionBusiness struct {
	reactionData    reactions.IData
	imageAnalyzer   images.IBusiness
	userBusniess    users.IBusiness
	articleBusiness articles.IBusiness
}

func NewBusiness(rd reactions.IData, ia images.IBusiness, ub users.IBusiness, ab articles.IBusiness) *reactionBusiness {
	return &reactionBusiness{
		reactionData:    rd,
		imageAnalyzer:   ia,
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
			ID:       user.ID,
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

func (rp *reactionBusiness) CountTotalArticleLikes(articleId uint) (int, error) {
	totalLikes, err := rp.reactionData.SelectCountLikes(articleId)
	if err != nil {
		return 0, err
	}
	return totalLikes, nil
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

	const SEXUAL_CONTENT_REPORT = 2
	if reportTypeId == SEXUAL_CONTENT_REPORT {
		go rb.followupNSFWReport(*article)
	}

	return nil, http.StatusCreated
}

func (rb *reactionBusiness) RemoveCommentsByArticleId(id uint) error {
	return rb.reactionData.DeleteCommentsByArticleId(id)
}
func (rb *reactionBusiness) RemoveLikesByArticleId(id uint) error {
	return rb.reactionData.DeleteLikesByArticleId(id)
}
func (rb *reactionBusiness) RemoveReportsByArticleId(id uint) error {
	return rb.reactionData.DeleteReportsByArticleId(id)
}

func (rb *reactionBusiness) RemoveCommentsByUserId(id uint) error {
	return rb.reactionData.DeleteCommentsByUserId(id)
}
func (rb *reactionBusiness) RemoveLikesByUserId(id uint) error {
	return rb.reactionData.DeleteLikesByUserId(id)
}
func (rb *reactionBusiness) RemoveReportsByUserId(id uint) error {
	return rb.reactionData.DeleteReportsByUserId(id)
}

func (rb *reactionBusiness) findUserAndArticle(username string, articleId uint) (*users.UserCore, *articles.ArticleCore) {
	userChannnel := make(chan *users.UserCore)
	articleChannnel := make(chan *articles.ArticleCore)

	go func() {
		user, err, _ := rb.userBusniess.FindUserByUsername(username)
		if err != nil {
			userChannnel <- nil
			return
		}
		userChannnel <- &user
	}()

	go func() {
		article, err, _ := rb.articleBusiness.FindArticleById(articleId)
		if err != nil {
			articleChannnel <- nil
			return
		}
		articleChannnel <- &article
	}()

	user := <-userChannnel
	article := <-articleChannnel

	return user, article
}

func (rb *reactionBusiness) followupNSFWReport(article articles.ArticleCore) {
	if article.Thumbnail == "" {
		return
	}

	isNSFW, err := rb.imageAnalyzer.IsNSFW(article.Thumbnail)
	if err != nil || !isNSFW {
		return
	}

	article.Nsfw = true
	_, _, _ = rb.articleBusiness.EditArticle(article)
}
