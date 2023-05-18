package main

import (
	"log"
	"pronaces_back/api"
	"pronaces_back/config"
	"pronaces_back/internal/infraestructure/database"

	"github.com/spf13/viper"
)

var port string

func main() {
	config.Init()
	dbHandler := database.Init()

	if viper.GetBool("database.migrate") {
		database.Migrate(dbHandler)
	}

	port := viper.GetString("server.port")

	if port == "" {
		log.Fatal("Port is not set")
		port = "8080"
	}

	api.Start(port, dbHandler)
}
