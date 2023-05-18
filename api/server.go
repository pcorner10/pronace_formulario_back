package api

import (
	"pronaces_back/internal/domain/service"

	"github.com/gin-gonic/gin"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// Leer token de la cabecera de la petición
		token := ctx.GetHeader("Authorization")
		if token == "" {
			ctx.JSON(400, gin.H{"error": "missing Authorization header"})
			ctx.Abort()
			return
		}

		// Validar token
		claims, err := service.ValidateJWT(token)
		if err != nil {
			ctx.JSON(401, gin.H{"error": "invalid token"})
			ctx.Abort()
			return
		}

		// Agregar usuario al contexto de la petición
		ctx.Set("user", claims)

		ctx.Next()
	}
}
