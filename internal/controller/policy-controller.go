package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/m-shahjalal/onepolicy-api/utils"
)

type PolicyController struct{}

func (ctrl *PolicyController) GetCookiePolicy(c *gin.Context) {
	println("call from cookie")
}
func (ctrl *PolicyController) CreateCookiePolicy(c *gin.Context) {
	data, err := c.GetRawData()
	if err != nil {
		c.JSON(400, gin.H{"error": "Failed to read request body"})
	}

	println("inputs", string(data))
	result, err := utils.GeneratePolicy(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate policy"})
	}
	println("results", result)
	c.JSON(http.StatusCreated, gin.H{"policy": result})
}
