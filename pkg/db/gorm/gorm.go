package db

import (
	"database/sql"
	"fmt"
	"os"
	"pronaces_back/pkg/domain"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type gormStore struct {
	db *gorm.DB
}

func NewGormStore() (domain.FormDB, error) {

	dbUser := os.Getenv("DB_USER")                                  // e.g. 'my-db-user'
	dbPwd := os.Getenv("DB_PASS")                                   // e.g. 'my-db-password'
	dbName := os.Getenv("DB_NAME")                                  // e.g. 'my-database'
	instanceConnectionName := os.Getenv("INSTANCE_CONNECTION_NAME") // e.g. 'project:region:instance'
	//usePrivate := os.Getenv("PRIVATE_IP")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=America/Mexico_City",
		instanceConnectionName,
		dbUser,
		dbPwd,
		dbName,
	)

	fmt.Println(dsn)

	conn, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}

	dbHandler, err := gorm.Open(postgres.New(postgres.Config{
		Conn: conn,
	}))

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
