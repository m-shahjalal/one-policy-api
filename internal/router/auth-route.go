package router

import (
	"github.com/gin-gonic/gin"
	"github.com/m-shahjalal/onepolicy-api/internal/controller"
	"github.com/m-shahjalal/onepolicy-api/internal/middleware"
)

func AuthRouter(r *gin.Engine) {
	authCtrl := controller.AuthController{}

	public := r.Group("/auth")
	public.POST("/register", authCtrl.Register)
	public.POST("/login", authCtrl.Login)
	public.POST("/refresh", authCtrl.RefreshToken)

	protected := r.Group("/auth")
	protected.Use(middleware.AuthMiddleware())

	protected.GET("/me", authCtrl.GetMe)
	protected.POST("/logout", authCtrl.Logout)
}
