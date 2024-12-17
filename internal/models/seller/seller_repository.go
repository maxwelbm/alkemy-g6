package modelsSeller

type SellerRepository interface {
	FindAll() (sellerMap map[int]Seller, err error)
}
