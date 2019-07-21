package controllers

import (
	"jdlgj/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

//CreateServiceCenter creates a new service center object
func CreateServiceCenter(c *gin.Context) {
	var serviceCenter models.ServiceCenter

	if err := c.BindJSON(&serviceCenter); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	models.GetDB().Create(&serviceCenter)
	c.JSON(http.StatusCreated, &serviceCenter)
}

//GetServiceCenters retrieves all ServiceCenters
func GetServiceCenters(c *gin.Context) {
	serviceCenters := []models.ServiceCenter{}
	serviceCenterResources := []models.ServiceCenterResource{}

	models.GetDB().Find(&serviceCenters)
	if len(serviceCenters) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No service centers found!"})
		return
	}

	for _, item := range serviceCenters {
		serviceCenterResources = append(serviceCenterResources, models.ServiceCenterResource{ID: item.ID, Name: item.Name, WorkingHour: item.WorkingHour, Address: item.Address, Contact: item.Contact})
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "ServiceCenters": serviceCenterResources})
}

//GetServiceCenter retrieves a service center by id
func GetServiceCenter(c *gin.Context) {
	id := c.Param("id")
	var serviceCenter models.ServiceCenter

	if err := models.GetDB().Where("id = ?", id).First(&serviceCenter).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Cannot find service center!"})
		return
	}

	serviceCenterResource := models.ServiceCenterResource{ID: serviceCenter.ID, Name: serviceCenter.Name, WorkingHour: serviceCenter.WorkingHour, Address: serviceCenter.Address, Contact: serviceCenter.Contact}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": serviceCenterResource})
}

//UpdateServiceCenter updates an existing ServiceCenter
func UpdateServiceCenter(c *gin.Context) {
	id := c.Param("id")
	var serviceCenter models.ServiceCenter

	if err := models.GetDB().Where("id = ?", id).First(&serviceCenter).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Cannot find service center!"})
		return
	}

	c.BindJSON(&serviceCenter)
	models.GetDB().Save(&serviceCenter)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Service center updated successfully!"})
}

//DeleteServiceCenter deletes an existing ServiceCenter
func DeleteServiceCenter(c *gin.Context) {
	id := c.Param("id")
	var serviceCenter models.ServiceCenter

	if err := models.GetDB().Where("id = ?", id).First(&serviceCenter).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Cannot find service center!"})
		return
	}

	models.GetDB().Delete(&serviceCenter)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Serivce center deleted successfully!"})
}
