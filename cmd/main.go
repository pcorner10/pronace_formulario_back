package main

import (
	"fmt"
	"log"
	"pronaces_back/config"
	"pronaces_back/pkg/app"
	"pronaces_back/pkg/db"
	"pronaces_back/pkg/ihttp"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func main() {
	config.Init()
	dbHandler := db.Init()

	if viper.GetBool("database.migrate") {
		db.Migrate(dbHandler)
	}

	port := viper.GetString("server.port")

	if port == "" {
		log.Fatal("Port is not set")
		port = "8080"
	}

	Start(port, dbHandler)
}

func Start(port string, dbHandler *gorm.DB) {
	fmt.Println("Starting server at port " + port)

	r := gin.Default()

	// Add CORS middleware
	config := cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		AllowCredentials: true,
	}

	r.Use(cors.New(config))

	authHandler := app.BindAuth(dbHandler)


	formHandler := app.BindForm(dbHandler)


	ihttp.SetupRoutes(r, authHandler, formHandler)

	r.Run(":" + port)
}
