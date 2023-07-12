package ihttp

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// Leer token de la cabecera de la petición
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.JSON(400, gin.H{"error": "missing Authorization header"})
			ctx.Abort()
			return
		}

		// Verificar el formato del token
		authHeaderParts := strings.Split(authHeader, " ")
		if len(authHeaderParts) != 2 || authHeaderParts[0] != "Bearer" {
			ctx.JSON(400, gin.H{"error": "invalid Authorization header format"})
			ctx.Abort()
			return
		}

		token := authHeaderParts[1]

		// Validar token
		claims, err := ValidateJWT(token)
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
