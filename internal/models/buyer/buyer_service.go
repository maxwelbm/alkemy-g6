package modelsBuyer

type BuyerService interface {
	GetAll() (buyers []Buyer, err error)
}
