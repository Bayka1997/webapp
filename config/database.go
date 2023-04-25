package database

import (
	"fmt"
	"webapp/helper"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConfig struct {
	Host     string `mapstructure:"DB_HOST"`
	Port     int64  `mapstructure:"DB_PORT"`
	Username string `mapstructure:"DB_USERNAME"`
	Password string `mapstructure:"DB_PASSWORD"`
	Database string `mapstructure:"DB_DATABASE"`
}

func DatabaseConnection() *gorm.DB {
	conf := getConf()
	sqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", conf.Host, conf.Port, conf.Username, conf.Password, conf.Database)
	db, err := gorm.Open(postgres.Open(sqlInfo), &gorm.Config{})
	helper.ErrorPanic(err)

	return db
}

func getConf() *DBConfig {
	conf := &DBConfig{}
	err := viper.Unmarshal(conf)
	helper.ErrorPanic(err)

	return conf
}
