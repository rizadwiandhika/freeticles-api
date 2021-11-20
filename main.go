package main

import (
	"log"

	"github.com/rizadwiandhika/miniproject-backend-alterra/config"
	"github.com/rizadwiandhika/miniproject-backend-alterra/migration"
	"github.com/rizadwiandhika/miniproject-backend-alterra/routes"
)

func main() {
	config.LoadEnv()
	config.ConnectDB()
	migration.AutoMigrate()

	e := routes.Setup()
	log.Fatalln(e.Start(":8080"))
}
