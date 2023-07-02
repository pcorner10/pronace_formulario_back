package ihttp

import (
	"net/http"
	"pronaces_back/pkg/domain"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService domain.AuthService
}

func NewAuthHandler(authService domain.AuthService) AuthHandler {
	return AuthHandler{
		authService: authService,
	}
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
	res, err := h.authService.Login(credentials.Email, credentials.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "authentication failed"})
		return
	}

	// Devolver token de autenticación
	ctx.JSON(http.StatusOK, gin.H{"res": res})
}

func (h *AuthHandler) Register(ctx *gin.Context) {
	// Leer datos del cuerpo de la petición
	var credentials = &domain.User{}

	err := ctx.BindJSON(&credentials)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Registrar usuario
	err = h.authService.Register(credentials)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Devolver token de autenticación
	ctx.JSON(http.StatusOK, gin.H{"message": "user created"})
}
