package models

import "errors"

var (
	ErrCarryNotFound = errors.New("carry not found")
)

type Carry struct {
	ID          int
	CID         string
	CompanyName string
	Address     string
	PhoneNumber string
	LocalityID  int
}

type CarryDTO struct {
	CID         *string
	CompanyName *string
	Address     *string
	PhoneNumber *string
	LocalityID  *int
}

type CarriesService interface {
	Create(carry CarryDTO) (carryReturn Carry, err error)
}

type CarriesRepository interface {
	Create(carry CarryDTO) (carryReturn Carry, err error)
}
