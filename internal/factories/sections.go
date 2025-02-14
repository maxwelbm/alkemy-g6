package factories

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/pkg/randstr"
)

type SectionFactory struct {
	db *sql.DB
}

func NewSectionFactory(db *sql.DB) *SectionFactory {
	return &SectionFactory{db: db}
}

func defaultSection() models.Section {
	return models.Section{
		SectionNumber:      randstr.Alphanumeric(8),
		CurrentTemperature: float64(10),
		MinimumTemperature: float64(1),
		CurrentCapacity:    10,
		MinimumCapacity:    1,
		MaximumCapacity:    100,
		WarehouseID:        1,
		ProductTypeID:      1,
	}
}

func (f SectionFactory) Build(section models.Section) models.Section {
	populateSectionParams(&section)

	return section
}

func (f *SectionFactory) Create(section models.Section) (record models.Section, err error) {
	populateSectionParams(&section)

	if err = f.checkWarehouseExists(section.WarehouseID); err != nil {
		return section, err
	}

	query := `
		INSERT INTO sections 
			(
			%s
			section_number,
			current_temperature,
			minimum_temperature,
			current_capacity,
			minimum_capacity,
			maximum_capacity,
			warehouse_id,
			product_type_id
			) 
		VALUES (%s?, ?, ?, ?, ?, ?, ?, ?)
	`

	switch section.ID {
	case 0:
		query = fmt.Sprintf(query, "", "")
	default:
		query = fmt.Sprintf(query, "id,", strconv.Itoa(section.ID)+",")
	}

	_, err = f.db.Exec(query,
		section.SectionNumber,
		section.CurrentTemperature,
		section.MinimumTemperature,
		section.CurrentCapacity,
		section.MinimumCapacity,
		section.MaximumCapacity,
		section.WarehouseID,
		section.ProductTypeID,
	)

	return section, err
}

func populateSectionParams(section *models.Section) {
	defaultSection := defaultSection()
	if section == nil {
		section = &defaultSection
	}

	if section.SectionNumber == "" {
		section.SectionNumber = defaultSection.SectionNumber
	}

	if section.CurrentTemperature == 0 {
		section.CurrentTemperature = defaultSection.CurrentTemperature
	}

	if section.MinimumTemperature == 0 {
		section.MinimumTemperature = defaultSection.MinimumTemperature
	}

	if section.CurrentCapacity == 0 {
		section.CurrentCapacity = defaultSection.CurrentCapacity
	}

	if section.MinimumCapacity == 0 {
		section.MinimumCapacity = defaultSection.MinimumCapacity
	}

	if section.MaximumCapacity == 0 {
		section.MaximumCapacity = defaultSection.MaximumCapacity
	}

	if section.WarehouseID == 0 {
		section.WarehouseID = defaultSection.WarehouseID
	}

	if section.ProductTypeID == 0 {
		section.ProductTypeID = defaultSection.ProductTypeID
	}
}

func (f *SectionFactory) checkWarehouseExists(warehouseID int) (err error) {
	var count int
	err = f.db.QueryRow(`SELECT COUNT(*) FROM warehouses WHERE id = ?`, warehouseID).Scan(&count)

	if err != nil {
		return
	}

	if count > 0 {
		return
	}

	err = f.createWarehouse()

	return
}

func (f *SectionFactory) createWarehouse() (err error) {
	warehouseFactory := NewWarehouseFactory(f.db)
	_, err = warehouseFactory.Create(models.Warehouse{})

	return
}
