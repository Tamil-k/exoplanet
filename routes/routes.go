package routes

import (
	"exoplanet-service/controllers"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {
	router.POST("/exoplanets", controllers.AddExoPlanet)
	router.GET("/exoplanets", controllers.ListExoplanets)
	router.GET("/exoplanets/:id", controllers.GetExoplanetByID)
	router.PUT("/exoplanets/:id", controllers.UpdateExoplanet)
	router.DELETE("/exoplanets/:id", controllers.DeleteExoplanet)
	router.GET("/exoplanets/:id/:crewcapacity/fuel", controllers.EstimateFuel)
}
