package http

import "github.com/gin-gonic/gin"

func SetupRoutes(router *gin.Engine) {

	private := router.Group("/api")
	{
		private.POST("/login", Login)
	}
}
