package modelsSeller

type SellerService interface {
	FindAll() (sellers map[int]Seller, err error)
}
