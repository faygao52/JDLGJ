package main

import (
	"jdlgj/controllers"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("Starting server..")

	router := gin.Default()

	router.GET("/health", controllers.HealthGET)

	v1 := router.Group("/api/v1")
	{
		lawFirm := v1.Group("/law-firms")
		{
			lawFirm.GET("/", controllers.GetLawFirms)
			lawFirm.GET("/:id", controllers.GetLawFirm)
			lawFirm.POST("/", controllers.CreateLawFirm)
			lawFirm.PUT("/:id", controllers.UpdateLawFirm)
			lawFirm.DELETE("/:id", controllers.DeleteLawFirm)
		}

		serviceCenter := v1.Group("/service-centers")
		{
			serviceCenter.GET("/", controllers.GetServiceCenters)
			serviceCenter.GET("/:id", controllers.GetServiceCenter)
			serviceCenter.POST("/", controllers.CreateServiceCenter)
			serviceCenter.PUT("/:id", controllers.UpdateServiceCenter)
			serviceCenter.DELETE("/:id", controllers.DeleteServiceCenter)
		}

		solvedCase := v1.Group("/cases")
		{
			solvedCase.GET("/", controllers.GetCases)
			solvedCase.POST("/", controllers.CreateCase)
			solvedCase.PUT("/:id", controllers.UpdateCase)
			solvedCase.DELETE("/:id", controllers.DeleteCase)
		}

		banner := v1.Group("/banners")
		{
			banner.GET("/", controllers.GetBanners)
			banner.POST("/", controllers.CreateBanner)
			banner.PUT("/:id", controllers.UpdateBanner)
			banner.DELETE("/:id", controllers.DeleteBanner)
		}
	}

	certsFolder := "./nginx/certs/api.jdlvguanjia.com/"
	router.RunTLS(":"+getEnv("PORT", "jingyi"), certsFolder+"2544652_api.jdlvguanjia.com.pem", certsFolder+"2544652_api.jdlvguanjia.com.key")
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
