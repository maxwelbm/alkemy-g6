package repository

import (
	"github.com/maxwelbm/alkemy-g6/internal/models"
)

func (e *EmployeesRepository) GetAll() (employees []models.Employees, err error) {
	query := "SELECT id, card_number_id, first_name, last_name, warehouse_id FROM employees"

	rows, err := e.DB.Query(query)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var employee models.Employees
		if err = rows.Scan(&employee.ID, &employee.CardNumberID, &employee.FirstName, &employee.LastName, &employee.WarehouseID); err != nil {
			return
		}
		employees = append(employees, employee)
	}

	// Check for errors after rows iteration
	if err = rows.Err(); err != nil {
		return
	}

	return
}
