package ihttp

import (
	"pronaces_back/pkg/app"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(router *gin.Engine, dbHandler *gorm.DB) {

	authHandler := app.BindAuth(dbHandler)

	public := router.Group("/api")
	{
		public.POST("/login", authHandler.Login)
		public.POST("/register", authHandler.Register)
	}

}
