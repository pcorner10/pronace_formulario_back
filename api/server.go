package api

import (
	"pronaces_back/internal/infraestructure/utility"

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
		claims, err := utility.ValidateJWT(token)
		if err != nil {
			ctx.JSON(401, gin.H{"error": "invalid token"})
			ctx.Abort()
			return
		}

		// Agregar usuario al contexto de la petición
		ctx.Set("user", claims.Email)

		ctx.Next()
	}
}
