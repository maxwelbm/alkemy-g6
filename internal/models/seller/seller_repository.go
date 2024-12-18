package modelsSeller

type SellerRepository interface {
	GetAll() (sellerMap []Seller, err error)
}
