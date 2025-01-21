package employeesrp

import models "github.com/maxwelbm/alkemy-g6/internal/models"

func (e *EmployeesRepository) Delete(id int) (err error) {
	var exists bool
	err = e.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM employees WHERE id = ?)", id).Scan(&exists)

	if err != nil {
		return
	}

	if !exists {
		err = models.ErrEmployeeNotFound
		return
	}

	query := "DELETE FROM employees WHERE id = ?"

	_, err = e.DB.Exec(query, id)
	if err != nil {
		return
	}

	return nil
}
