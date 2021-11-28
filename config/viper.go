package config

import (
	"path/filepath"

	"github.com/spf13/viper"
)

type Env struct {
	JWT_SECRET  string `mapstructure:"JWT_SECRET"`
	DB_NAME     string `mapstructure:"DB_NAME"`
	DB_USERNAME string `mapstructure:"DB_USERNAME"`
	DB_PASSWORD string `mapstructure:"DB_PASSWORD"`
	DB_HOST     string `mapstructure:"DB_HOST"`
	DB_PORT     string `mapstructure:"DB_PORT"`
}

var ENV Env

func LoadConfig(path string) error {
	viper.AddConfigPath(path)  // __dir of main.go
	viper.SetConfigName("app") // .env file name
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	error := viper.Unmarshal(&ENV)
	return error
}

func LoadEnv() {
	envFile := filepath.Join(".env")
	viper.SetConfigFile(envFile)

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(&ENV); err != nil {
		panic(err)
	}
}
