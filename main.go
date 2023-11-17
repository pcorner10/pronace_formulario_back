package main

import (
	"log"
	"os"
	"pronaces_back/pkg/app"
	db "pronaces_back/pkg/db/gorm"
	"pronaces_back/pkg/domain"
	"pronaces_back/pkg/ihttp"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		log.Print("Port is not set")
		port = "4000"
	}
	log.Printf("listening on port %s", port)
	Start(port)
}

func Start(port string) {
	//config.Init()
	log.Println("Starting server at port " + port)

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
	var weddingDbHandler domain.WeddingDB
	var err error

	dbHandler, weddingDbHandler, err = db.NewGormStore()

	if err != nil {
		panic(err)
	}

	service := app.NewSurveyService(dbHandler)
	handler := ihttp.NewHandler(service)

	weddingService := app.NewWeddingService(weddingDbHandler)
	weddingHandler := ihttp.NewWeddingHandler(weddingService)

	ihttp.SetupRoutes(r, handler, weddingHandler)

	r.Run(":" + port)
}
