package repository

import (
	"errors"
	"fmt"

	modelsBuyer "github.com/maxwelbm/alkemy-g6/internal/models/buyer"
)

func (r *BuyerRepository) GetById(id int) (buyer modelsBuyer.Buyer, err error) {
	buyer, ok := r.Buyers[id]
	if !ok {
		err = errors.New(fmt.Sprintf("Id %d not found in the base!", id))
	}
	return
}
