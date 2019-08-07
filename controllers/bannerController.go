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

//CreateBanner creates a new banner object
func CreateBanner(c *gin.Context) {
	var banner models.Banner

	if err := c.BindJSON(&banner); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var resource = repository.Create(&banner)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "banner": resource})
}

//ListBanners retrieves all banners
func ListBanners(c *gin.Context) {
	page := c.DefaultQuery("page", "0")
	size := c.DefaultQuery("size", "10")
	orderBy := strings.Split(c.DefaultQuery("orderBy", "createdAt desc"), ",")
	banners := []models.Banner{}
	data := repository.List(&banners, page, size, orderBy)

	bannerResources := []models.BannerResource{}
	for _, item := range banners {
		resource, ok := item.ToResource().(models.BannerResource)
		if ok {
			bannerResources = append(bannerResources, resource)
		}
	}

	paginationResource := base.PaginationResource{
		TotalElement:   data.TotalRecord,
		DataCollection: bannerResources,
		CurrentPage:    data.Page,
		TotalPages:     data.TotalPage,
	}
	c.JSON(http.StatusOK, paginationResource)
}

//GetBanner retrieve banner by id
func GetBanner(c *gin.Context) {
	id, err := uuid.FromString(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var banner models.Banner

	resource, err := repository.FindByID(&banner, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Cannot find banner!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "banner": resource})
}

//UpdateBanner updates an existing banner
func UpdateBanner(c *gin.Context) {
	id, err := uuid.FromString(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var banner models.Banner
	resource, err := repository.Update(c, &banner, id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Cannot find banner!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "banner": resource})
}

//DeleteBanner deletes an existing banner
func DeleteBanner(c *gin.Context) {
	id, err := uuid.FromString(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	var banner models.Banner

	if err := repository.DeleteByID(&banner, id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Cannot find banner!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Banner deleted successfully!"})
}
