package router

import "github.com/gin-gonic/gin"

func RootRoutes(r *gin.Engine) {
	r.Group("/")

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "Hello form onepolicy-api"})
	})
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "pong"})
	})
	r.GET("/version", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"version": "1.0.0"})
	})
	r.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "OK"})
	})
	r.GET("/status", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"status": "OK"})
	})
	r.GET("/ready", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"status": "OK"})
	})
	r.GET("/live", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"status": "OK"})
	})
}
