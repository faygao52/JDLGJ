package main

import (
	"jdlgj/auth"
	"jdlgj/controllers"
	"jdlgj/core"

	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("Starting server..")

	router := gin.Default()

	router.GET("/health", controllers.HealthGET)
	authRoutes := router.Group("/auth")
	{
		authRoutes.GET("/wcLogin", controllers.LoginByWechat)
		// auth.POST("/login", controllers.Login)
	}

	v1 := router.Group("/api/v1")
	v1.Use(auth.Required())
	{
		lawFirm := v1.Group("/law-firms")
		{
			lawFirm.GET("", controllers.ListLawFirms)
			lawFirm.GET("/:id", controllers.GetLawFirm)
			lawFirm.POST("/", controllers.CreateLawFirm)
			lawFirm.PUT("/:id", controllers.UpdateLawFirm)
			lawFirm.DELETE("/:id", controllers.DeleteLawFirm)
		}

		serviceCenter := v1.Group("/service-centers")
		{
			serviceCenter.GET("", controllers.ListServiceCenters)
			serviceCenter.GET("/:id", controllers.GetServiceCenter)
			serviceCenter.POST("/", controllers.CreateServiceCenter)
			serviceCenter.PUT("/:id", controllers.UpdateServiceCenter)
			serviceCenter.DELETE("/:id", controllers.DeleteServiceCenter)
		}

		solvedCase := v1.Group("/cases")
		{
			solvedCase.GET("", controllers.ListCases)
			solvedCase.GET("/:id", controllers.GetCase)
			solvedCase.POST("/", controllers.CreateCase)
			solvedCase.PUT("/:id", controllers.UpdateCase)
			solvedCase.DELETE("/:id", controllers.DeleteCase)
		}

		banner := v1.Group("/banners")
		{
			banner.GET("", controllers.ListBanners)
			banner.GET("/:id", controllers.GetBanner)
			banner.POST("/", controllers.CreateBanner)
			banner.PUT("/:id", controllers.UpdateBanner)
			banner.DELETE("/:id", controllers.DeleteBanner)
		}

		message := v1.Group("/messages")
		{
			message.GET("", controllers.ListMessages)
			message.POST("/", controllers.CreateMessage)
			message.DELETE("/:id", controllers.DeleteMessage)
		}
	}

	if core.GetEnv("ENV", "development") == "production" {
		certsFolder := "./nginx/certs/api.jdlvguanjia.com/"
		router.RunTLS(":"+core.GetEnv("PORT", "8080"), certsFolder+"2544652_api.jdlvguanjia.com.pem", certsFolder+"2544652_api.jdlvguanjia.com.key")
	} else {
		router.Run(":" + core.GetEnv("PORT", "8080"))
	}

}
