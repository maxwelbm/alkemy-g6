package loaders

import (
	"encoding/json"
	"os"

	"github.com/maxwelbm/alkemy-g6/internal/models"
)

func NewSectionJSONFile(path string) *SectionJSONFile {
	return &SectionJSONFile{
		path: path,
	}
}

type SectionJSONFile struct {
	path string
}

type SectionJSON struct {
	ID                 int     `json:"id"`
	SectionNumber      string  `json:"section_number"`
	CurrentTemperature float64 `json:"current_temperature"`
	MinimumTemperature float64 `json:"minimum_temperature"`
	CurrentCapacity    int     `json:"current_capacity"`
	MinimumCapacity    int     `json:"minimum_capacity"`
	MaximumCapacity    int     `json:"maximum_capacity"`
	WarehouseID        int     `json:"warehouse_id"`
	ProductTypeID      int     `json:"product_type_id"`
}

func (l *SectionJSONFile) Load() (sections map[int]models.Section, err error) {
	// open file
	file, err := os.Open(l.path)
	if err != nil {
		return
	}
	defer file.Close()

	// decode file
	var sectionsJSON []SectionJSON
	err = json.NewDecoder(file).Decode(&sectionsJSON)
	if err != nil {
		return
	}

	// serialize sections
	sections = make(map[int]models.Section)
	for _, s := range sectionsJSON {
		sections[s.ID] = models.Section{
			ID:                 s.ID,
			SectionNumber:      s.SectionNumber,
			CurrentTemperature: s.CurrentTemperature,
			MinimumTemperature: s.MinimumTemperature,
			CurrentCapacity:    s.CurrentCapacity,
			MinimumCapacity:    s.MinimumCapacity,
			MaximumCapacity:    s.MaximumCapacity,
			WarehouseID:        s.WarehouseID,
			ProductTypeID:      s.ProductTypeID,
		}
	}

	return
}
