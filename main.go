package main

import (
	"jdlgj/controllers"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("Starting server..")

	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		// lawFirm := v1.Group("/law-firm")
		// {
		// 	lawFirm.GET("/", getLawFirms)
		// 	lawFirm.GET("/:id", getLawFirm)
		// 	lawFirm.POST("/", createLawFirm)
		// 	lawFirm.PUT("/:id", updateLawFirm)
		// 	lawFirm.DELETE("/:id", deleteLawFirm)
		// }

		// serviceCenter := v1.Group("/service-center")
		// {
		// 	serviceCenter.GET("/", getServiceCenters)
		// 	serviceCenter.GET("/:id", getServiceCenter)
		// 	serviceCenter.POST("/", createServiceCenter)
		// 	serviceCenter.PUT("/:id", updateServiceCenter)
		// 	serviceCenter.DELETE("/:id", deleteServiceCenter)
		// }

		// solvedCase := v1.Group("/case")
		// {
		// 	solvedCase.GET("/", getBanners)
		// 	solvedCase.POST("/", createCase)
		// 	solvedCase.PUT("/:id", updateCase)
		// 	solvedCase.DELETE("/:id", deleteCase)
		// }

		banner := v1.Group("/banner")
		{
			banner.GET("/", controllers.GetBanners)
			banner.POST("/", controllers.CreateBanner)
			banner.PUT("/:id", controllers.UpdateBanner)
			banner.DELETE("/:id", controllers.DeleteBanner)
		}
	}

	router.Run(":8000")
}
