package ihttp

import (
	"fmt"
	"pronaces_back/pkg/domain"

	"github.com/gin-gonic/gin"
)

type FormHandler struct {
	formService domain.FormService
}

func NewFormHandler(formService domain.FormService) *FormHandler {
	return &FormHandler{
		formService: formService,
	}
}

func (h *FormHandler) CreateForm0(ctx *gin.Context) {
	res := &domain.Form0Request{}

	err := ctx.BindJSON(res)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Perform a type assertion to convert the interface{} returned by CreateForm to a *entity.Table0
	err = h.formService.CreateForm0(*res)
	if err != nil {
		fmt.Println("error", err.Error())
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// return the response: "suvcessfully created"
	ctx.JSON(200, gin.H{"status": "successfully created"})
}

func (h *FormHandler) CretateForm1(ctx *gin.Context) {
	res := &domain.Form1{}

	err := ctx.BindJSON(&res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = h.formService.CreateForm1(*res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"status": "successfully created"})
}

func (h *FormHandler) CretateForm2(ctx *gin.Context) {
	res := &domain.Form2{}

	err := ctx.BindJSON(&res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = h.formService.CreateForm2(*res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"status": "successfully created"})
}

func (h *FormHandler) CretateForm3(ctx *gin.Context) {
	res := &domain.Form3{}

	err := ctx.BindJSON(&res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = h.formService.CreateForm3(*res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"status": "successfully created"})
}

func (h *FormHandler) CretateForm4(ctx *gin.Context) {
	res := &domain.Form4Request{}

	err := ctx.BindJSON(&res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = h.formService.CreateForm4(*res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"status": "successfully created"})
}

func (h *FormHandler) CretateForm5(ctx *gin.Context) {
	res := &domain.Form5Request{}

	err := ctx.BindJSON(&res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = h.formService.CreateForm5(*res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"status": "successfully created"})
}

func (h *FormHandler) CretateForm6(ctx *gin.Context) {
	res := &domain.Form6Request{}

	err := ctx.BindJSON(&res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = h.formService.CreateForm6(*res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"status": "successfully created"})
}

func (h *FormHandler) CretateForm7(ctx *gin.Context) {
	res := &domain.Form7Request{}

	err := ctx.BindJSON(&res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = h.formService.CreateForm7(*res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"status": "successfully created"})
}

func (h *FormHandler) CretateForm8(ctx *gin.Context) {
	res := &domain.Form8Request{}

	err := ctx.BindJSON(&res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = h.formService.CreateForm8(*res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"status": "successfully created"})
}

func (h *FormHandler) CretateForm8_1(ctx *gin.Context) {
	res := &domain.Form8_1Request{}

	err := ctx.BindJSON(&res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = h.formService.CreateForm8_1(*res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"status": "successfully created"})
}

func (h *FormHandler) CretateForm9(ctx *gin.Context) {
	res := &domain.Form9Request{}

	err := ctx.BindJSON(&res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = h.formService.CreateForm9(*res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"status": "successfully created"})
}

func (h *FormHandler) CretateForm9_1(ctx *gin.Context) {
	res := &domain.Form9_1Request{}

	err := ctx.BindJSON(&res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = h.formService.CreateForm9_1(*res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"status": "successfully created"})
}

func (h *FormHandler) CretateForm10(ctx *gin.Context) {
	res := &domain.Form10Request{}

	err := ctx.BindJSON(&res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = h.formService.CreateForm10(*res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"status": "successfully created"})
}

func (h *FormHandler) CretateForm10_1(ctx *gin.Context) {
	res := &domain.Form10_1Request{}

	err := ctx.BindJSON(&res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = h.formService.CreateForm10_1(*res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"status": "successfully created"})
}

func (h *FormHandler) CretateForm11(ctx *gin.Context) {
	res := &domain.Form11Request{}

	err := ctx.BindJSON(&res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = h.formService.CreateForm11(*res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"status": "successfully created"})
}

func (h *FormHandler) CretateForm12(ctx *gin.Context) {
	res := &domain.Form12{}

	err := ctx.BindJSON(&res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = h.formService.CreateForm12(*res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"status": "successfully created"})
}

func (h *FormHandler) CretateForm13(ctx *gin.Context) {
	res := &domain.Form13{}

	err := ctx.BindJSON(&res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = h.formService.CreateForm13(*res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"status": "successfully created"})
}

func (h *FormHandler) CretateForm14(ctx *gin.Context) {
	res := &domain.Form14Request{}

	err := ctx.BindJSON(&res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = h.formService.CreateForm14(*res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"status": "successfully created"})
}
