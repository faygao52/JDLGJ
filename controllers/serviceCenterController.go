package controllers

import (
	"strings"
	"jdlgj/models"
	"jdlgj/models/base"
	"jdlgj/repository"
	"net/http"
	uuid "github.com/satori/go.uuid"

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

	var resource = repository.Create(&serviceCenter)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "data":resource})
}

//ListServiceCenters retrieves all ServiceCenters
func ListServiceCenters(c *gin.Context) {
	page := c.DefaultQuery("page", "0")
	size := c.DefaultQuery("size", "10")
	orderBy := strings.Split(c.DefaultQuery("orderBy", "id"), ",")
	serviceCenters := []models.ServiceCenter{}
	data := repository.List(&serviceCenters, page, size, orderBy)

	serviceCenterResources := []models.ServiceCenterResource{}
	for _, item := range serviceCenters {
		resource,ok := item.ToResource().(models.ServiceCenterResource)
		if ok {
			serviceCenterResources = append(serviceCenterResources, resource)
		}
	}
	
	paginationResource := base.PaginationResource {
		TotalElement: data.TotalRecords,
		DataCollection: serviceCenterResources,
		CurrentPage: data.CurrentPage,
		TotalPages: data.TotalPages,
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": paginationResource})
}

//GetServiceCenter retrieves a service center by id
func GetServiceCenter(c *gin.Context) {
	id, err :=  uuid.FromString(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	var serviceCenter models.ServiceCenter

	resource, err := repository.FindByID(&serviceCenter, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Cannot find service center!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": resource})
}

//UpdateServiceCenter updates an existing ServiceCenter
func UpdateServiceCenter(c *gin.Context) {
	id, err :=  uuid.FromString(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var serviceCenter models.ServiceCenter
	resource, err := repository.Update(c, &serviceCenter, id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Cannot find service center!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": resource})
}

//DeleteServiceCenter deletes an existing ServiceCenter
func DeleteServiceCenter(c *gin.Context) {
	id, err :=  uuid.FromString(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	var serviceCenter models.ServiceCenter

	if err := repository.DeleteByID(&serviceCenter, id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Cannot find service center!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Serivce center deleted successfully!"})
}