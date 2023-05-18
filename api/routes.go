package api

import (
	"pronaces_back/internal/app/handler"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {

	authHandler := handler.NewAuthHandler()

	private := router.Group("/api")
	private.Use(JWTMiddleware())
	{
		private.POST("/login", authHandler.Login)
	}
}
