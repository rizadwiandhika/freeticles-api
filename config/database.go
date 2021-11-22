package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		ENV.DB_USERNAME,
		ENV.DB_PASSWORD,
		ENV.DB_NAME,
	)

	loggerWriter := log.New(os.Stdout, "\r\n", log.LstdFlags)
	loggerConfig := logger.Config{
		SlowThreshold: time.Second,
		LogLevel:      logger.Silent,
		Colorful:      true,
	}

	connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.New(loggerWriter, loggerConfig)})
	if err != nil {
		panic(err)
	}

	DB = connection
}
