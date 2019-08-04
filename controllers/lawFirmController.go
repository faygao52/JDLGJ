package controllers

import (
	"jdlgj/models"
	"jdlgj/models/base"
	"jdlgj/repository"
	"strings"

	"net/http"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
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

	var resource = repository.Create(&lawFirm)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "lawFirm": resource})
}

//ListLawFirms retrieves all LawFirms
func ListLawFirms(c *gin.Context) {
	page := c.DefaultQuery("page", "0")
	size := c.DefaultQuery("size", "10")
	orderBy := strings.Split(c.DefaultQuery("orderBy", "id"), ",")
	lawFirms := []models.LawFirm{}
	data := repository.List(&lawFirms, page, size, orderBy)

	lawFirmResources := []models.LawFirmResource{}
	for _, item := range lawFirms {
		resource, ok := item.ToResource().(models.LawFirmResource)
		if ok {
			lawFirmResources = append(lawFirmResources, resource)
		}
	}

	paginationResource := base.PaginationResource{
		TotalElement:   data.TotalRecords,
		DataCollection: lawFirmResources,
		CurrentPage:    data.CurrentPage,
		TotalPages:     data.TotalPages,
	}
	c.JSON(http.StatusOK, paginationResource)
}

//GetLawFirm retrieve law firm by its id
func GetLawFirm(c *gin.Context) {
	id, err := uuid.FromString(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	var lawFirm models.LawFirm

	resource, err := repository.FindByID(&lawFirm, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Cannot find law firm!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "lawFirm": resource})
}

//UpdateLawFirm updates an existing LawFirm
func UpdateLawFirm(c *gin.Context) {
	id, err := uuid.FromString(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	var lawFirm models.LawFirm
	resource, err := repository.Update(c, &lawFirm, id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Cannot find law firm!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "lawFirm": resource})
}

//DeleteLawFirm deletes an existing LawFirm
func DeleteLawFirm(c *gin.Context) {
	id, err := uuid.FromString(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	var lawFirm models.LawFirm

	if err := repository.DeleteByID(&lawFirm, id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Cannot find law firm!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Law firm deleted successfully!"})
}
