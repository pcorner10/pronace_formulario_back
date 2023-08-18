package ihttp

import (
	"fmt"
	"pronaces_back/pkg/domain"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service domain.FormService
}

func NewHandler(Service domain.FormService) *Handler {
	return &Handler{
		service: Service,
	}
}

func (h *Handler) CreateForm0(ctx *gin.Context) {
	res := &domain.Form0Request{}

	err := ctx.BindJSON(res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Perform a type assertion to convert the interface{} returned by CreateForm to a *entity.Table0
	err = h.service.CreateForm0(*res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// return the response: "suvcessfully created"
	ctx.JSON(200, gin.H{"status": "successfully created"})
}

func (h *Handler) CretateForm123(ctx *gin.Context) {
	res := &domain.Form123Request{}

	err := ctx.BindJSON(&res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(res)
	err = h.service.CreateForm123(*res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"status": "successfully created"})
}

func (h *Handler) CretateForm4(ctx *gin.Context) {
	res := &domain.Form4Request{}

	err := ctx.BindJSON(&res)
	fmt.Println("resssss", res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = h.service.CreateForm4(*res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"status": "successfully created"})
}

func (h *Handler) CretateForm5(ctx *gin.Context) {
	res := &domain.Form5Request{}

	err := ctx.BindJSON(&res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = h.service.CreateForm5(*res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"status": "successfully created"})
}

func (h *Handler) CretateForm6(ctx *gin.Context) {
	res := &domain.Form6Request{}

	err := ctx.BindJSON(&res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = h.service.CreateForm6(*res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"status": "successfully created"})
}

func (h *Handler) CretateForm7(ctx *gin.Context) {
	res := &domain.Form7Request{}

	err := ctx.BindJSON(&res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// print res with all the data
	fmt.Printf("%+v\n", res)

	err = h.service.CreateForm7(*res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"status": "successfully created"})
}

func (h *Handler) CretateForm8(ctx *gin.Context) {
	res := &domain.Form8Request{}

	err := ctx.BindJSON(&res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = h.service.CreateForm8(*res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"status": "successfully created"})
}

func (h *Handler) CretateForm8_1(ctx *gin.Context) {
	res := &domain.Form8_1Request{}

	err := ctx.BindJSON(&res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = h.service.CreateForm8_1(*res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"status": "successfully created"})
}

func (h *Handler) CretateForm9(ctx *gin.Context) {
	res := &domain.Form9Request{}

	err := ctx.BindJSON(&res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	fmt.Printf("%+v\n", res)

	err = h.service.CreateForm9(*res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"status": "successfully created"})
}

func (h *Handler) CretateForm9_1(ctx *gin.Context) {
	res := &domain.Form9_1Request{}

	err := ctx.BindJSON(&res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = h.service.CreateForm9_1(*res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"status": "successfully created"})
}

func (h *Handler) CretateForm10(ctx *gin.Context) {
	res := &domain.Form10Request{}

	err := ctx.BindJSON(&res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = h.service.CreateForm10(*res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"status": "successfully created"})
}

func (h *Handler) CretateForm10_1(ctx *gin.Context) {
	res := &domain.Form10_1Request{}

	err := ctx.BindJSON(&res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = h.service.CreateForm10_1(*res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"status": "successfully created"})
}

func (h *Handler) CretateForm11(ctx *gin.Context) {
	res := &domain.Form11Request{}

	err := ctx.BindJSON(&res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = h.service.CreateForm11(*res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"status": "successfully created"})
}

func (h *Handler) CretateForm1213(ctx *gin.Context) {
	res := &domain.Form1213Request{}

	err := ctx.BindJSON(&res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = h.service.CreateForm1213(*res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"status": "successfully created"})
}

func (h *Handler) CretateForm14(ctx *gin.Context) {
	res := &domain.Form14Request{}

	err := ctx.BindJSON(&res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = h.service.CreateForm14(*res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"status": "successfully created"})
}
