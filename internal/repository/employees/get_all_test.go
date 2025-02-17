package employeesrp

import (
	"testing"

	"github.com/maxwelbm/alkemy-g6/internal/factories"
	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/pkg/testdb"
	"github.com/stretchr/testify/require"
)

func TestEmployeesRepository_GetAll(t *testing.T) {
	db, truncate, teardown := testdb.NewConn(t)
	defer teardown()

	factory := factories.NewEmployeeFactory(db)
	employee1 := factory.Build(models.Employee{ID: 1})
	employee2 := factory.Build(models.Employee{ID: 2})

	type expected struct {
		employees []models.Employee
	}

	tests := []struct {
		name  string
		setup func()
		expected
	}{
		{
			name: "When employees are registered",
			setup: func() {
				_, err := factory.Create(employee1)
				require.NoError(t, err)
				_, err = factory.Create(employee2)
				require.NoError(t, err)
			},
			expected: expected{
				employees: []models.Employee{employee1, employee2},
			},
		},
		{
			name: "When no employees are registered",
			expected: expected{
				employees: []models.Employee(nil),
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
			got, err := rp.GetAll()

			// Assert

			require.NoError(t, err)
			require.Equal(t, tt.expected.employees, got)
		})
	}
}
