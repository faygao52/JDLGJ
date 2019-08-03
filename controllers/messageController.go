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

//CreateMessage creates a new message object
func CreateMessage(c *gin.Context) {
	var message models.Message

	if err := c.BindJSON(&message); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var resource = repository.Create(&message)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "data": resource})
}

//ListMessages retrieves all messages
func ListMessages(c *gin.Context) {
	page := c.DefaultQuery("page", "0")
	size := c.DefaultQuery("size", "10")
	orderBy := strings.Split(c.DefaultQuery("orderBy", "id"), ",")
	messages := []models.Message{}
	data := repository.List(&messages, page, size, orderBy)

	reosources := []models.MessageResource{}
	for _, item := range messages {
		resource, ok := item.ToResource().(models.MessageResource)
		if ok {
			reosources = append(reosources, resource)
		}
	}

	paginationResource := base.PaginationResource{
		TotalElement:   data.TotalRecords,
		DataCollection: reosources,
		CurrentPage:    data.CurrentPage,
		TotalPages:     data.TotalPages,
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": paginationResource})
}

//DeleteMessage deletes an existing message
func DeleteMessage(c *gin.Context) {
	id, err := uuid.FromString(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	var message models.Message

	if err := repository.DeleteByID(&message, id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Cannot find message!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Message deleted successfully!"})
}
