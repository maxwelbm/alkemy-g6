package models

type ProductRepository interface {
	GetAll() (list map[int]Product, err error)
}
