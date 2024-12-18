package models

type SectionService interface {
	GetAll() (sections map[int]Section, err error)
}
