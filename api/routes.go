package api

import (
	"pronaces_back/internal/app/handler"
	"pronaces_back/internal/app/usecase"
	"pronaces_back/internal/infraestructure/orm"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(router *gin.Engine, dbHandler *gorm.DB) {

	// Crear un nuevo repositorio de usuarios utilizando la conexión a la base de datos.
	userRepo := orm.NewUserRepository(dbHandler)

	// Crear un nuevo caso de uso de autenticación utilizando el repositorio de usuarios.
	authUseCase := usecase.NewAuthUseCase(userRepo)

	// Crear un nuevo manejador de autenticación utilizando el caso de uso de autenticación.
	authHandler := handler.NewAuthHandler(authUseCase)

	public := router.Group("/public")
	{
		public.POST("/login", authHandler.Login)
	}

	private := router.Group("/api")
	private.Use(JWTMiddleware())
	{
		private.POST("/login", nil)
	}
}
