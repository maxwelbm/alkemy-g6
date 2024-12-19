package modelsBuyer

type BuyerService interface {
	GetAll() (buyers []Buyer, err error)
	GetById(id int) (buyer Buyer, err error)
}
