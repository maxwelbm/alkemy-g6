package carries_controller

import "github.com/maxwelbm/alkemy-g6/internal/models"

type CarriesDefault struct {
	sv models.CarriesService
}

func NewCarriesDefault(sv models.CarriesService) *CarriesDefault {
	return &CarriesDefault{sv: sv}
}

type CarriesReqJSON struct {
	CID         string `json:"cid"`
	CompanyName string `json:"company_name"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
	LocalityID  int    `json:"locality_id"`
}

type FullCarryJSON struct {
	ID          int    `json:"id"`
	CID         string `json:"cid"`
	CompanyName string `json:"company_name"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
	LocalityID  int    `json:"locality_id"`
}

type CarriesResJSON struct {
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}
