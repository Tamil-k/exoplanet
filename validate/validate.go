package validate

import (
	"errors"
	"exoplanet-service/models"
	"strings"
)

func ValidateExoplanet(exo models.Exoplanet) error {
	if strings.TrimSpace(exo.Name) == "" {
		return errors.New("planet name cannot be empty")
	}
	if strings.TrimSpace(exo.Description) == "" {
		return errors.New("planet description cannot be empty")
	}
	if exo.Distance <= 10 || exo.Distance >= 1000 {
		return errors.New("distance from earth must be between 10 and 1000 light years")
	}
	if exo.Radius <= 0.1 || exo.Radius >= 10 {
		return errors.New("planet radius must be between 0.1 and 10 times Earth's radius")
	}
	if exo.Type != models.Terrestrial && exo.Type != models.GasGiant {
		return errors.New("exoplanet type must be either Terrestrial or GasGiant")
	}
	if exo.Type == models.Terrestrial && (exo.Mass <= 0.1 || exo.Mass >= 10) {
		return errors.New("terrestrial planet mass must be between 0.1 and 10 Earth masses")
	}
	return nil
}
