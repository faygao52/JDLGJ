package controllers

import (
	"jdlgj/models"
	"jdlgj/models/base"
	"jdlgj/repository"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
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

	var resource = repository.Create(&caseStudy)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "case": resource})
}

//ListCases retrieves all cases
func ListCases(c *gin.Context) {
	page := c.DefaultQuery("page", "0")
	size := c.DefaultQuery("size", "10")
	orderBy := strings.Split(c.DefaultQuery("orderBy", "id"), ",")
	cases := []models.Case{}
	data := repository.List(&cases, page, size, orderBy)

	caseResources := []models.CaseResource{}
	for _, item := range cases {
		resource, ok := item.ToResource().(models.CaseResource)
		if ok {
			caseResources = append(caseResources, resource)
		}
	}

	paginationResource := base.PaginationResource{
		TotalElement:   data.TotalRecords,
		DataCollection: caseResources,
		CurrentPage:    data.CurrentPage,
		TotalPages:     data.TotalPages,
	}
	c.JSON(http.StatusOK, paginationResource)
}

//GetCase retrieve case by id
func GetCase(c *gin.Context) {
	id, err := uuid.FromString(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	var caseModel models.Case
	resource, err := repository.FindByID(&caseModel, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Cannot find case!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "case": resource})
}

//UpdateCase updates an existing case
func UpdateCase(c *gin.Context) {
	id, err := uuid.FromString(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var caseStudy models.Case
	resource, err := repository.Update(c, &caseStudy, id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Cannot find case!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "case": resource})
}

//DeleteCase deletes an existing case
func DeleteCase(c *gin.Context) {
	id, err := uuid.FromString(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var caseStudy models.Case

	if err := repository.DeleteByID(&caseStudy, id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Cannot find case!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Case deleted successfully!"})
}
