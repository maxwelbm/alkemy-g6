package service

import "github.com/maxwelbm/alkemy-g6/internal/models"

type LocalityDefault struct {
	rp models.LocalityRepository
}

func NewLocalitiesService(rp models.LocalityRepository) *LocalityDefault {
	return &LocalityDefault{rp: rp}
}

func (s *LocalityDefault) ReportSellers(id int) (reports []models.LocalitySellersReport, err error) {
	reports, err = s.rp.ReportSellers(id)
	return
}

func (s *LocalityDefault) ReportCarries(id int) (reports []models.LocalityCarriesReport, err error) {
	reports, err = s.rp.ReportCarries(id)
	return
}

func (s *LocalityDefault) Create(locDTO models.LocalityDTO) (loc models.Locality, err error) {
	loc, err = s.rp.Create(locDTO)
	return
}
