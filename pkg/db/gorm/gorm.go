package db

import (
	"fmt"
	"log"
	"os"
	"pronaces_back/pkg/domain"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type gormStore struct {
	db *gorm.DB
}

func NewGormStore() (domain.FormDB, error) {

	dbHandler, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	if err != nil {
		return nil, err
	}

	if os.Getenv("MIGRATE") == "true" {
		Migrate(dbHandler)
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
	return &gormStore{db: dbHandler}, nil
}
