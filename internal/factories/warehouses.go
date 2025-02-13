package factories

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/pkg/randstr"
)

type WarehouseFactory struct {
	db *sql.DB
}

func NewWarehouseFactory(db *sql.DB) *WarehouseFactory {
	return &WarehouseFactory{db: db}
}

func defaultWarehouse() models.Warehouse {
	return models.Warehouse{
		Address:            randstr.Alphanumeric(16),
		Telephone:          randstr.Alphanumeric(9),
		WarehouseCode:      randstr.Alphanumeric(8),
		MinimumCapacity:    10,
		MinimumTemperature: 10,
	}
}

func (f WarehouseFactory) Build(warehouse models.Warehouse) models.Warehouse {
	populateWarehouseParams(&warehouse)

	return warehouse
}

func (f *WarehouseFactory) Create(warehouse models.Warehouse) (record models.Warehouse, err error) {
	populateWarehouseParams(&warehouse)

	query := `
		INSERT INTO warehouses 
			(
			%s
			address,
			telephone,
			warehouse_code,
			minimum_capacity,
			minimum_temperature
			) 
		VALUES (%s?, ?, ?, ?, ?, ?)
	`

	switch warehouse.ID {
	case 0:
		query = fmt.Sprintf(query, "", "")
	default:
		query = fmt.Sprintf(query, "id,", strconv.Itoa(warehouse.ID)+",")
	}

	result, err := f.db.Exec(query,
		warehouse.Address,
		warehouse.Telephone,
		warehouse.WarehouseCode,
		warehouse.MinimumCapacity,
		warehouse.MinimumTemperature,
	)

	if err != nil {
		return record, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return record, err
	}

	warehouse.ID = int(id)

	return warehouse, err
}

func populateWarehouseParams(warehouse *models.Warehouse) {
	defaultWarehouse := defaultWarehouse()
	if warehouse == nil {
		warehouse = &defaultWarehouse
	}

	if warehouse.Address == "" {
		warehouse.Address = defaultWarehouse.Address
	}

	if warehouse.Telephone == "" {
		warehouse.Telephone = defaultWarehouse.Telephone
	}

	if warehouse.WarehouseCode == "" {
		warehouse.WarehouseCode = defaultWarehouse.WarehouseCode
	}

	if warehouse.MinimumCapacity == 0 {
		warehouse.MinimumCapacity = defaultWarehouse.MinimumCapacity
	}

	if warehouse.MinimumTemperature == 0 {
		warehouse.MinimumTemperature = defaultWarehouse.MinimumTemperature
	}
}
