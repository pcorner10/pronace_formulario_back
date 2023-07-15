package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"pronaces_back/pkg/domain"

	"cloud.google.com/go/cloudsqlconn"
	"cloud.google.com/go/cloudsqlconn/postgres/pgxv4"
)

type psxStore struct {
	db      *sql.DB
	cleanup func() error
}

func NewPsxStore() (domain.FormDB, error) {
	cleanup, err := pgxv4.RegisterDriver("cloudsql-postgres", cloudsqlconn.WithIAMAuthN())
	if err != nil {
		log.Fatalf("Error on pgxv4.RegisterDriver: %v", err)
	}

	dsn := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable", os.Getenv("INSTANCE_CONNECTION_NAME"), os.Getenv("DB_USER"), os.Getenv("DB_NAME"))
	db, err := sql.Open("cloudsql-postgres", dsn)
	if err != nil {
		log.Fatalf("Error on sql.Open: %v", err)
	}

	fmt.Println("Successfully connected to database!")
	return &psxStore{db: db, cleanup: cleanup}, nil
}

