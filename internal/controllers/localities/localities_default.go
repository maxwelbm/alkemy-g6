package localities_controller

import "github.com/maxwelbm/alkemy-g6/internal/models"

type LocalityController struct {
	sv models.LocalityService
}

func NewLocalityController(sv models.LocalityService) *LocalityController {
	return &LocalityController{sv: sv}
}

type LocalityResJSON struct {
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}
