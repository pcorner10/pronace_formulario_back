package handler

import (
	"fmt"
	apprepository "pronaces_back/internal/app/repository"
	usecase "pronaces_back/internal/app/usecase"

	"github.com/gin-gonic/gin"
)

type FormHandler struct {
	FormUseCase *usecase.FormUseCase
}

func NewFormHandler(formUseCase *usecase.FormUseCase) *FormHandler {
	return &FormHandler{
		FormUseCase: formUseCase,
	}
}

func (h *FormHandler) CreateForm0(ctx *gin.Context) {
	res := &apprepository.Form0Request{}

	err := ctx.BindJSON(res)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Perform a type assertion to convert the interface{} returned by CreateForm to a *entity.Table0
	err = h.FormUseCase.CreateForm0(*res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// return the response: "suvcessfully created"
	ctx.JSON(200, gin.H{"status": "successfully created"})
}

func (h *FormHandler) CretateForm1(ctx *gin.Context) {
	res := &apprepository.Form1{}

	err := ctx.BindJSON(&res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = h.FormUseCase.CreateForm1(*res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"status": "successfully created"})
}

func (h *FormHandler) CretateForm2(ctx *gin.Context) {
	res := &apprepository.Form2{}

	err := ctx.BindJSON(&res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = h.FormUseCase.CreateForm2(*res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"status": "successfully created"})
}

func (h *FormHandler) CretateForm3(ctx *gin.Context) {
	res := &apprepository.Form3{}

	err := ctx.BindJSON(&res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = h.FormUseCase.CreateForm3(*res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"status": "successfully created"})
}

func (h *FormHandler) CretateForm4(ctx *gin.Context) {
	res := &apprepository.Form4Request{}

	err := ctx.BindJSON(&res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = h.FormUseCase.CreateForm4(*res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"status": "successfully created"})
}

func (h *FormHandler) CretateForm5(ctx *gin.Context) {
	res := &apprepository.Form5Request{}

	err := ctx.BindJSON(&res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = h.FormUseCase.CreateForm5(*res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"status": "successfully created"})
}

func (h *FormHandler) CretateForm6(ctx *gin.Context) {
	res := &apprepository.Form6Request{}

	err := ctx.BindJSON(&res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = h.FormUseCase.CreateForm6(*res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"status": "successfully created"})
}

func (h *FormHandler) CretateForm7(ctx *gin.Context) {
	res := &apprepository.Form7Request{}

	err := ctx.BindJSON(&res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = h.FormUseCase.CreateForm7(*res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"status": "successfully created"})
}

func (h *FormHandler) CretateForm8(ctx *gin.Context) {
	res := &apprepository.Form8Request{}

	err := ctx.BindJSON(&res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = h.FormUseCase.CreateForm8(*res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"status": "successfully created"})
}

func (h *FormHandler) CretateForm8_1(ctx *gin.Context) {
	res := &apprepository.Form8_1Request{}

	err := ctx.BindJSON(&res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = h.FormUseCase.CreateForm8_1(*res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"status": "successfully created"})
}

func (h *FormHandler) CretateForm9(ctx *gin.Context) {
	res := &apprepository.Form9Request{}

	err := ctx.BindJSON(&res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = h.FormUseCase.CreateForm9(*res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"status": "successfully created"})
}

func (h *FormHandler) CretateForm9_1(ctx *gin.Context) {
	res := &apprepository.Form9_1Request{}

	err := ctx.BindJSON(&res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = h.FormUseCase.CreateForm9_1(*res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"status": "successfully created"})
}

func (h *FormHandler) CretateForm10(ctx *gin.Context) {
	res := &apprepository.Form10Request{}

	err := ctx.BindJSON(&res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = h.FormUseCase.CreateForm10(*res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"status": "successfully created"})
}

func (h *FormHandler) CretateForm10_1(ctx *gin.Context) {
	res := &apprepository.Form10_1Request{}

	err := ctx.BindJSON(&res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = h.FormUseCase.CreateForm10_1(*res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"status": "successfully created"})
}

func (h *FormHandler) CretateForm11(ctx *gin.Context) {
	res := &apprepository.Form11Request{}

	err := ctx.BindJSON(&res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = h.FormUseCase.CreateForm11(*res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"status": "successfully created"})
}

func (h *FormHandler) CretateForm12(ctx *gin.Context) {
	res := &apprepository.Form12{}

	err := ctx.BindJSON(&res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = h.FormUseCase.CreateForm12(*res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"status": "successfully created"})
}

func (h *FormHandler) CretateForm13(ctx *gin.Context) {
	res := &apprepository.Form13{}

	err := ctx.BindJSON(&res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = h.FormUseCase.CreateForm13(*res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"status": "successfully created"})
}

func (h *FormHandler) CretateForm14(ctx *gin.Context) {
	res := &apprepository.Form14Request{}

	err := ctx.BindJSON(&res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = h.FormUseCase.CreateForm14(*res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"status": "successfully created"})
}
