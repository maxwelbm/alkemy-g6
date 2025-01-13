package service

import (
	"github.com/maxwelbm/alkemy-g6/internal/models"
)

func NewProductRecordsService(repo models.ProductRecordsRepository) *ProductRecordDefault {
	return &ProductRecordDefault{repo: repo}
}

type ProductRecordDefault struct {
	repo models.ProductRecordsRepository
}

func (s *ProductRecordDefault) Create(warehouse models.ProductRecordDTO) (w models.ProductRecord, err error) {
	w, err = s.repo.Create(warehouse)
	return
}
