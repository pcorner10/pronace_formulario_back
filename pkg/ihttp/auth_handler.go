package ihttp

import (
	"net/http"
	"pronaces_back/pkg/domain"

	"github.com/gin-gonic/gin"
)

func (h *Handler) HelloWorld(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Hello World!"})
}

func (h *Handler) Login(ctx *gin.Context) {
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
	res, err := h.service.LoginUser(credentials.Email, credentials.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "authentication failed"})
		return
	}

	// Devolver token de autenticación
	ctx.JSON(http.StatusOK, gin.H{"res": res})
}

func (h *Handler) Register(ctx *gin.Context) {
	// Leer datos del cuerpo de la petición
	var credentials = domain.User{}

	err := ctx.BindJSON(&credentials)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Registrar usuario
	err = h.service.RegisterUser(&credentials)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Devolver token de autenticación
	ctx.JSON(http.StatusOK, gin.H{"message": "user created"})
}
