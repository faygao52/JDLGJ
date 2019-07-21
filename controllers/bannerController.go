package controllers

import (
	"jdlgj/models"
	"net/http"

	"github.com/gin-gonic/gin"
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

	models.GetDB().Create(&banner)
	c.JSON(http.StatusCreated, &banner)
}

//GetBanners retrieves all banners
func GetBanners(c *gin.Context) {
	banners := []models.Banner{}
	bannerResources := []models.BannerResource{}

	models.GetDB().Find(&banners)
	if len(banners) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No banners found!"})
		return
	}

	for _, item := range banners {
		bannerResources = append(bannerResources, models.BannerResource{ID: item.ID, Title: item.Title, ImageURI: item.ImageURI, Link: item.Link, Visible: item.Visible})
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "banners": bannerResources})
}

//UpdateBanner updates an existing banner
func UpdateBanner(c *gin.Context) {
	id := c.Param("id")
	var banner models.Banner

	if err := models.GetDB().Where("id = ?", id).First(&banner).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Cannot find banner!"})
		return
	}

	c.BindJSON(&banner)
	models.GetDB().Save(&banner)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "banner updated successfully!"})
}

//DeleteBanner deletes an existing banner
func DeleteBanner(c *gin.Context) {
	id := c.Param("id")
	var banner models.Banner

	if err := models.GetDB().Where("id = ?", id).First(&banner).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Cannot find banner!"})
		return
	}

	models.GetDB().Delete(&banner)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Banner deleted successfully!"})
}
