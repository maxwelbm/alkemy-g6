package modelsBuyer

type BuyerRepository interface {
	GetAll() (buyers []Buyer, err error)
	GetById(id int) (buyer Buyer, err error)
}
