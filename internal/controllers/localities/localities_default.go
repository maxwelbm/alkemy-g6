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
