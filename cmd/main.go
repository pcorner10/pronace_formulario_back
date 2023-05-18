package main

import (
	"log"
	"pronaces_back/api"

	"github.com/spf13/viper"
)

func main() {

	log.Println("Starting server...")
	port := viper.GetString("server.port")

	if port == "" {
		log.Fatal("Port is not set")
	}

	api.Start(port)
}
