package data

import (
	"errors"
	"fmt"

	"github.com/rizadwiandhika/miniproject-backend-alterra/features/reactions"
	"gorm.io/gorm"
)

type reactionRepository struct {
	db *gorm.DB
}

func NewMySQLRepository(db *gorm.DB) *reactionRepository {
	return &reactionRepository{
		db: db,
	}
}

func (rr *reactionRepository) SelectLike(like reactions.LikeCore) (reactions.LikeCore, error) {
	fetchedLike := Like{}
	err := rr.db.Where("user_id = ? AND article_id = ?", like.UserID, like.ArticleID).First(&fetchedLike).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return fetchedLike.toLikeCore(), err
	}
	return fetchedLike.toLikeCore(), nil
}

func (rr *reactionRepository) SelectCommentsByArticleId(articleId uint) ([]reactions.CommentCore, error) {
	comments := SliceComment{}
	err := rr.db.Where("article_id = ?", articleId).Find(&comments).Error
	if err != nil {
		return nil, err
	}
	return comments.toSliceCommentCore(), nil
}

func (rr *reactionRepository) SelectUserReport(userId uint, articleId uint) (reactions.ReportCore, error) {
	fetchedReport := Report{}
	err := rr.db.Where("user_id = ? AND article_id = ?", userId, articleId).First(&fetchedReport).Error
	if err != nil && err == gorm.ErrUnsupportedRelation {
		fmt.Println("unsupported!")
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		return fetchedReport.toReportCore(), err
	}
	return fetchedReport.toReportCore(), nil
}

func (rr *reactionRepository) InsertLike(like reactions.LikeCore) error {
	newLike := Like{UserID: like.UserID, ArticleID: like.ArticleID}
	return rr.db.Create(&newLike).Error
}

func (rr *reactionRepository) InsertComment(comment reactions.CommentCore) error {
	newComment := Comment{
		UserID:    comment.UserID,
		ArticleID: comment.ArticleID,
		Commentar: comment.Commentar,
	}
	return rr.db.Create(&newComment).Error
}

func (rr *reactionRepository) InsertReport(report reactions.ReportCore) error {
	newReport := Report{
		UserID:       report.UserID,
		ArticleID:    report.ArticleID,
		ReportTypeID: report.ReportTypeID,
	}
	return rr.db.Create(&newReport).Error
}

func (rr *reactionRepository) DeleteLikeById(likeId uint) error {
	return rr.db.Delete(&Like{}, likeId).Error
}
