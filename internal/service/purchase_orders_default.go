package service

import "github.com/maxwelbm/alkemy-g6/internal/models"

type PurchaseOrdersDefault struct {
	repo models.PurchaseOrdersRepository
}

func NewPurchaseOrdersService(rp models.PurchaseOrdersRepository) *PurchaseOrdersDefault {
	return &PurchaseOrdersDefault{repo: rp}
}

func (s *PurchaseOrdersDefault) Create(purchaseOrdersDTO models.PurchaseOrdersDTO) (po models.PurchaseOrders, err error) {
	po, err = s.repo.Create(purchaseOrdersDTO)
	return
}
