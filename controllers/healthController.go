package controllers

import "github.com/gin-gonic/gin"

//HealthGET add healthy check for app
func HealthGET(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "UP",
	})
}
