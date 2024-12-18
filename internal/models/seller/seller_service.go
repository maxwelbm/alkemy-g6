package modelsSeller

type SellerService interface {
	GetAll() (sellers []Seller, err error)
}
