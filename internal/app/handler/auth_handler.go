package handler

import (
	"pronaces_back/internal/domain/repository"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	userRepo repository.UserRepository
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

func (h *AuthHandler) Login(ctx *gin.Context) {

	// Leer datos del cuerpo de la petición
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := ctx.BindJSON(&credentials)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "invalid request"})
		return
	}

}
