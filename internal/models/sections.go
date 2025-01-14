package models

import "errors"

var (
	ErrSectionNotFound = errors.New("Section not found")
)

type Section struct {
	ID                 int
	SectionNumber      string
	CurrentTemperature float64
	MinimumTemperature float64
	CurrentCapacity    int
	MinimumCapacity    int
	MaximumCapacity    int
	WarehouseID        int
	ProductTypeID      int
}

type SectionDTO struct {
	SectionNumber      *string  `json:"section_number,omitempty"`
	CurrentTemperature *float64 `json:"current_temperature,omitempty"`
	MinimumTemperature *float64 `json:"minimum_temperature,omitempty"`
	CurrentCapacity    *int     `json:"current_capacity,omitempty"`
	MinimumCapacity    *int     `json:"minimum_capacity,omitempty"`
	MaximumCapacity    *int     `json:"maximum_capacity,omitempty"`
	WarehouseID        *int     `json:"warehouse_id,omitempty"`
	ProductTypeID      *int     `json:"product_type_id,omitempty"`
}

type ProductReport struct {
	SectionID     int
	SectionNumber string
	ProductsCount int
}

type SectionService interface {
	GetAll() (sections []Section, err error)
	GetByID(id int) (section Section, err error)
	GetReportProducts(sectionId int) (reportProducts []ProductReport, err error)
	Create(sec SectionDTO) (newSection Section, err error)
	Update(id int, sec SectionDTO) (updateSection Section, err error)
	Delete(id int) (err error)
}

type SectionRepository interface {
	GetAll() (sections []Section, err error)
	GetByID(id int) (section Section, err error)
	GetReportProducts(sectionId int) (reportProducts []ProductReport, err error)
	Create(sec SectionDTO) (newSection Section, err error)
	Update(id int, sec SectionDTO) (updateSection Section, err error)
	Delete(id int) (err error)
}
