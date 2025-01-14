package service

import (
	"github.com/maxwelbm/alkemy-g6/internal/models"
)

type SectionsController struct {
	rp models.SectionRepository
}

func NewSectionService(rp models.SectionRepository) *SectionsController {
	return &SectionsController{
		rp: rp,
	}
}

func (s *SectionsController) GetAll() (sections []models.Section, err error) {
	return s.rp.GetAll()

}

func (s *SectionsController) GetByID(id int) (section models.Section, err error) {
	return s.rp.GetByID(id)

}

func (s *SectionsController) GetReportProducts(sectionID int) (reportProducts []models.ProductReport, err error) {
	reportProducts, err = s.rp.GetReportProducts(sectionID)
	return
}

func (s *SectionsController) Create(sec models.SectionDTO) (newSection models.Section, err error) {
	return s.rp.Create(sec)

}

func (s *SectionsController) Update(id int, sec models.SectionDTO) (updateSection models.Section, err error) {
	return s.rp.Update(id, sec)

}

func (s *SectionsController) Delete(id int) (err error) {
	return s.rp.Delete(id)

}
