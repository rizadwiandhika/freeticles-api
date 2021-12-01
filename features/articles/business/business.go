package business

import (
	"errors"
	"net/http"

	"github.com/rizadwiandhika/miniproject-backend-alterra/features/articles"
	"github.com/rizadwiandhika/miniproject-backend-alterra/features/bookmarks"
	"github.com/rizadwiandhika/miniproject-backend-alterra/features/reactions"
	"github.com/rizadwiandhika/miniproject-backend-alterra/features/users"
	"github.com/rizadwiandhika/miniproject-backend-alterra/helpers"
	"github.com/rizadwiandhika/miniproject-backend-alterra/third-parties/translate"
)

type articleBusiness struct {
	articleData      articles.IData
	translateApp     translate.IBusiness
	userBusiness     users.IBusiness
	reactionBusiness reactions.IBusiness
	bookmarkBusiness bookmarks.IBusiness
}

func NewBusiness(
	data articles.IData,
	t translate.IBusiness,
	userBusiness users.IBusiness,
	reactionBusiness reactions.IBusiness,
	bookmarkBusiness bookmarks.IBusiness,
) *articleBusiness {
	return &articleBusiness{
		articleData:      data,
		translateApp:     t,
		userBusiness:     userBusiness,
		reactionBusiness: reactionBusiness,
		bookmarkBusiness: bookmarkBusiness,
	}
}

func (ab *articleBusiness) FindArticles(params articles.QueryParams) ([]articles.ArticleCore, error, int) {
	articlesData, err := ab.articleData.SelectArticles(params)
	if err != nil {
		return nil, err, http.StatusInternalServerError
	}

	for i := range articlesData {
		userData, err, _ := ab.userBusiness.FindUserById(articlesData[i].AuthorID)
		if err != nil {
			return nil, err, http.StatusInternalServerError
		}

		totalLikes, err := ab.reactionBusiness.CountTotalArticleLikes(articlesData[i].ID)
		if err != nil {
			return nil, err, http.StatusInternalServerError
		}

		articlesData[i].Author.Username = userData.Username
		articlesData[i].Author.Email = userData.Email
		articlesData[i].Author.Name = userData.Name
		articlesData[i].Likes = totalLikes
	}

	return articlesData, nil, http.StatusOK
}

func (ab *articleBusiness) FindArticleById(id uint) (articles.ArticleCore, error, int) {
	articleData, err := ab.articleData.SelectArticleById(id)
	if err != nil {
		return articles.ArticleCore{}, err, http.StatusInternalServerError
	}
	if articleData.IsNotFound() {
		return articleData, errors.New("Article not found"), http.StatusNotFound
	}

	userData, err, _ := ab.userBusiness.FindUserById(articleData.AuthorID)
	if err != nil {
		return articles.ArticleCore{}, err, http.StatusInternalServerError
	}

	totalLikes, err := ab.reactionBusiness.CountTotalArticleLikes(articleData.ID)
	if err != nil {
		return articles.ArticleCore{}, err, http.StatusInternalServerError
	}

	articleData.Likes = totalLikes
	articleData.Author.Username = userData.Username
	articleData.Author.Email = userData.Email
	articleData.Author.Name = userData.Name

	return articleData, nil, http.StatusOK
}

func (ab *articleBusiness) FindTranslatedArticleById(id uint, lang string) (articles.ArticleCore, error, int) {
	articleData, err := ab.articleData.SelectArticleById(id)
	if err != nil {
		return articles.ArticleCore{}, err, http.StatusInternalServerError
	}
	if articleData.IsNotFound() {
		return articleData, errors.New("Article not found"), http.StatusNotFound
	}

	userData, err, _ := ab.userBusiness.FindUserById(articleData.AuthorID)
	if err != nil {
		return articles.ArticleCore{}, err, http.StatusInternalServerError
	}

	articleData.Author.Username = userData.Username
	articleData.Author.Email = userData.Email
	articleData.Author.Name = userData.Name

	translatedTitle, err := ab.translateApp.Translate(translate.TranslateCore{
		Target: lang,
		Text:   articleData.Title,
	})
	if err != nil {
		return articles.ArticleCore{}, err, http.StatusInternalServerError
	}

	translatedSubtitle, err := ab.translateApp.Translate(translate.TranslateCore{
		Target: lang,
		Text:   articleData.Subtitle,
	})
	if err != nil {
		return articles.ArticleCore{}, err, http.StatusInternalServerError
	}

	articleData.Title = translatedTitle
	articleData.Subtitle = translatedSubtitle

	return articleData, nil, http.StatusOK
}

func (ab *articleBusiness) FindUserArticles(username string) ([]articles.ArticleCore, error, int) {
	user, err, _ := ab.userBusiness.FindUserByUsername(username)
	if err != nil {
		return nil, err, http.StatusInternalServerError
	}

	userArticles, err := ab.articleData.SelectArticlesByAuthorId(user.ID)
	if err != nil {
		return nil, err, http.StatusInternalServerError
	}

	for i := range userArticles {
		totalLikes, err := ab.reactionBusiness.CountTotalArticleLikes(userArticles[i].ID)
		if err != nil {
			return nil, err, http.StatusInternalServerError
		}
		userArticles[i].Likes = totalLikes
	}

	return userArticles, nil, http.StatusOK
}

func (ab *articleBusiness) RemoveArticleById(id uint) (error, int) {
	err := ab.articleData.DeleteArticleTags(id)
	if err != nil {
		return err, http.StatusInternalServerError
	}

	err = ab.articleData.DeleteArticleById(id)
	if err != nil {
		return err, http.StatusInternalServerError
	}

	go ab.reactionBusiness.RemoveCommentsByArticleId(id)
	go ab.reactionBusiness.RemoveLikesByArticleId(id)
	go ab.reactionBusiness.RemoveReportsByArticleId(id)

	return nil, http.StatusAccepted
}

func (ab *articleBusiness) RemoveUserArticles(userID uint) (error, int) {
	userArticles, err := ab.articleData.SelectArticlesByAuthorId(userID)
	if err != nil {
		return err, http.StatusInternalServerError
	}

	articleIDs := make([]uint, len(userArticles))
	for i, a := range userArticles {
		articleIDs[i] = a.ID
	}

	err = ab.articleData.DeleteTagByArticleIds(articleIDs)
	if err != nil {
		return err, http.StatusInternalServerError
	}

	err = ab.articleData.DeleteArticlesByUserId(userID)
	if err != nil {
		return err, http.StatusInternalServerError
	}
	return nil, http.StatusOK
}

func (ab *articleBusiness) CreateArticle(article articles.ArticleCore) (articles.ArticleCore, error, int) {
	createdArticle, err := ab.articleData.InsertArticle(article)
	if err != nil {
		return article, err, http.StatusInternalServerError
	}

	return createdArticle, nil, http.StatusOK
}

func (ab *articleBusiness) EditArticle(updatedArticle articles.ArticleCore) (articles.ArticleCore, error, int) {
	previousArticle, err := ab.articleData.SelectArticleById(updatedArticle.ID)
	if err != nil {
		return updatedArticle, err, http.StatusInternalServerError
	}
	if previousArticle.IsNotFound() {
		return updatedArticle, errors.New("Article not found"), http.StatusNotFound
	}

	if previousArticle.Thumbnail != updatedArticle.Thumbnail {
		err = helpers.DeleteFile(previousArticle.Thumbnail)
		if err != nil {
			return updatedArticle, err, http.StatusInternalServerError
		}
	}

	// Edit article is only allowed for title, subtitle, thumbnail, nsfw, and content
	previousArticle.Title = updatedArticle.Title
	previousArticle.Subtitle = updatedArticle.Subtitle
	previousArticle.Thumbnail = updatedArticle.Thumbnail
	previousArticle.Content = updatedArticle.Content
	previousArticle.Nsfw = updatedArticle.Nsfw
	previousArticle.Tags = updatedArticle.Tags

	editedArticle, err := ab.articleData.UpdateArticle(previousArticle)
	if err != nil {
		return updatedArticle, err, http.StatusInternalServerError
	}

	return editedArticle, nil, http.StatusOK
}
