package service

import (
	models "github.com/maxwelbm/alkemy-g6/internal/models/sections"
)

type SectionsDefault struct {
	repo models.SectionRepository
}

func NewSectionDefault(repo models.SectionRepository) *SectionsDefault {
	return &SectionsDefault{
		repo: repo,
	}
}

func (s *SectionsDefault) GetAll() (sections []models.Section, err error) {
	sections, err = s.repo.GetAll()
	return
}

func (s *SectionsDefault) GetById(id int) (section models.Section, err error) {
	section, err = s.repo.GetById(id)
	return
}

func (s *SectionsDefault) Create(sec models.SectionDTO) (newSection models.Section, err error) {
	newSection, err = s.repo.Create(sec)
	return
}

func (s *SectionsDefault) Update(id int, sec models.SectionDTO) (updateSection models.Section, err error) {
	updateSection, err = s.repo.Update(id, sec)
	return
}

func (s *SectionsDefault) Delete(id int) (err error) {
	err = s.repo.Delete(id)
	return
}
