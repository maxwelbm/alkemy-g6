package models

type SectionRepository interface {
	GetAll() (sections []Section, err error)
}
