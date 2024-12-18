package modelsSeller

type SellerService interface {
	GetAll() (sellers []Seller, err error)
	GetById(id int) (seller Seller, err error)
	PatchSeller(seller Seller) error
	Delete(id int) (err error)
}
