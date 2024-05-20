package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	DatabaseURL      string
	ServerPort       string
	JwtSecretKey     string
	RefreshSecretKey string
}

var config Config

func LoadConfig() {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	config = Config{
		DatabaseURL:      viper.GetString("DATABASE_URL"),
		ServerPort:       viper.GetString("SERVER_PORT"),
		JwtSecretKey:     viper.GetString("JWT_SECRET_KEY"),
		RefreshSecretKey: viper.GetString("REFRESH_SECRET_KEY"),
	}
}

func GetConfig() Config {
	return config
}
