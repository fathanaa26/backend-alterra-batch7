package main

import (
	"blog-app/router"
	"net/http"

	"blog-app/config"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	v1 := server.Group("/v1")
	v1.GET("/", func(r *gin.Context) {
		r.JSON(
			http.StatusOK,
			gin.H{
				"data": "v1",
			})
	})
	router.BlogRouter(v1)
	config.ConnectDb()
	server.Run(":3000")
}
