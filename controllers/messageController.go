package controllers

import (
	"jdlgj/auth"
	"jdlgj/models"
	"jdlgj/models/base"
	"jdlgj/repository"
	"net/http"
	"strings"

	"github.com/biezhi/gorm-paginator/pagination"
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
	claim, ok := c.Get("user")
	if ok {
		user, ok := claim.(auth.JWTClaims)
		if ok {
			message.UserID = user.OpenID
			var resource = repository.Create(&message)
			c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": resource})
		}
	}
	c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": "Not able to get current user"})
}

//ListMessages retrieves all messages
func ListMessages(c *gin.Context) {
	var data *pagination.Paginator

	page := c.DefaultQuery("page", "0")
	size := c.DefaultQuery("size", "10")
	orderBy := strings.Split(c.DefaultQuery("orderBy", "created_at desc"), ",")
	answer := c.Query("answered")
	messages := []models.Message{}
	if answer != "" {
		data = repository.FilterBy(&messages, page, size, orderBy, answer, "answered")
	} else {
		data = repository.List(&messages, page, size, orderBy)
	}

	reosources := []models.MessageResource{}
	for _, item := range messages {
		resource, ok := item.ToResource().(models.MessageResource)
		if ok {
			reosources = append(reosources, resource)
		}
	}

	paginationResource := base.PaginationResource{
		TotalElement:   data.TotalRecord,
		DataCollection: reosources,
		CurrentPage:    data.Page,
		TotalPages:     data.TotalPage,
	}
	c.JSON(http.StatusOK, paginationResource)
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

//UpdateMessage updates an existing message
func UpdateMessage(c *gin.Context) {
	id, err := uuid.FromString(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var message models.Message
	resource, err := repository.Update(c, &message, id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Cannot find message!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": resource})
}
