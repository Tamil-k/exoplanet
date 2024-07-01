package services

import (
	"exoplanet-service/models"
	"exoplanet-service/services/gas_giant"
	"exoplanet-service/services/terrestrial"
)

type GravityCalculator interface {
	CalculateGravity() float64
}

func CalculateGravity(exoplanet *models.Exoplanet) float64 {

	var gravityCalculator GravityCalculator

	switch exoplanet.Type {

	case models.GasGiant:
		gravityCalculator = &gas_giant.GasGiant{Radius: exoplanet.Radius}

	case models.Terrestrial:
		gravityCalculator = &terrestrial.Terrestrial{Radius: exoplanet.Radius, Mass: exoplanet.Mass}
	default:
		return 0

	}
	return gravityCalculator.CalculateGravity()
}
