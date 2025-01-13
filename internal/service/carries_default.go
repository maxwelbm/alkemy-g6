package service

import (
	"github.com/maxwelbm/alkemy-g6/internal/models"
)

type CarriesDefault struct {
	rp models.CarriesRepository
}

func NewCarriesService(rp models.CarriesRepository) *CarriesDefault {
	return &CarriesDefault{
		rp: rp,
	}
}

func (s *CarriesDefault) Create(carry models.CarryDTO) (carryToReturn models.Carry, err error) {
	carryToReturn, err = s.rp.Create(carry)
	return
}
