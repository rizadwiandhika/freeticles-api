package data_test

import (
	"database/sql"
	"fmt"
	"os"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/rizadwiandhika/miniproject-backend-alterra/features/articles"
	"github.com/rizadwiandhika/miniproject-backend-alterra/features/articles/data"
)

var (
	gormDB  *gorm.DB
	sqlMock sqlmock.Sqlmock

	repository  articles.IData
	articleCore articles.ArticleCore
)

func TestMain(m *testing.M) {
	var db *sql.DB
	var err error

	db, sqlMock, err = sqlmock.New()
	assert.Equal(&testing.T{}, err, nil)

	mysqlConfig := mysql.Config{Conn: db, SkipInitializeWithVersion: true}
	// dbLogger := logger.Default.LogMode(logger.Info)

	gormDB, err = gorm.Open(mysql.New(mysqlConfig), &gorm.Config{})
	assert.Equal(&testing.T{}, err, nil)

	repository = data.NewMySQLRepository(gormDB)
	articleCore = articles.ArticleCore{
		ID:       1,
		AuthorID: 1,
		Tags: []articles.TagCore{
			{Tag: "tag1"},
			{Tag: "tag2"},
			{Tag: "tag3"},
		},
		Title:     "title",
		Subtitle:  "subtitle",
		Content:   "content",
		Thumbnail: "google.com/img.jpg",
		Nsfw:      false,
	}

	os.Exit(m.Run())
}

func TestSelectArticleById(t *testing.T) {
	t.Run("valid - SelectArticleById", func(t *testing.T) {
		sqlMock.ExpectQuery(
			regexp.QuoteMeta("SELECT * FROM `articles` WHERE `articles`.`id` = ? ORDER BY `articles`.`id` LIMIT 1"),
		).WillReturnRows(
			sqlmock.NewRows(
				[]string{"id", "author_id", "title", "subtitle", "content", "thumbnail", "nsfw"},
			).AddRow(
				articleCore.ID, articleCore.AuthorID, articleCore.Title, articleCore.Subtitle, articleCore.Content, articleCore.Thumbnail, articleCore.Nsfw,
			),
		)

		sqlMock.ExpectQuery(
			regexp.QuoteMeta("SELECT * FROM `tags` WHERE `tags`.`article_id` = ?"),
		).WillReturnRows(
			sqlmock.NewRows(
				[]string{"id", "article_id", "tag"},
			).AddRow(
				1, 1, "tag1",
			).AddRow(
				2, 1, "tag2",
			),
		)

		result, err := repository.SelectArticleById(1)
		assert.Equal(t, articleCore.ID, result.ID)
		assert.Nil(t, err)
	})

	t.Run("valid - empty user when record not found", func(t *testing.T) {
		sqlMock.ExpectQuery(
			regexp.QuoteMeta("SELECT * FROM `articles` WHERE `articles`.`id` = ? ORDER BY `articles`.`id` LIMIT 1"),
		).WillReturnRows(
			sqlmock.NewRows(
				[]string{"id", "author_id", "title", "subtitle", "content", "thumbnail", "nsfw"},
			),
		)

		sqlMock.ExpectQuery(
			regexp.QuoteMeta("SELECT * FROM `tags` WHERE `tags`.`article_id` = ?"),
		).WillReturnRows(
			sqlmock.NewRows(
				[]string{"id", "article_id", "tag"},
			),
		)

		fetchedArticle, err := repository.SelectArticleById(1)

		assert.Equal(t, uint(0), fetchedArticle.ID)
		assert.Nil(t, err)
	})
}

func TestInsertArticle(t *testing.T) {
	t.Run("valid - InsertArticle", func(t *testing.T) {
		// sqlMock.ExpectPrepare(
		// 	"^INSERT INTO *",
		// ).ExpectExec().WithArgs()

		/* sqlMock.ExpectExec(
			"INSERT INTO `tags` (.+)",
		).WithArgs(
			1, "tag1",
		).WillReturnResult(sqlmock.NewResult(1, 1))

		sqlMock.ExpectExec(
			"INSERT INTO `tags` (.+)",
		).WithArgs(
			1, "tag2",
		).WillReturnResult(sqlmock.NewResult(1, 1))

		sqlMock.ExpectExec(
			"INSERT INTO `tags` (.+)",
		).WithArgs(
			1, "tag3",
		).WillReturnResult(sqlmock.NewResult(1, 1)) */

		/* sqlMock.ExpectExec(
			"INSERT INTO `tags` (.+)",
		).WithArgs(
			1, "tag1", sqlmock.AnyArg(), sqlmock.AnyArg(),
		).WithArgs(
			1, "tag2", sqlmock.AnyArg(), sqlmock.AnyArg(),
		).WithArgs(
			1, "tag3", sqlmock.AnyArg(), sqlmock.AnyArg(),
		).WillReturnResult(sqlmock.NewResult(3, 3)) */

		newArticle := articles.ArticleCore{
			AuthorID: 1,
			Tags: []articles.TagCore{
				{Tag: "tag1"},
				{Tag: "tag2"},
				{Tag: "tag3"},
			},
			Title:     "title",
			Subtitle:  "subtitle",
			Content:   "content",
			Thumbnail: "google.com/img.jpg",
			Nsfw:      false,
		}

		sqlMock.ExpectBegin()

		sqlMock.ExpectExec(
			"INSERT INTO `articles` (.+)",
		).WithArgs(
			newArticle.AuthorID,
			newArticle.Title,
			newArticle.Subtitle,
			newArticle.Content,
			newArticle.Thumbnail,
			newArticle.Nsfw,
			sqlmock.AnyArg(),
			sqlmock.AnyArg(),
		).WillReturnResult(sqlmock.NewResult(1, 1))

		sqlMock.ExpectExec(
			"INSERT INTO `tags` (.+)",
		).WithArgs(
			1, "tag1", sqlmock.AnyArg(), sqlmock.AnyArg(),
			1, "tag2", sqlmock.AnyArg(), sqlmock.AnyArg(),
			1, "tag3", sqlmock.AnyArg(), sqlmock.AnyArg(),
		).WillReturnResult(sqlmock.NewResult(3, 3))

		sqlMock.ExpectCommit()

		createdArticle, err := repository.InsertArticle(newArticle)

		fmt.Println("createdArticle: ", createdArticle.ID)
		assert.Greater(t, createdArticle.ID, uint(0))
		assert.Equal(t, newArticle.AuthorID, createdArticle.AuthorID)
		assert.Nil(t, err)
	})
}
