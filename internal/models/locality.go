package models

import "errors"

var (
	ErrLocalityNotFound = errors.New("Locality not found")
)

type Locality struct {
	ID           int
	LocalityName string
	ProvinceName string
	CountryName  string
}

type LocalityDTO struct {
	LocalityName string
	ProvinceName string
	CountryName  string
}

type LocalityService interface {
	Create(locDTO LocalityDTO) (loc Locality, err error)
}

type LocalityRepository interface {
	Create(locDTO LocalityDTO) (loc Locality, err error)
}
