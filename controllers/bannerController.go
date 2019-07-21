package controllers

import (
	"jdlgj/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

//CreateBanner creates a new banner object
func CreateBanner(c *gin.Context) {
	var banner models.Banner

	if err := c.BindJSON(&banner); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	db.Create(&banner)
	c.JSON(http.StatusCreated, &banner)
}

//GetBanners retrieves all banners
func GetBanners(c *gin.Context) {
	var banners []models.Banner
	var bannerResources []models.BannerResource

	db.Find(&banners)
	if len(banners) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No banners found!"})
		return
	}

	for _, item := range banners {
		bannerResources = append(bannerResources, models.BannerResource{ID: item.ID, Title: item.Title, ImageURI: item.ImageURI, Link: item.Link, Visible: item.Visible})
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": bannerResources})
}

//UpdateBanner updates an existing banner
func UpdateBanner(c *gin.Context) {
	id := c.Param("id")
	var banner models.Banner

	if err := db.Where("id = ?", id).First(&banner).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Cannot find banner!"})
		return
	}

	c.BindJSON(&banner)
	db.Save(&banner)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "banner updated successfully!"})
}

//DeleteBanner deletes an existing banner
func DeleteBanner(c *gin.Context) {
	id := c.Param("id")
	var banner models.Banner

	if err := db.Where("id = ?", id).First(&banner).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Cannot find banner!"})
		return
	}

	db.Delete(&banner)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Banner deleted successfully!"})
}
