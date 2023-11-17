package ihttp

import (
	"pronaces_back/pkg/domain"

	"github.com/gin-gonic/gin"
)

type WeddingHandler struct {
	service domain.WeddingService
}

func NewWeddingHandler(Service domain.WeddingService) *WeddingHandler {
	return &WeddingHandler{
		service: Service,
	}
}

func (h *WeddingHandler) FirstOrCreateUserWedding(ctx *gin.Context) {
	res := &domain.UserWedding{}

	err := ctx.BindJSON(res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Perform a type assertion to convert the interface{} returned by CreateForm to a *entity.Table0
	_, err = h.service.FirstOrCreateUserWedding(*res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// return the response: "suvcessfully created"
	ctx.JSON(200, gin.H{"status": "successfully created"})
}

func (h *WeddingHandler) ConfirmarInvitado(ctx *gin.Context) {
	res := &domain.InvitadosConfirmados{}

	err := ctx.BindJSON(&res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	_, err = h.service.ConfirmarInvitado(*res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"status": "successfully created"})
}
