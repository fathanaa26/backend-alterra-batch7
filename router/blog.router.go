package router

import (
	"blog-app/controller"

	"github.com/gin-gonic/gin"
)

func BlogRouter(rg *gin.RouterGroup) {

	blogRoute := rg.Group("/blog")

	blogRoute.GET("/", controller.BlogsHandler)
	blogRoute.GET("/:id", controller.BlogIdHandler)
	blogRoute.POST("/", controller.AddBlogHandler)

}
