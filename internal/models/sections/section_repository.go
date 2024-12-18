package models

type SectionRepository interface {
	GetAll() (sections map[int]Section, err error)
}
