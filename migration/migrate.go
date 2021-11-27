package migration

import (
	"encoding/json"

	"github.com/rizadwiandhika/miniproject-backend-alterra/config"
	article "github.com/rizadwiandhika/miniproject-backend-alterra/features/articles/data"
	reaction "github.com/rizadwiandhika/miniproject-backend-alterra/features/reactions/data"
	user "github.com/rizadwiandhika/miniproject-backend-alterra/features/users/data"
	"golang.org/x/crypto/bcrypt"
)

func AutoMigrate() {
	if err := config.DB.Exec("DROP TABLE IF EXISTS tags").Error; err != nil {
		panic(err)
	}
	if err := config.DB.Exec("DROP TABLE IF EXISTS comments").Error; err != nil {
		panic(err)
	}
	if err := config.DB.Exec("DROP TABLE IF EXISTS likes").Error; err != nil {
		panic(err)
	}
	if err := config.DB.Exec("DROP TABLE IF EXISTS reports").Error; err != nil {
		panic(err)
	}
	if err := config.DB.Exec("DROP TABLE IF EXISTS report_types").Error; err != nil {
		panic(err)
	}
	if err := config.DB.Exec("DROP TABLE IF EXISTS articles").Error; err != nil {
		panic(err)
	}
	if err := config.DB.Exec("DROP TABLE IF EXISTS users").Error; err != nil {
		panic(err)
	}

	err := config.DB.AutoMigrate(
		&user.User{},
		&article.Article{},
		&article.Tag{},
		&reaction.Comment{},
		&reaction.Like{},
		&reaction.ReportType{},
		&reaction.Report{},
	)

	if err != nil {
		panic(err)
	}

	populateDBWithDummyData()
}

func populateDBWithDummyData() {
	pw1, err := bcrypt.GenerateFromPassword([]byte("riza123"), 14)
	if err != nil {
		panic(err)
	}
	pw2, err := bcrypt.GenerateFromPassword([]byte("hernowo123"), 14)
	if err != nil {
		panic(err)
	}
	pw3, err := bcrypt.GenerateFromPassword([]byte("hammim123"), 14)
	if err != nil {
		panic(err)
	}

	usr1 := user.User{
		Username: "riza.dwii",
		Email:    "riza@mail.com",
		Role:     "admin",
		Password: string(pw1),
		Name:     "Riza Dwi Andhika",
	}
	usr2 := user.User{
		Username: "hernowoari",
		Email:    "hernowo@mail.com",
		Password: string(pw2),
		Name:     "Hernowo Ari Sutanto",
	}
	usr3 := user.User{
		Username: "hammim",
		Email:    "hammim@mail.com",
		Password: string(pw3),
		Name:     "Hammim Eka",
	}

	err = config.DB.Create(&usr1).Error
	if err != nil {
		panic(err)
	}
	err = config.DB.Create(&usr2).Error
	if err != nil {
		panic(err)
	}
	err = config.DB.Create(&usr3).Error
	if err != nil {
		panic(err)
	}

	arr1 := article.Article{
		Title:    "Test Article",
		Content:  `Some random content`,
		AuthorID: 1,
		Tags: []article.Tag{
			{Tag: "anime"},
			{Tag: "sport"},
			{Tag: "fun"},
		},
	}
	arr2 := article.Article{
		Title:    "Another Article",
		Content:  `Again... random content`,
		AuthorID: 2,
		Tags: []article.Tag{
			{Tag: "car"},
			{Tag: "education"},
			{Tag: "hobby"},
		},
	}

	err = config.DB.Create(&arr1).Error
	if err != nil {
		panic(err)
	}
	err = config.DB.Create(&arr2).Error
	if err != nil {
		panic(err)
	}

	rt1 := reaction.ReportType{
		Name:        "spam",
		Description: "Article is considered as spam",
	}
	rt2 := reaction.ReportType{
		Name:        "sexual content",
		Description: "Article content may not appropriate for certain users",
	}
	rt3 := reaction.ReportType{
		Name:        "just don't like",
		Description: "Users somehow just don't like the article content",
	}

	err = config.DB.Create(&rt1).Error
	if err != nil {
		panic(err)
	}
	err = config.DB.Create(&rt2).Error
	if err != nil {
		panic(err)
	}
	err = config.DB.Create(&rt3).Error
	if err != nil {
		panic(err)
	}

}

func outputAsJSON(v interface{}) string {
	JSON, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(JSON)
}
