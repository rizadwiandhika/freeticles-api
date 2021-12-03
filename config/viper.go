package config

import (
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type Env struct {
	JWT_SECRET           string `mapstructure:"JWT_SECRET"`
	CLOUDMERSIVE_API_KEY string `mapstructure:"CLOUDMERSIVE_API_KEY"`
	DB_NAME              string `mapstructure:"DB_NAME"`
	DB_USERNAME          string `mapstructure:"DB_USERNAME"`
	DB_PASSWORD          string `mapstructure:"DB_PASSWORD"`
	DB_HOST              string `mapstructure:"DB_HOST"`
	DB_PORT              string `mapstructure:"DB_PORT"`
	PORT                 string
}

var ENV Env
var WORKING_DIR string

func LoadEnv() {
	envFile := filepath.Join(".env")
	viper.SetConfigFile(envFile)

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(&ENV); err != nil {
		panic(err)
	}

	ENV.PORT = os.Getenv("PORT")
	if ENV.PORT == "" {
		ENV.PORT = "8080"
	}

	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	WORKING_DIR = pwd
}
