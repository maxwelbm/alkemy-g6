package models

type SectionService interface {
	GetAll() (sections []Section, err error)
}
