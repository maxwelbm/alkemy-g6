package service

import (
	"github.com/maxwelbm/alkemy-g6/internal/models"
)

type SectionsDefault struct {
	rp models.SectionRepository
}

func NewSectionsService(rp models.SectionRepository) *SectionsDefault {
	return &SectionsDefault{
		rp: rp,
	}
}

func (s *SectionsDefault) GetAll() (sections []models.Section, err error) {
	return s.rp.GetAll()
}

func (s *SectionsDefault) GetByID(id int) (section models.Section, err error) {
	return s.rp.GetByID(id)
}

func (s *SectionsDefault) GetReportProducts(sectionID int) (reportProducts []models.ProductReport, err error) {
	reportProducts, err = s.rp.GetReportProducts(sectionID)
	return
}

func (s *SectionsDefault) Create(sec models.SectionDTO) (newSection models.Section, err error) {
	return s.rp.Create(sec)
}

func (s *SectionsDefault) Update(id int, sec models.SectionDTO) (updateSection models.Section, err error) {
	return s.rp.Update(id, sec)
}

func (s *SectionsDefault) Delete(id int) (err error) {
	return s.rp.Delete(id)
}
