package sellerRepository

import modelsSeller "github.com/maxwelbm/alkemy-g6/internal/models/seller"

type SellerRepository struct {
	Sellers map[int]modelsSeller.Seller
	NextID  int
}

func NewSellerRepository(db map[int]modelsSeller.Seller) *SellerRepository {
	repo := &SellerRepository{
		Sellers: make(map[int]modelsSeller.Seller),
		NextID:  1,
	}
	//	repo.loadProducts(jsonPath)
	return repo
}
