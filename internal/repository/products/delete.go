package repository

import "database/sql"

func (p *Products) Delete(id int) (err error) {
	query := "DELETE FROM products WHERE `id` = ?"
	result, err := p.DB.Exec(query, id)
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
