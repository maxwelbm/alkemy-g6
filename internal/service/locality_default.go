package service

import "github.com/maxwelbm/alkemy-g6/internal/models"

type LocalityDefault struct {
	rp models.LocalityRepository
}

func NewLocalityDefault(rp models.LocalityRepository) *LocalityDefault {
	return &LocalityDefault{rp: rp}
}

func (s *LocalityDefault) ReportSellers(id int) (report models.LocalitySellersReport, err error) {
	report, err = s.rp.ReportSellers(id)
	return
}

func (s *LocalityDefault) Create(locDTO models.LocalityDTO) (loc models.Locality, err error) {
	loc, err = s.rp.Create(locDTO)
	return
}
