package controllers

import (
	"exoplanet-service/models"
	"exoplanet-service/services"
	"exoplanet-service/validate"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

func AddExoPlanet(c *gin.Context) {

	var exoplanet models.Exoplanet
	if err := c.ShouldBindJSON(&exoplanet); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Json Error format in request body": err.Error()})
		return
	}

	if err := validate.ValidateExoplanet(exoplanet); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			IsError: true,
			Message: err.Error(),
		})
		return
	}

	id, err := services.NewExoPlanet().AddExoplanet(exoplanet)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, models.Response{
			IsError: true,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, models.Response{
		Message: "Planet added successfully",
		Data:    id,
	})

}
func ListExoplanets(c *gin.Context) {
	expoPlanets, err := services.NewExoPlanet().ListExoplanets()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, models.Response{})
	}
	c.JSON(http.StatusOK, models.Response{
		Message: "All exoplanets fetched successfully",
		Data:    expoPlanets,
	})

}
func GetExoplanetByID(c *gin.Context) {
	id := c.Param("id")

	if strings.TrimSpace(id) == "" {
		c.JSON(http.StatusBadRequest, models.Response{
			IsError: true,
			Message: "Id cannot be blank",
		})
	}
	exoplanet, err := services.NewExoPlanet().GetExoplanetByID(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.Response{
			IsError: true,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, models.Response{
		Message: "Exoplanet fetched successfully",
		Data:    exoplanet,
	})

}

func DeleteExoplanet(c *gin.Context) {
	id := c.Param("id")

	if strings.TrimSpace(id) == "" {
		c.JSON(http.StatusBadRequest, models.Response{
			IsError: true,
			Message: "Id cannot be blank",
		})
	}
	if err := services.NewExoPlanet().DeleteExoplanet(id); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.Response{
			IsError: true,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, models.Response{
		Message: "Exoplanet deleted successfully",
	})

}

func UpdateExoplanet(c *gin.Context) {
	id := c.Param("id")

	if strings.TrimSpace(id) == "" {
		c.JSON(http.StatusBadRequest, models.Response{
			IsError: true,
			Message: "Id cannot be blank",
		})
	}
	var exoplanet models.Exoplanet
	if err := c.ShouldBindJSON(&exoplanet); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Json Error format in request body": err.Error()})
		return
	}
	exoplanet.ID = id
	if err := validate.ValidateExoplanet(exoplanet); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			IsError: true,
			Message: err.Error(),
		})
		return
	}
	if err := services.NewExoPlanet().UpdateExoplanet(id, exoplanet); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.Response{
			IsError: true,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, models.Response{
		Message: "Exoplanet updated successfully",
	})

}

func EstimateFuel(c *gin.Context) {
	id := c.Param("id")

	if strings.TrimSpace(id) == "" {
		c.JSON(http.StatusBadRequest, models.Response{
			IsError: true,
			Message: "Id cannot be blank",
		})
	}

	crewCapacity := c.Param("crewcapacity")
	if strings.TrimSpace(crewCapacity) == "" {
		c.JSON(http.StatusBadRequest, models.Response{
			IsError: true,
			Message: "Id cannot be blank",
		})
	}
	crew, err := strconv.Atoi(crewCapacity)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			IsError: true,
			Message: err.Error(),
		})
		return
	}

	respone, err := services.NewExoPlanet().EstimateFuel(id, crew)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.Response{
			IsError: true,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, models.Response{
		Message: "Exoplanet fuel estimated successfully",
		Data:    respone,
	})

}
