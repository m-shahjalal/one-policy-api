package router

import "github.com/gin-gonic/gin"

func RootRoutes(r *gin.Engine) {
	r.Group("/")

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "Hello form onepolicy-api"})
	})

	r.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "OK"})
	})

}
