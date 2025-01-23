package service_test

import (
	"errors"
	"testing"

	"github.com/go-sql-driver/mysql"
	"github.com/maxwelbm/alkemy-g6/internal/models"
	warehousesrp "github.com/maxwelbm/alkemy-g6/internal/repository/warehouses"
	"github.com/maxwelbm/alkemy-g6/internal/service"
	"github.com/maxwelbm/alkemy-g6/pkg/mysqlerr"
	"github.com/stretchr/testify/require"
)

var warehousesFixture = []models.Warehouse{
	{ID: 1, Address: "123 Main St", Telephone: "555-1234", WarehouseCode: "WH001", MinimumCapacity: 100, MinimumTemperature: -10},
	{ID: 2, Address: "100 Main Av", Telephone: "555-9876", WarehouseCode: "WH002", MinimumCapacity: 50, MinimumTemperature: -5},
	{ID: 3, Address: "215 Main Ro", Telephone: "555-2830", WarehouseCode: "WH003", MinimumCapacity: 80, MinimumTemperature: 0},
}

func TestWarehousesDefault_GetAll(t *testing.T) {
	tests := []struct {
		name            string
		warehouse       []models.Warehouse
		err             error
		wantedWarehouse []models.Warehouse
		wantedErr       error
	}{
		{
			name:            "When the repository returns a warehouse",
			warehouse:       warehousesFixture,
			err:             nil,
			wantedWarehouse: warehousesFixture,
			wantedErr:       nil,
		},
		{
			name:            "When the repository returns an error",
			warehouse:       []models.Warehouse{},
			err:             errors.New("internal error"),
			wantedWarehouse: []models.Warehouse{},
			wantedErr:       errors.New("internal error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			rp := warehousesrp.NewWarehousesRepositoryMock()
			rp.On("GetAll").Return(tt.warehouse, tt.err)
			sv := service.NewWarehousesService(rp)

			// Act
			warehouse, err := sv.GetAll()

			// Assert
			require.Equal(t, tt.wantedWarehouse, warehouse)
			require.Equal(t, tt.wantedErr, err)
		})
	}
}

func TestWarehousesDefault_GetByID(t *testing.T) {
	tests := []struct {
		name            string
		warehouse       models.Warehouse
		err             error
		wantedWarehouse models.Warehouse
		wantedErr       error
	}{
		{
			name:            "When the repository returns a warehouse",
			warehouse:       warehousesFixture[0],
			err:             nil,
			wantedWarehouse: warehousesFixture[0],
			wantedErr:       nil,
		},
		{
			name:            "When the repository returns an error",
			warehouse:       models.Warehouse{},
			err:             models.ErrWareHouseNotFound,
			wantedWarehouse: models.Warehouse{},
			wantedErr:       models.ErrWareHouseNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			rp := warehousesrp.NewWarehousesRepositoryMock()
			rp.On("GetByID", tt.warehouse.ID).Return(tt.warehouse, tt.err)
			sv := service.NewWarehousesService(rp)

			// Act
			warehouse, err := sv.GetByID(tt.warehouse.ID)

			// Assert
			require.Equal(t, tt.wantedWarehouse, warehouse)
			require.Equal(t, tt.wantedErr, err)
		})
	}
}

func TestWarehousesDefault_Create(t *testing.T) {
	tests := []struct {
		name            string
		warehouse       models.Warehouse
		err             error
		wantedWarehouse models.Warehouse
		wantedErr       error
	}{
		{
			name:            "When the repository returns a warehouse",
			warehouse:       warehousesFixture[0],
			err:             nil,
			wantedWarehouse: warehousesFixture[0],
			wantedErr:       nil,
		},
		{
			name:            "When the repository returns an error",
			warehouse:       warehousesFixture[0],
			err:             &mysql.MySQLError{Number: mysqlerr.CodeDuplicateEntry},
			wantedWarehouse: warehousesFixture[0],
			wantedErr:       &mysql.MySQLError{Number: mysqlerr.CodeDuplicateEntry},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			rp := warehousesrp.NewWarehousesRepositoryMock()
			sv := service.NewWarehousesService(rp)
			dto := models.WarehouseDTO{
				Address:            &tt.warehouse.Address,
				Telephone:          &tt.warehouse.Telephone,
				WarehouseCode:      &tt.warehouse.WarehouseCode,
				MinimumCapacity:    &tt.warehouse.MinimumCapacity,
				MinimumTemperature: &tt.warehouse.MinimumTemperature,
			}
			// Act
			rp.On("Create", dto).Return(tt.warehouse, tt.err)
			warehouse, err := sv.Create(dto)

			// Assert
			require.Equal(t, tt.wantedWarehouse, warehouse)
			require.Equal(t, tt.wantedErr, err)
		})
	}
}

func TestWarehousesDefault_Update(t *testing.T) {
	tests := []struct {
		name            string
		warehouse       models.Warehouse
		err             error
		wantedWarehouse models.Warehouse
		wantedErr       error
	}{
		{
			name:            "When the repository returns a warehouse",
			warehouse:       warehousesFixture[0],
			err:             nil,
			wantedWarehouse: warehousesFixture[0],
			wantedErr:       nil,
		},
		{
			name:            "When the repository returns an error",
			warehouse:       warehousesFixture[0],
			err:             models.ErrWareHouseNotFound,
			wantedWarehouse: warehousesFixture[0],
			wantedErr:       models.ErrWareHouseNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			rp := warehousesrp.NewWarehousesRepositoryMock()
			sv := service.NewWarehousesService(rp)
			dto := models.WarehouseDTO{
				Address:            &tt.warehouse.Address,
				Telephone:          &tt.warehouse.Telephone,
				WarehouseCode:      &tt.warehouse.WarehouseCode,
				MinimumCapacity:    &tt.warehouse.MinimumCapacity,
				MinimumTemperature: &tt.warehouse.MinimumTemperature,
			}
			// Act
			rp.On("Update", 1, dto).Return(tt.warehouse, tt.err)
			warehouse, err := sv.Update(1, dto)

			// Assert
			require.Equal(t, tt.wantedWarehouse, warehouse)
			require.Equal(t, tt.wantedErr, err)
		})
	}
}

func TestWarehousesDefault_Delete(t *testing.T) {
	tests := []struct {
		name      string
		err       error
		wantedErr error
	}{
		{
			name:      "When the repository deletes the warehouse sucessfully",
			err:       nil,
			wantedErr: nil,
		},
		{
			name:      "When the repository returns an error",
			err:       models.ErrWareHouseNotFound,
			wantedErr: models.ErrWareHouseNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			rp := warehousesrp.NewWarehousesRepositoryMock()
			sv := service.NewWarehousesService(rp)

			// Act
			rp.On("Delete", 1).Return(tt.err)
			err := sv.Delete(1)

			// Assert
			require.Equal(t, tt.wantedErr, err)
		})
	}
}
