package employeesrp

import (
	"testing"

	"github.com/maxwelbm/alkemy-g6/internal/factories"
	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/pkg/testdb"
	"github.com/stretchr/testify/require"
)

func TestEmployeesRepository_GetReportInboundOrders(t *testing.T) {
	db, truncate, teardown := testdb.NewConn(t)
	defer teardown()

	type expected struct {
		employee []models.EmployeeReportInbound
		err      error
	}
	tests := []struct {
		name  string
		setup func()
		id    int
		expected
	}{
		{
			name: "When retrieving all employees inbound orders reports",
			setup: func() {
				factory := factories.NewEmployeeFactory(db)
				_, err := factory.Create(models.Employee{
					CardNumberID: "a12d",
					FirstName:    "Tiffany",
					LastName:     "Young",
					WarehouseID:  1,
				})
				require.NoError(t, err)
				_, err = factory.Create(models.Employee{
					CardNumberID: "a2221",
					FirstName:    "Lily",
					LastName:     "Mars",
					WarehouseID:  1,
				})
				require.NoError(t, err)
			},
			expected: expected{
				employee: []models.EmployeeReportInbound{
					{
						ID:           1,
						CardNumberID: "a12d",
						FirstName:    "Tiffany",
						LastName:     "Young",
						WarehouseID:  1,
						CountReports: 0,
					},
					{
						ID:           2,
						CardNumberID: "a2221",
						FirstName:    "Lily",
						LastName:     "Mars",
						WarehouseID:  1,
						CountReports: 0,
					},
				},
				err: nil,
			},
		},
		{
			name: "When retrieving single employee inbound orders report",
			setup: func() {
				factory := factories.NewEmployeeFactory(db)
				_, err := factory.Create(models.Employee{
					CardNumberID: "a12d",
					FirstName:    "Tiffany",
					LastName:     "Young",
					WarehouseID:  1,
				})
				require.NoError(t, err)
				_, err = factory.Create(models.Employee{
					CardNumberID: "a2221",
					FirstName:    "Lily",
					LastName:     "Mars",
					WarehouseID:  1,
				})
				require.NoError(t, err)
			},
			id: 1,
			expected: expected{
				employee: []models.EmployeeReportInbound{
					{
						ID:           1,
						CardNumberID: "a12d",
						FirstName:    "Tiffany",
						LastName:     "Young",
						WarehouseID:  1,
						CountReports: 0,
					},
				},
				err: nil,
			},
		},
		{
			name: "When reporting one inbound order by id and the employee has inbound orders",
			setup: func() {
				factory := factories.NewEmployeeFactory(db)
				_, err := factory.Create(models.Employee{
					CardNumberID: "a12d",
					FirstName:    "Tiffany",
					LastName:     "Young",
					WarehouseID:  1,
				})
				require.NoError(t, err)
				inboundOrdersFactory := factories.NewInboundOrderFactory(db)
				_, err = inboundOrdersFactory.Create(models.InboundOrder{OrderNumber: 1, EmployeeID: 1})
				require.NoError(t, err)
				_, err = inboundOrdersFactory.Create(models.InboundOrder{OrderNumber: 2, EmployeeID: 2})
				require.NoError(t, err)
			},
			id: 1,
			expected: expected{
				employee: []models.EmployeeReportInbound{
					{
						ID:           1,
						CardNumberID: "a12d",
						FirstName:    "Tiffany",
						LastName:     "Young",
						WarehouseID:  1,
						CountReports: 1,
					},
				},
				err: nil,
			},
		},
		{
			name: "When employee is not found",
			id:   1,
			expected: expected{
				err: models.ErrEmployeeNotFound,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup
			t.Cleanup(truncate)
			if tt.setup != nil {
				tt.setup()
			}

			// Arrange
			rp := NewEmployeesRepository(db)
			// Act
			got, err := rp.GetReportInboundOrders(tt.id)

			// Assert
			if tt.err != nil {
				require.Equal(t, tt.err.Error(), err.Error())
			}
			require.Equal(t, tt.expected.employee, got)
		})
	}
}
