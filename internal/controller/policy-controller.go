package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/m-shahjalal/onepolicy-api/config"
	"github.com/m-shahjalal/onepolicy-api/internal/model"
	"github.com/m-shahjalal/onepolicy-api/utils"
)

type PolicyController struct{}

func (ctrl *PolicyController) GetCookiePolicy(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello from cookie policy",
	})
}

func (ctrl PolicyController) CreateCookiePolicy(c *gin.Context) {
	data, err := c.GetRawData()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read request body",
		})
		return
	}

	if len(data) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Request body is empty",
		})
		return
	}

	markdown, err := utils.GeneratePolicy(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to generate policy",
		})
		return
	}

	policy := &model.Policy{
		Inputs:      string(data),
		Markdown:    markdown,
		Policy_type: model.Cookie,
		Effect_date: time.Now(),
	}

	if err := config.DB.Create(&policy).Error; err != nil {
		println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create policy",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": policy,
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
