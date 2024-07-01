package services

import (
	"exoplanet-service/models"
	"exoplanet-service/storage"
	"github.com/google/uuid"
	"math"
)

type ExoplanetInterface interface {
	AddExoplanet(exoplanet models.Exoplanet) (string, error)
	ListExoplanets() ([]models.Exoplanet, error)
	GetExoplanetByID(id string) (*models.Exoplanet, error)
	UpdateExoplanet(id string, exoplanet models.Exoplanet) error
	DeleteExoplanet(id string) error
	EstimateFuel(id string, crewCapacity int) (*models.FuelEstimationResponse, error)
}

type Exoplanet struct {
}

func NewExoPlanet() ExoplanetInterface {
	return &Exoplanet{}
}

func (e *Exoplanet) AddExoplanet(exoplanet models.Exoplanet) (string, error) {
	exoplanet.ID = uuid.NewString()
	storage.Store.AddExoplanet(exoplanet)
	return exoplanet.ID, nil
}

func (e *Exoplanet) ListExoplanets() ([]models.Exoplanet, error) {
	exoplanets := storage.Store.ListExoplanets()
	return exoplanets, nil
}

func (e *Exoplanet) GetExoplanetByID(id string) (*models.Exoplanet, error) {

	exoplanet, err := storage.Store.GetExoplanetByID(id)
	if err != nil {

		return nil, err
	}
	return exoplanet, nil
}

func (e *Exoplanet) UpdateExoplanet(id string, exoplanet models.Exoplanet) error {

	if err := storage.Store.UpdateExoplanet(id, exoplanet); err != nil {
		return err
	}
	return nil
}

func (e *Exoplanet) DeleteExoplanet(id string) error {

	if err := storage.Store.DeleteExoplanet(id); err != nil {

		return err
	}
	return nil
}

func (e *Exoplanet) EstimateFuel(id string, crewCapacity int) (*models.FuelEstimationResponse, error) {

	exoplanet, err := storage.Store.GetExoplanetByID(id)
	if err != nil {
		return nil, err
	}

	fuelUnits := CalculateFuel(crewCapacity, exoplanet)
	return &models.FuelEstimationResponse{FuelUnits: math.Round(fuelUnits*100) / 100}, nil

}

func CalculateFuel(crewCapacity int, exoplanet *models.Exoplanet) float64 {
	g := CalculateGravity(exoplanet)
	return float64(exoplanet.Distance) / (g * g) * float64(crewCapacity)
}
