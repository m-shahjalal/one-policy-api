package router

import (
	"github.com/gin-gonic/gin"
	"github.com/m-shahjalal/onepolicy-api/internal/controller"
)

func AuthRouter(r *gin.Engine) {
	authCtrl := controller.AuthController{}

	public := r.Group("/auth")
	{
		public.POST("/register", authCtrl.Register)
		public.POST("/login", authCtrl.Login)
		public.POST("/forgot-password", authCtrl.ForgotPassword)
		public.POST("/reset-password", authCtrl.ResetPassword)
		public.POST("/refresh-token", authCtrl.RefreshToken)
	}

	protected := r.Group("/auth")
	{
		protected.POST("/logout", authCtrl.Logout)
		protected.GET("/me", authCtrl.GetMe)
		protected.PUT("/profile", authCtrl.UpdateProfile)
	}
}
