package main

import (
	"exoplanet-service/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routes.InitializeRoutes(r)
	r.Run(":9000")
}
