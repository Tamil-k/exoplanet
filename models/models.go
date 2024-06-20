package models

type ExoplanetType string

const (
	GasGiant    ExoplanetType = "GasGiant"
	Terrestrial ExoplanetType = "Terrestrial"
)

type Exoplanet struct {
	ID          string        `json:"id,omitempty"`
	Name        string        `json:"name,omitempty"`
	Description string        `json:"description,omitempty"`
	Distance    int           `json:"distance,omitempty"`
	Radius      float64       `json:"radius,omitempty"`
	Mass        float64       `json:"mass,omitempty"`
	Type        ExoplanetType `json:"type,omitempty"`
}

type FuelEstimationRequest struct {
	CrewCapacity int `json:"crewCapacity"`
}

type FuelEstimationResponse struct {
	FuelUnits float64 `json:"fuelUnits"`
}

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	IsError bool        `json:"isError,omitempty"`
}
