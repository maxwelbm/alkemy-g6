package repository

import (

	models "github.com/maxwelbm/alkemy-g6/internal/models"

)

func (e *EmployeesRepository) Update(employees models.EmployeesDTO, id int) (newEmployees models.Employees, err error) {
	// Check if the employee exists
	var exists bool
	err = e.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM employees WHERE id = ?)", id).Scan(&exists)
	if err != nil {
		return
	}
	if !exists {
		err = models.ErrEmployeesNotFound
		return
	}
	// Update the employees
	query := `UPDATE employees SET 
		card_number_id = COALESCE(NULLIF(?, ''), card_number_id), 
		first_name = COALESCE(NULLIF(?, ''), first_name),
		last_name = COALESCE(NULLIF(?, ''), last_name),
		warehouse_id = COALESCE(NULLIF(?, ''), warehouse_id)
	WHERE id = ?`
	res, err := e.DB.Exec(query, employees.CardNumberID, employees.FirstName, employees.LastName, employees.WarehouseID, id)
	// Check for errors
	if err != nil {
		return
	}
	// Check if the employee was updated
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return
	}
	// If the employee was not updated, return an error
	if rowsAffected == 0 {
		err = models.ErrorNoChangesMade
		return
	}

	// Retrieve the updated employee
	err = e.DB.QueryRow("SELECT id, card_number_id, first_name, last_name, warehouse_id FROM employees WHERE id = ?", id).Scan(
		&newEmployees.ID, &newEmployees.CardNumberID, &newEmployees.FirstName, &newEmployees.LastName, &newEmployees.WarehouseID)
	if err != nil {
		return
	}

	return
}
