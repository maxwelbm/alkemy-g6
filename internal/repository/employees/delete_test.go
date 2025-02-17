package employeesrp

import (
	"testing"

	"github.com/maxwelbm/alkemy-g6/internal/factories"
	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/pkg/testdb"
	"github.com/stretchr/testify/require"
)

func TestEmployeesRepository_Delete(t *testing.T) {
	db, truncate, teardown := testdb.NewConn(t)
	defer teardown()

	factory := factories.NewEmployeeFactory(db)

	type expected struct {
		err error
	}

	tests := []struct {
		name  string
		setup func()
		id    int
		expected
	}{
		{
			name: "When successfully deleting a new Employee",
			setup: func() {
				_, err := factory.Create(models.Employee{ID: 1})
				require.NoError(t, err)
			},
			id: 1,
			expected: expected{
				err: nil,
			},
		},
		{
			name: "Error - When trying to delete a employee that does not exist",
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
			err := rp.Delete(tt.id)

			// Assert
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
			}
		})
	}
}
