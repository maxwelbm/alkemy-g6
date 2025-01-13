package repository

import (
	"database/sql"
	"errors"


	models "github.com/maxwelbm/alkemy-g6/internal/models"

)

func (r *EmployeesRepository) GetByID(id int) (employees models.Employees, err error) {
	query := "SELECT id, card_number_id, first_name, last_name, warehouse_id FROM employees WHERE id = ?"

	row := r.DB.QueryRow(query, id)

	err = row.Scan(&employees.ID, &employees.CardNumberID, &employees.FirstName, &employees.LastName, &employees.WarehouseID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = errors.New("Id not found")
			return
		}
		return
	}

	return
}
