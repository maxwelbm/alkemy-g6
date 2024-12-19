package modelsBuyer

type BuyerRepository interface {
	GetAll() (buyers []Buyer, err error)
	GetById(id int) (buyer Buyer, err error)
	GetByCardNumberId(cardNumberId string) (buyer Buyer, err error)
	PostBuyer(buyer Buyer) (buyerReturn Buyer, err error)
	PatchBuyer(buyer Buyer) (err error)
}