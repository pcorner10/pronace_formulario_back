package handler

import (
	"net/http"
	usecase "pronaces_back/internal/app/usecase"
	"pronaces_back/internal/domain/entity"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	AuthUseCase *usecase.AuthUseCase
}

func NewAuthHandler(authUseCase *usecase.AuthUseCase) *AuthHandler {
	return &AuthHandler{AuthUseCase: authUseCase}
}

func (h *AuthHandler) Login(ctx *gin.Context) {
	// Leer datos del cuerpo de la petición
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := ctx.BindJSON(&credentials)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	// Autenticar usuario
	res, err := h.AuthUseCase.Login(credentials.Email, credentials.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "authentication failed"})
		return
	}

	// Devolver token de autenticación
	ctx.JSON(http.StatusOK, gin.H{"res": res})
}

func (h *AuthHandler) Register(ctx *gin.Context) {
	// Leer datos del cuerpo de la petición
	var credentials = &entity.User{}

	err := ctx.BindJSON(&credentials)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Registrar usuario
	err = h.AuthUseCase.Register(credentials)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Devolver token de autenticación
	ctx.JSON(http.StatusOK, gin.H{"message": "user created"})
}
