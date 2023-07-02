package ihttp

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, authHandler *AuthHandler, formHandler *FormHandler) {

	public := router.Group("/api")
	{
		public.POST("/login", authHandler.Login)
		public.POST("/register", authHandler.Register)
	}

	private := router.Group("/api")
	private.Use(JWTMiddleware())

	forms := private.Group("/forms")
	{
		forms.POST("/form0", formHandler.CreateForm0)
		forms.POST("/form1", formHandler.CretateForm1)
		forms.POST("/form2", formHandler.CretateForm2)
		forms.POST("/form3", formHandler.CretateForm3)
		forms.POST("/form4", formHandler.CretateForm4)
		forms.POST("/form5", formHandler.CretateForm5)
		forms.POST("/form6", formHandler.CretateForm6)
		forms.POST("/form7", formHandler.CretateForm7)
		forms.POST("/form8", formHandler.CretateForm8)
		forms.POST("/form8_1", formHandler.CretateForm8_1)
		forms.POST("/form9", formHandler.CretateForm9)
		forms.POST("/form9_1", formHandler.CretateForm9_1)
		forms.POST("/form10", formHandler.CretateForm10)
		forms.POST("/form10_1", formHandler.CretateForm10_1)
		forms.POST("/form11", formHandler.CretateForm11)
		forms.POST("/form12", formHandler.CretateForm12)
		forms.POST("/form13", formHandler.CretateForm13)
		forms.POST("/form14", formHandler.CretateForm14)

	}

}
