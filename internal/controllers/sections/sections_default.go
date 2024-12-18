package sections

import models "github.com/maxwelbm/alkemy-g6/internal/models/sections"

type SectionsDefault struct {
	// sv is the service used by the handler
	sv models.SectionService
}

type SectionFullJSON struct {
	ID                 int     `json:"id"`
	SectionNumber      int     `json:"section_number"`
	CurrentTemperature float64 `json:"current_temperature"`
	MinimumTemperature float64 `json:"minimum_temperature"`
	CurrentCapacity    int     `json:"current_capacity"`
	MinimumCapacity    int     `json:"minimum_capacity"`
	MaximumCapacity    int     `json:"maximum_capacity"`
	WarehouseID        int     `json:"warehouse_id"`
	ProductTypeID      int     `json:"product_type_id"`
}

func NewSectionsDefault(sv models.SectionService) *SectionsDefault {
	return &SectionsDefault{
		sv: sv,
	}
}
