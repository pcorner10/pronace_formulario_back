package handler

import (
	"fmt"
	usecase "pronaces_back/internal/app/usecase"
	"pronaces_back/internal/domain/entity"

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

func (h *FormHandler) CreateForm(ctx *gin.Context) {
	res := &entity.Table0{}

	err := ctx.BindJSON(res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// print res showing field names and values
	fmt.Printf("%+v\n", res)

	// Perform a type assertion to convert the interface{} returned by CreateForm to a *entity.Table0
	res, err = h.FormUseCase.CreateForm(*res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"data": res})
}

func (h *FormHandler) GetForm1(ctx *gin.Context) {
	res := &entity.Table1{}

	err := ctx.BindJSON(&res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"data": res})
}

func (h *FormHandler) GetForm2(ctx *gin.Context) {
	res := &entity.Table2{}

	err := ctx.BindJSON(&res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"data": res})
}

func (h *FormHandler) GetForm3(ctx *gin.Context) {
	res := &entity.Table3{}

	err := ctx.BindJSON(&res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"data": res})
}

func (h *FormHandler) GetForm4(ctx *gin.Context) {
	res := &entity.Table4{}

	err := ctx.BindJSON(&res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"data": res})
}

func (h *FormHandler) GetForm5(ctx *gin.Context) {
	res := &entity.Table5{}

	err := ctx.BindJSON(&res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"data": res})
}

func (h *FormHandler) GetForm6(ctx *gin.Context) {
	res := &entity.Table6{}

	err := ctx.BindJSON(&res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"data": res})
}

func (h *FormHandler) GetForm7(ctx *gin.Context) {
	res := &entity.Table7{}

	err := ctx.BindJSON(&res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"data": res})
}

func (h *FormHandler) GetForm8(ctx *gin.Context) {
	res := &entity.Table8{}

	err := ctx.BindJSON(&res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"data": res})
}

func (h *FormHandler) GetForm8_1(ctx *gin.Context) {
	res := &entity.Table8_1{}

	err := ctx.BindJSON(&res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"data": res})
}

func (h *FormHandler) GetForm9(ctx *gin.Context) {
	res := &entity.Table9{}

	err := ctx.BindJSON(&res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"data": res})
}

func (h *FormHandler) GetForm9_1(ctx *gin.Context) {
	res := &entity.Table9_1{}

	err := ctx.BindJSON(&res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"data": res})
}

func (h *FormHandler) GetForm10(ctx *gin.Context) {
	res := &entity.Table10{}

	err := ctx.BindJSON(&res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"data": res})
}

func (h *FormHandler) GetForm10_1(ctx *gin.Context) {
	res := &entity.Table10_1{}

	err := ctx.BindJSON(&res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"data": res})
}

func (h *FormHandler) GetForm11(ctx *gin.Context) {
	res := &entity.Table11{}

	err := ctx.BindJSON(&res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"data": res})
}

func (h *FormHandler) GetForm11_1(ctx *gin.Context) {
	res := &entity.Table11{}

	err := ctx.BindJSON(&res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"data": res})
}

func (h *FormHandler) GetForm12(ctx *gin.Context) {
	res := &entity.Table12{}

	err := ctx.BindJSON(&res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"data": res})
}

func (h *FormHandler) GetForm13(ctx *gin.Context) {
	res := &entity.Table13{}

	err := ctx.BindJSON(&res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"data": res})
}

func (h *FormHandler) GetForm14(ctx *gin.Context) {
	res := &entity.Table14{}

	err := ctx.BindJSON(&res)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"data": res})
}
