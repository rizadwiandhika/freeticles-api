package main

import (
	"github.com/rizadwiandhika/miniproject-backend-alterra/config"
)

func main() {
	config.LoadEnv()
	config.ConnectDB()

	// e := routes.Setup()
	// log.Fatalln(e.Start(":8080"))
}
