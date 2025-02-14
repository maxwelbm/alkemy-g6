package employeesrp

import (
	"testing"

	"github.com/go-sql-driver/mysql"
	"github.com/maxwelbm/alkemy-g6/internal/factories"
	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/pkg/mysqlerr"
	"github.com/maxwelbm/alkemy-g6/pkg/testdb"
	"github.com/stretchr/testify/require"
)

func TestEmployeesRepository_Create(t *testing.T) {
	db, truncate, teardown := testdb.NewConn(t)
	defer teardown()

	factory := factories.NewEmployeeFactory(db)
	fixture := factory.Build(models.Employee{ID: 1})

	type arg struct {
		dto models.EmployeeDTO
	}

	type expected struct {
		employee models.Employee
		err      error
	}

	tests := []struct {
		name  string
		setup func()
		arg
		expected
	}{
		{
			name: "When successfully creating a new Employee",
			setup: func() {
				_, err := factories.NewWarehouseFactory(db).Create(models.Warehouse{ID: 1})
				require.NoError(t, err)
			},
			arg: arg{
				dto: models.EmployeeDTO{
					ID:           &fixture.ID,
					CardNumberID: &fixture.CardNumberID,
					FirstName:    &fixture.FirstName,
					LastName:     &fixture.LastName,
					WarehouseID:  &fixture.WarehouseID,
				},
			},
			expected: expected{
				employee: fixture,
				err:      nil,
			},
		},
		{
			name: "Error - When creating a duplicated Employee",
			setup: func() {
				_, err := factory.Create(fixture)
				require.NoError(t, err)
			},
			arg: arg{
				dto: models.EmployeeDTO{
					ID:           &fixture.ID,
					CardNumberID: &fixture.CardNumberID,
					FirstName:    &fixture.FirstName,
					LastName:     &fixture.LastName,
					WarehouseID:  &fixture.WarehouseID,
				},
			},
			expected: expected{
				err: &mysql.MySQLError{Number: mysqlerr.CodeDuplicateEntry},
			},
		},
		{
			name: "Error - When creating an Employee with a non-existent Warehouse",
			arg: arg{
				dto: models.EmployeeDTO{
					ID:           &fixture.ID,
					CardNumberID: &fixture.CardNumberID,
					FirstName:    &fixture.FirstName,
					LastName:     &fixture.LastName,
					WarehouseID:  &fixture.WarehouseID,
				},
			},
			expected: expected{
				err: &mysql.MySQLError{Number: mysqlerr.CodeCannotAddOrUpdateChildRow},
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
			got, err := rp.Create(tt.dto)

			// Assert
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
			}
			require.Equal(t, tt.expected.employee, got)
		})
	}
}
