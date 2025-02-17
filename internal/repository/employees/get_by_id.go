package employeesrp

import (
	"database/sql"
	"errors"

	models "github.com/maxwelbm/alkemy-g6/internal/models"
)

func (e *EmployeesRepository) GetByID(id int) (employees models.Employee, err error) {
	query := "SELECT id, card_number_id, first_name, last_name, warehouse_id FROM employees WHERE id = ?"

	row := e.DB.QueryRow(query, id)

	err = row.Scan(&employees.ID, &employees.CardNumberID, &employees.FirstName, &employees.LastName, &employees.WarehouseID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = models.ErrEmployeeNotFound
			return
		}

		return
	}

	return
}
