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

type LocalitySellersReport struct {
	ID           int
	LocalityName string
	SellersCount int
}

type LocalityService interface {
	Create(locDTO LocalityDTO) (loc Locality, err error)
	ReportSellers(id int) (report LocalitySellersReport, err error)
}

type LocalityRepository interface {
	Create(locDTO LocalityDTO) (loc Locality, err error)
	ReportSellers(id int) (report LocalitySellersReport, err error)
}
