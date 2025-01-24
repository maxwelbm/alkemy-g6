package service

import "github.com/maxwelbm/alkemy-g6/internal/models"

type ProductBatchesService struct {
	rp models.ProductBatchesRepository
}

func NewProductBatchesService(rp models.ProductBatchesRepository) *ProductBatchesService {
	return &ProductBatchesService{rp: rp}
}

func (s *ProductBatchesService) Create(prodBatches models.ProductBatchesDTO) (newProdBatches models.ProductBatches, err error) {
	newProdBatches, err = s.rp.Create(prodBatches)
	return
}
