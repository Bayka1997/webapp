package app

import (
	database "webapp/config"
	"webapp/helper"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func Run() {
	// load env
	err := godotenv.Load()
	helper.ErrorPanic(err)
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	viper.AutomaticEnv()

	// connect database
	database.DatabaseConnection()
}
