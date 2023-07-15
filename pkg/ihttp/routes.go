package ihttp

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, handler *Handler) {

	router.GET("/", handler.HelloWorld)

	public := router.Group("/public")
	{
		public.POST("/login", handler.Login)
		public.POST("/register", handler.Register)
	}

	private := router.Group("/api")
	private.Use(JWTMiddleware())

	forms := private.Group("/forms")
	{
		forms.POST("/form0", handler.CreateForm0)
		forms.POST("/form123", handler.CretateForm123)
		forms.POST("/form4", handler.CretateForm4)
		forms.POST("/form5", handler.CretateForm5)
		forms.POST("/form6", handler.CretateForm6)
		forms.POST("/form7", handler.CretateForm7)
		forms.POST("/form8", handler.CretateForm8)
		forms.POST("/form8_1", handler.CretateForm8_1)
		forms.POST("/form9", handler.CretateForm9)
		forms.POST("/form9_1", handler.CretateForm9_1)
		forms.POST("/form10", handler.CretateForm10)
		forms.POST("/form10_1", handler.CretateForm10_1)
		forms.POST("/form11", handler.CretateForm11)
		forms.POST("/form12", handler.CretateForm1213)
		forms.POST("/form14", handler.CretateForm14)

	}

}
