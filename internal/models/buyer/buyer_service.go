package modelsBuyer

type BuyerService interface {
	GetAll() (buyers []Buyer, err error)
	GetById(id int) (buyer Buyer, err error)
	GetByCardNumberId(cardNumberId string) (buyer Buyer, err error)
	PostBuyer(buyer Buyer) (buyerReturn Buyer, err error)
	PatchBuyer(buyer BuyerDTO) (buyerReturn Buyer, err error)
}
