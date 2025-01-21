package productsrp

import (
	"github.com/maxwelbm/alkemy-g6/internal/models"
)

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
		err = models.ErrProductNotFound
		return
	}
	
	return nil
}
