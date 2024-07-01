package main

import (
	"exoplanet-service/routes"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	r := gin.Default()
	routes.InitializeRoutes(r)
	if err := r.Run(":9000"); err != nil {
		logrus.Fatal("router listening and serve error : ", err)
	}
}
