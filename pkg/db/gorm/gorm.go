package db

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbHandler *gorm.DB

func NewGormStore() *gorm.DB {

	var err error

	// load config from config.yml

	dsn := fmt.Sprintf("postgres://%s:%s@%s:5432/%s",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"))

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

	fmt.Println("Successfully connected to database!")
	return dbHandler
}

func GetDB() *gorm.DB {
	return dbHandler
}
