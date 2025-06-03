package router

import (
	"github.com/gin-gonic/gin"
	"github.com/m-shahjalal/onepolicy-api/internal/controller"
)

func PolicyRoutes(r *gin.Engine) {
	policyCtrl := controller.PolicyController{}
	group := r.Group("/policies")

	group.GET("/cookies", policyCtrl.GetCookiePolicy)
	group.POST("/cookies", policyCtrl.CreateCookiePolicy)
	group.GET("cookies/:id", policyCtrl.GetPolicyById)

}
