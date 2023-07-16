package main

import (
	"fmt"
	"log"
	"os"
	"pronaces_back/config"
	"pronaces_back/pkg/app"
	db "pronaces_back/pkg/db/gorm"
	"pronaces_back/pkg/domain"
	"pronaces_back/pkg/ihttp"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		log.Print("Port is not set")
		port = "8080"
	}
	log.Printf("listening on port %s", port)
	Start(port)
}

func Start(port string) {
	config.Init()
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

	var dbHandler domain.FormDB
	var err error

	if viper.GetBool("server.localenv") {
		gin.SetMode(gin.DebugMode)
		dbHandler, err = db.NewGormStore()
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	if err != nil {
		panic(err)
	}

	service := app.NewSurveyService(dbHandler)
	handler := ihttp.NewHandler(service)

	ihttp.SetupRoutes(r, handler)

	r.Run(":" + port)
}