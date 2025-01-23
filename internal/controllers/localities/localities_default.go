package localitiesctl

import "github.com/maxwelbm/alkemy-g6/internal/models"

type LocalitiesController struct {
	sv models.LocalityService
}

func NewLocalitiesController(sv models.LocalityService) *LocalitiesController {
	return &LocalitiesController{sv: sv}
}

type LocalityResJSON struct {
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}

type FullLocalitySON struct {
	ID           int    `json:"id"`
	LocalityName string `json:"locality_name"`
	ProvinceName string `json:"province_name"`
	CountryName  string `json:"country_name"`
}
