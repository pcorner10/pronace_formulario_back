package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Start(port string, dbHandler *gorm.DB) {
	fmt.Println("Starting server at port " + port)

	r := gin.Default()

	SetupRoutes(r, dbHandler)

	r.Run(":" + port)
}
