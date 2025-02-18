package employeesrp

import (
	"github.com/maxwelbm/alkemy-g6/internal/models"
)

func (e *EmployeesRepository) Create(employees models.EmployeeDTO) (newEmployees models.Employee, err error) {
	query := "INSERT INTO employees (card_number_id, first_name, last_name, warehouse_id) VALUES (?, ?, ?, ?)"

	result, err := e.DB.Exec(query, employees.CardNumberID, employees.FirstName, employees.LastName, employees.WarehouseID)
	if err != nil {
		return
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return
	}

	query = "SELECT id, card_number_id, first_name, last_name, warehouse_id FROM employees WHERE id = ?"
	err = e.DB.
		QueryRow(query, lastInsertID).
		Scan(&newEmployees.ID, &newEmployees.CardNumberID, &newEmployees.FirstName, &newEmployees.LastName, &newEmployees.WarehouseID)

	if err != nil {
		return
	}

	return
}
