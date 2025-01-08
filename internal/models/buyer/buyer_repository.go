package models

type BuyerRepository interface {
	GetAll() (buyers []Buyer, err error)
	GetById(id int) (buyer Buyer, err error)
	GetByCardNumberId(cardNumberId string) (buyer Buyer, err error)
	Create(buyer Buyer) (buyerReturn Buyer, err error)
	Update(buyer BuyerDTO) (buyerReturn Buyer, err error)
	Delete(id int) (err error)
}
