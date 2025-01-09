package models

import "errors"

var (
	ErrBuyerNotFound = errors.New("Buyer not found ")
)

type Buyer struct {
	Id           int
	CardNumberId string
	FirstName    string
	LastName     string
}

type BuyerDTO struct {
	Id           *int    `json:"id,omitempty"`
	CardNumberId *string `json:"card_number_id,omitempty"`
	FirstName    *string `json:"first_name,omitempty"`
	LastName     *string `json:"last_name,omitempty"`
}

type BuyerService interface {
	GetAll() (buyers []Buyer, err error)
	GetById(id int) (buyer Buyer, err error)
	GetByCardNumberId(cardNumberId string) (buyer Buyer, err error)
	Create(buyer BuyerDTO) (buyerReturn Buyer, err error)
	Update(id int, buyer BuyerDTO) (buyerReturn Buyer, err error)
	Delete(id int) (err error)
}

type BuyerRepository interface {
	GetAll() (buyers []Buyer, err error)
	GetById(id int) (buyer Buyer, err error)
	GetByCardNumberId(cardNumberId string) (buyer Buyer, err error)
	Create(buyer BuyerDTO) (buyerReturn Buyer, err error)
	Update(id int, buyer BuyerDTO) (buyerReturn Buyer, err error)
	Delete(id int) (err error)
}
