package api

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

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

	SetupRoutes(r, dbHandler)

	r.Run(":" + port)
}
