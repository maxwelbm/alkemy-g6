package service

import (
	"errors"

	models "github.com/maxwelbm/alkemy-g6/internal/models/sections"
	"github.com/maxwelbm/alkemy-g6/internal/repository"
)

type SectionsDefault struct {
	repo repository.RepoDB
}

var (
	ErrWareHousesNotFound = errors.New("Warehouse not found")
)

func NewSectionDefault(repo repository.RepoDB) *SectionsDefault {
	return &SectionsDefault{
		repo: repo,
	}
}

func (s *SectionsDefault) GetAll() (sections []models.Section, err error) {
	sections, err = s.repo.SectionsDB.GetAll()
	return
}

func (s *SectionsDefault) GetById(id int) (section models.Section, err error) {
	section, err = s.repo.SectionsDB.GetById(id)
	return
}

func (s *SectionsDefault) Create(sec models.SectionDTO) (newSection models.Section, err error) {
	if _, err = s.repo.WarehouseDB.GetById(*sec.WarehouseID); err != nil {
		err = ErrWareHousesServiceNotFound
		return
	}

	newSection, err = s.repo.SectionsDB.Create(sec)
	return
}

func (s *SectionsDefault) Update(id int, sec models.SectionDTO) (updateSection models.Section, err error) {
	if sec.WarehouseID != nil {
		if _, err = s.repo.WarehouseDB.GetById(*sec.WarehouseID); err != nil {
			err = ErrWareHousesNotFound
			return
		}
	}

	updateSection, err = s.repo.SectionsDB.Update(id, sec)
	return
}

func (s *SectionsDefault) Delete(id int) (err error) {
	err = s.repo.SectionsDB.Delete(id)
	return
}
