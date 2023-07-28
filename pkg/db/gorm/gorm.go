package db

import (
	"fmt"
	"log"
	"os"
	"pronaces_back/pkg/domain"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type gormStore struct {
	db *gorm.DB
}

func NewGormStore() (domain.FormDB, error) {

	local := true

	// try to catch the config.yml file
	viper.SetConfigName("config") // name of config file (without extension)

	viper.AddConfigPath("config/") // optionally look for config in the working directory

	viper.AutomaticEnv() // read in environment variables that match

	viper.SetConfigType("yml") // REQUIRED if the config file does not have the extension in the name

	if err := viper.ReadInConfig(); err != nil {

		local = false
	}

	var (
		dbUrl   string
		migrate string
	)
	switch local {
	case true:
		fmt.Println("Using config file:", viper.ConfigFileUsed())
		dbUrl = viper.GetString("database.DATABASE_URL")
		migrate = viper.GetString("database.MIGRATE")

	case false:
		fmt.Println("Using Heroku config")
		dbUrl = os.Getenv("DATABASE_URL")
		migrate = os.Getenv("MIGRATE")
	}

	dbHandler, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	if err != nil {
		return nil, err
	}

	if migrate == "true" {
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
