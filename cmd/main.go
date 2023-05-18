package main

import (
	"pronaces_back/config"
	"pronaces_back/db"

	"github.com/spf13/viper"
)

func main() {
	config.Init()
	dbHandler := db.Init()

	if viper.GetBool("database.migrate") {
		dbHandler = dbHandler.Debug()
	}

}
