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

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}

	r.Use(cors.New(config))

	SetupRoutes(r, dbHandler)

	r.Run(":" + port)
}
