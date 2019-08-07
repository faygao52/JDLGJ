package controllers

import (
	"jdlgj/models"
	"jdlgj/models/base"
	"jdlgj/repository"
	"net/http"
	"strings"

	"github.com/biezhi/gorm-paginator/pagination"
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
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "serviceCenter": resource})
}

//ListServiceCenters retrieves all ServiceCenters
func ListServiceCenters(c *gin.Context) {
	var data *pagination.Paginator

	page := c.DefaultQuery("page", "0")
	size := c.DefaultQuery("size", "10")
	orderBy := strings.Split(c.DefaultQuery("orderBy", "id"), ",")
	query := c.Query("q")
	serviceCenters := []models.ServiceCenter{}
	if query != "" {
		data = repository.SearchAll(&serviceCenters, page, size, orderBy, query, []string{"name", "address"})
	} else {
		data = repository.List(&serviceCenters, page, size, orderBy)
	}

	serviceCenterResources := []models.ServiceCenterResource{}
	for _, item := range serviceCenters {
		resource, ok := item.ToResource().(models.ServiceCenterResource)
		if ok {
			serviceCenterResources = append(serviceCenterResources, resource)
		}
	}

	paginationResource := base.PaginationResource{
		TotalElement:   data.TotalRecord,
		DataCollection: serviceCenterResources,
		CurrentPage:    data.Page,
		TotalPages:     data.TotalPage,
	}
	c.JSON(http.StatusOK, paginationResource)
}

//GetServiceCenter retrieves a service center by id
func GetServiceCenter(c *gin.Context) {
	id, err := uuid.FromString(c.Param("id"))
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

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "serviceCenter": resource})
}

//UpdateServiceCenter updates an existing ServiceCenter
func UpdateServiceCenter(c *gin.Context) {
	id, err := uuid.FromString(c.Param("id"))
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

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "serviceCenter": resource})
}

//DeleteServiceCenter deletes an existing ServiceCenter
func DeleteServiceCenter(c *gin.Context) {
	id, err := uuid.FromString(c.Param("id"))
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
