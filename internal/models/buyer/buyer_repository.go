package modelsBuyer

type BuyerRepository interface {
	GetAll() (buyers []Buyer, err error)
}
