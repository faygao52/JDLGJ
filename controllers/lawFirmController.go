package controllers

import (
	"jdlgj/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

//CreateLawFirm creates a new LawFirm object
func CreateLawFirm(c *gin.Context) {
	var lawFirm models.LawFirm

	if err := c.BindJSON(&lawFirm); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	models.GetDB().Create(&lawFirm)
	c.JSON(http.StatusCreated, &lawFirm)
}

//GetLawFirms retrieves all LawFirms
func GetLawFirms(c *gin.Context) {
	lawFirms := []models.LawFirm{}
	lawFirmResources := []models.LawFirmResource{}

	models.GetDB().Find(&lawFirms)
	if len(lawFirms) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No law firms found!"})
		return
	}

	for _, item := range lawFirms {
		lawFirmResources = append(lawFirmResources, models.LawFirmResource{ID: item.ID, Title: item.Title, WorkingHour: item.WorkingHour, Address: item.Address, Description: item.Description, Contact: item.Contact, Review: item.Review, Services: item.Services})
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "lawFirms": lawFirmResources})
}

//GetLawFirm retrieve law firm by its id
func GetLawFirm(c *gin.Context) {
	id := c.Param("id")
	var lawFirm models.LawFirm

	if err := models.GetDB().Where("id = ?", id).First(&lawFirm).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Cannot find law firm!"})
		return
	}

	lawFirmResource := models.LawFirmResource{ID: lawFirm.ID, Title: lawFirm.Title, WorkingHour: lawFirm.WorkingHour, Address: lawFirm.Address, Description: lawFirm.Description, Contact: lawFirm.Contact, Review: lawFirm.Review, Services: lawFirm.Services}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": lawFirmResource})
}

//UpdateLawFirm updates an existing LawFirm
func UpdateLawFirm(c *gin.Context) {
	id := c.Param("id")
	var lawFirm models.LawFirm

	if err := models.GetDB().Where("id = ?", id).First(&lawFirm).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Cannot find law firm!"})
		return
	}

	c.BindJSON(&lawFirm)
	models.GetDB().Save(&lawFirm)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Law firm updated successfully!"})
}

//DeleteLawFirm deletes an existing LawFirm
func DeleteLawFirm(c *gin.Context) {
	id := c.Param("id")
	var lawFirm models.LawFirm

	if err := models.GetDB().Where("id = ?", id).First(&lawFirm).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Cannot find law firm!"})
		return
	}

	models.GetDB().Delete(&lawFirm)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Law Firm deleted successfully!"})
}
