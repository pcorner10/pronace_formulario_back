package api

import (
	"pronaces_back/config"
	"pronaces_back/internal/infraestructure/database"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Start(port string) {
	config.Init()
	dbHandler := database.Init()

	if viper.GetBool("database.migrate") {
		database.Migrate(dbHandler)
	}

	r := gin.Default()
	SetupRoutes(r, dbHandler)

	r.Run(":" + port)
}
