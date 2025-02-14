package factories

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/pkg/randstr"
)

type BuyerFactory struct {
	db *sql.DB
}

func NewBuyerFactory(db *sql.DB) *BuyerFactory {
	return &BuyerFactory{db: db}
}

func defaultBuyer() models.Buyer {
	return models.Buyer{
		CardNumberID: randstr.Alphanumeric(8),
		FirstName:    randstr.Chars(8),
		LastName:     randstr.Chars(8),
	}
}

func (f BuyerFactory) Build(buyer models.Buyer) models.Buyer {
	populateBuyerParams(&buyer)

	return buyer
}

func (f *BuyerFactory) Create(buyer models.Buyer) (record models.Buyer, err error) {
	populateBuyerParams(&buyer)

	query := `
		INSERT INTO buyers 
			(
			%s
			card_number_id, 
			first_name, 
			last_name
			) 
		VALUES (%s?, ?, ?)
	`

	switch buyer.ID {
	case 0:
		query = fmt.Sprintf(query, "", "")
	default:
		query = fmt.Sprintf(query, "id,", strconv.Itoa(buyer.ID)+",")
	}

	result, err := f.db.Exec(query,
		buyer.CardNumberID,
		buyer.FirstName,
		buyer.LastName,
	)

	if err != nil {
		return record, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return record, err
	}

	buyer.ID = int(id)

	return buyer, err
}

func populateBuyerParams(buyer *models.Buyer) {
	defaultBuyer := defaultBuyer()
	if buyer == nil {
		buyer = &defaultBuyer
	}

	if buyer.CardNumberID == "" {
		buyer.CardNumberID = defaultBuyer.CardNumberID
	}

	if buyer.FirstName == "" {
		buyer.FirstName = defaultBuyer.FirstName
	}

	if buyer.LastName == "" {
		buyer.LastName = defaultBuyer.LastName
	}
}
