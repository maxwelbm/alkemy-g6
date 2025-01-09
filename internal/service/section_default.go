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
	sections, err = s.rp.GetAll()
	return
}

func (s *SectionsController) GetById(id int) (section models.Section, err error) {
	section, err = s.rp.GetById(id)
	return
}

func (s *SectionsController) Create(sec models.SectionDTO) (newSection models.Section, err error) {
	// if _, err = s.rp.WarehouseDB.GetById(*sec.WarehouseID); err != nil {
	// 	err = ErrWareHousesServiceNotFound
	// 	return
	// }

	newSection, err = s.rp.Create(sec)
	return
}

func (s *SectionsController) Update(id int, sec models.SectionDTO) (updateSection models.Section, err error) {
	// if sec.WarehouseID != nil {
	// 	if _, err = s.rp.WarehouseDB.GetById(*sec.WarehouseID); err != nil {
	// 		err = ErrWareHousesNotFound
	// 		return
	// 	}
	// }

	updateSection, err = s.rp.Update(id, sec)
	return
}

func (s *SectionsController) Delete(id int) (err error) {
	err = s.rp.Delete(id)
	return
}
