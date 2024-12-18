package models

type SectionRepository interface {
	GetAll() (sections []Section, err error)
	GetById(id int) (section Section, err error)
	Create(sec SectionDTO) (newSection Section, err error)
}
