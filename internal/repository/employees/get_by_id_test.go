package employeesrp

import (
	"testing"

	"github.com/maxwelbm/alkemy-g6/internal/factories"
	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/pkg/testdb"
	"github.com/stretchr/testify/require"
)

func TestEmployeesRepository_GetByID(t *testing.T) {
	db, truncate, teardown := testdb.NewConn(t)
	defer teardown()

	factory := factories.NewEmployeeFactory(db)
	employee := factory.Build(models.Employee{ID: 1})

	type expected struct {
		employees models.Employee
		err       error
	}

	tests := []struct {
		name  string
		setup func()
		id    int
		expected
	}{
		{
			name: "When the employee is found",
			setup: func() {
				_, err := factory.Create(employee)
				require.NoError(t, err)
			},
			id: 1,
			expected: expected{
				employees: employee,
			},
		},
		{
			name: "When the employee is not	found",
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
			got, err := rp.GetByID(tt.id)

			// Assert

			require.ErrorIs(t, err, tt.expected.err)
			require.Equal(t, tt.expected.employees, got)
		})
	}
}
