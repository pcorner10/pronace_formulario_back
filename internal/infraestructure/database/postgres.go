package database

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbHandler *gorm.DB

func Init() *gorm.DB {

	var err error

	// load config from config.yml

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		viper.GetString("database.host"),
		viper.GetString("database.user"),
		viper.GetString("database.password"),
		viper.GetString("database.dbname"),
		viper.GetString("database.port"))

	dbHandler, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	sql, err := dbHandler.DB()
	if err != nil {
		panic(err)
	}

	sql.SetMaxIdleConns(10)
	sql.SetMaxOpenConns(100)
	sql.SetConnMaxIdleTime(30)
	sql.SetConnMaxLifetime(60)

	return dbHandler
}

func GetDB() *gorm.DB {
	return dbHandler
}
