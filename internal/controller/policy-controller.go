package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/m-shahjalal/onepolicy-api/config"
	"github.com/m-shahjalal/onepolicy-api/internal/model"
	"github.com/m-shahjalal/onepolicy-api/utils"
)

type PolicyController struct{}

func (ctrl *PolicyController) GetAllPolicies(c *gin.Context) {
	userId := utils.GetUserID(c)

	policies := []model.Policy{}
	config.DB.Find(&policies).Where("user_id = ?", userId)
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Policies fetched successfully",
		"data":    policies,
	})
}

func (ctrl *PolicyController) UpdatePolicy(c *gin.Context) {
	id := c.Params.ByName("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Policy ID is required",
		})
		return
	}

	inputs := c.Request.Body
	if inputs == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Request body is empty",
		})
		return
	}

	var policy model.Policy
	if err := c.ShouldBindJSON(&policy); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// TODO: Update policy
	result := config.DB.Model(&model.Policy{}).Where("id = ?", id).Updates(policy)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update policy",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully updated policy",
		"data":    policy,
		"success": true,
	})
}

func (ctrl *PolicyController) GetPolicyById(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Policy ID is required",
		})
		return
	}

	var policy model.Policy
	result := config.DB.First(&policy, "id = ?", id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Policy not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": policy,
	})
}

func (ctrl *PolicyController) DeletePolicy(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Policy ID is required",
		})
		return
	}

	result := config.DB.Delete(&model.Policy{}, "id = ?", id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete policy",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully deleted policy",
		"success": true,
	})
}

func (ctrl *PolicyController) EditPolicyWithPrompt(c *gin.Context) {
	inputs := c.Request.Body
	if inputs == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Request body is empty",
		})
		return
	}

	var policy model.Policy
	if err := c.ShouldBindJSON(&policy); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// TODO: Edit policy with prompt
	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully edited policy",
		"data":    policy,
		"success": true,
	})
}
