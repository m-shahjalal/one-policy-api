package router

import (
	"github.com/gin-gonic/gin"
	"github.com/m-shahjalal/onepolicy-api/internal/controller"
)

func PolicyRoutes(r *gin.Engine) {
	policyCtrl := controller.PolicyController{}
	group := r.Group("/policies")

	group.GET("/", policyCtrl.GetAllPolicies)
	group.GET("/:id", policyCtrl.GetPolicyById)

	group.PUT("/:id/edit", policyCtrl.UpdatePolicy)
	group.DELETE("/:id", policyCtrl.DeletePolicy)

	group.POST("/", policyCtrl.EditPolicyWithPrompt)
}
