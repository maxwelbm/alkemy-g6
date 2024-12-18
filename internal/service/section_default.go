package service

import (
	models "github.com/maxwelbm/alkemy-g6/internal/models/sections"
	repository "github.com/maxwelbm/alkemy-g6/internal/repository/sections"
)

type SectionsDefault struct {
	repo repository.Sections
}

func NewSectionDefault(repo repository.Sections) *SectionsDefault {
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
