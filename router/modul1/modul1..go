package modul1

import (
	handler "simple-crud-go/handler/modul1"
	mdwr "simple-crud-go/middleware"

	"github.com/gin-gonic/gin"
)

func Router(router *gin.Engine) {
	public := router.Group("/")
	private := router.Group("/")
	public.POST("api/data", handler.DataHandler)
	public.POST("api/data/add", handler.CreateHandler)
	public.PUT("api/data/update", handler.UpdateHandler)
	public.DELETE("api/data/delete", handler.DeleteHandler)

	// authentication needs
	private.Use(mdwr.JWTMiddleware())
	{
		private.POST("api/data/private/add", handler.CreateHandler)
		private.PUT("api/data/private/update", handler.UpdateHandler)
		private.DELETE("api/data/private", handler.DeleteHandler)
	}
}
