package service

import "github.com/maxwelbm/alkemy-g6/internal/models"

type ProductBatchesDefault struct {
	rp models.ProductBatchesRepository
}

func NewProductBatchesDefault(rp models.ProductBatchesRepository) *ProductBatchesDefault {
	return &ProductBatchesDefault{rp: rp}
}

func (s *ProductBatchesDefault) Create(prodBatches models.ProductBatchesDTO) (newProdBatches models.ProductBatches, err error) {
	newProdBatches, err = s.rp.Create(prodBatches)
	return
}
