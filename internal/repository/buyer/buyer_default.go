package buyerRepository

import (
	"errors"

	modelsBuyer "github.com/maxwelbm/alkemy-g6/internal/models/buyer"
)

type BuyerRepository struct {
	Buyers map[int]modelsBuyer.Buyer
	NextID int
}

var ErrorIdNotFound = errors.New("Id not found")

func NewBuyerRepository(db map[int]modelsBuyer.Buyer) *BuyerRepository {
	repo := &BuyerRepository{
		Buyers: db,
		NextID: 1,
	}
	return repo
}
