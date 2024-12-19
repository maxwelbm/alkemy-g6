package modelsSeller

type SellerService interface {
	GetAll() (sellers []Seller, err error)
	GetById(id int) (seller Seller, err error)
	GetByCid(cid int) (seller Seller, err error)
	PostSeller(seller Seller) error
	PatchSeller(seller Seller) error
	Delete(id int) (err error)
}
