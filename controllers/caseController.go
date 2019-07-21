package controllers

import (
	"jdlgj/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

//CreateCase creates a new case object
func CreateCase(c *gin.Context) {
	var caseStudy models.Case

	if err := c.BindJSON(&caseStudy); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	models.GetDB().Create(&caseStudy)
	c.JSON(http.StatusCreated, &caseStudy)
}

//GetCases retrieves all cases
func GetCases(c *gin.Context) {
	cases := []models.Case{}
	caseResources := []models.CaseResource{}

	models.GetDB().Find(&cases)
	if len(cases) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No cases found!"})
		return
	}

	for _, item := range cases {
		caseResources = append(caseResources, models.CaseResource{ID: item.ID, Catalog: item.Catalog, Question: item.Question, Answer: item.Answer, Contact: item.Contact, Lawyer: item.Lawyer})
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "cases": caseResources})
}

//UpdateCase updates an existing case
func UpdateCase(c *gin.Context) {
	id := c.Param("id")
	var caseStudy models.Case

	if err := models.GetDB().Where("id = ?", id).First(&caseStudy).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Cannot find case!"})
		return
	}

	c.BindJSON(&caseStudy)
	models.GetDB().Save(&caseStudy)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "case updated successfully!"})
}

//DeleteCase deletes an existing case
func DeleteCase(c *gin.Context) {
	id := c.Param("id")
	var caseStudy models.Case

	if err := models.GetDB().Where("id = ?", id).First(&caseStudy).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Cannot find case!"})
		return
	}

	models.GetDB().Delete(&caseStudy)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "case deleted successfully!"})
}
