package repository

import "database/sql"

func (e *EmployeesRepository) Delete(id int) (err error) {
	query := "DELETE FROM employees WHERE id = ?"
	result, err := e.DB.Exec(query, id)
	if err != nil {
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
