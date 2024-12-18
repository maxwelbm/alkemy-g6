package models

type SectionService interface {
	GetAll() (sections []Section, err error)
	GetById(id int) (section Section, err error)
}
