package migration

import (
	"github.com/rizadwiandhika/miniproject-backend-alterra/config"
	article "github.com/rizadwiandhika/miniproject-backend-alterra/features/articles/data"
	user "github.com/rizadwiandhika/miniproject-backend-alterra/features/users/data"
	"golang.org/x/crypto/bcrypt"
)

func AutoMigrate() {
	if err := config.DB.Exec("DROP TABLE IF EXISTS users").Error; err != nil {
		panic(err)
	}
	if err := config.DB.Exec("DROP TABLE IF EXISTS tags").Error; err != nil {
		panic(err)
	}
	if err := config.DB.Exec("DROP TABLE IF EXISTS articles").Error; err != nil {
		panic(err)
	}

	err := config.DB.AutoMigrate(&user.User{}, &article.Article{}, &article.Tag{})
	if err != nil {
		panic(err)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("riza123"), 14)
	if err != nil {
		panic(err)
	}

	usr1 := user.User{
		Username: "riza.dwii",
		Email:    "rizadwiandhika@mail.com",
		Password: string(hashedPassword),
		Name:     "Riza Dwi Andhika",
	}
	usr2 := user.User{
		Username: "hernowoari",
		Email:    "owo@mail.com",
		Password: string(hashedPassword),
		Name:     "Hernowo Ari Sutanto",
	}

	tags := []article.Tag{
		{ArticleID: 1, Tag: "anime"},
		{ArticleID: 1, Tag: "sport"},
		{ArticleID: 1, Tag: "fun"},
	}
	arr := article.Article{
		Title:    "Test Article",
		Content:  `Some random content`,
		AuthorID: 1,
		Tags:     tags,
	}
	arr2 := article.Article{
		Title:    "Another Article",
		Content:  `Again... random content`,
		AuthorID: 2,
		Tags:     tags,
	}

	err = config.DB.Create(&usr1).Error
	if err != nil {
		panic(err)
	}
	err = config.DB.Create(&usr2).Error
	if err != nil {
		panic(err)
	}

	err = config.DB.Create(&arr).Error
	if err != nil {
		panic(err)
	}
	err = config.DB.Create(&arr2).Error
	if err != nil {
		panic(err)
	}

}
