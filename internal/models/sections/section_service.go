package models

type SectionService interface {
	GetAll() (sections []Section, err error)
	GetById(id int) (section Section, err error)
	Create(sec SectionDTO) (newSection Section, err error)
	Delete(id int) (err error)
}
