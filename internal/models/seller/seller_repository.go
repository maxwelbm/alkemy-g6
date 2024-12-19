package modelsSeller

type SellerRepository interface {
	GetAll() (sellerMap []Seller, err error)
	GetById(id int) (sellerMap Seller, err error)
	PatchSeller(seller Seller) error
	Delete(id int) (err error)
}
