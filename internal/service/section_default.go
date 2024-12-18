package service

import (
	models "github.com/maxwelbm/alkemy-g6/internal/models/sections"
	section_repository "github.com/maxwelbm/alkemy-g6/internal/repository/sections"
)

type SectionsDefault struct {
	repo section_repository.Sections
}

func NewSectionDefault(repo section_repository.Sections) *SectionsDefault {
	return &SectionsDefault{
		repo: repo,
	}
}

func (s *SectionsDefault) GetAll() (sections map[int]models.Section, err error) {
	sections, err = s.repo.GetAll()
	return
}
