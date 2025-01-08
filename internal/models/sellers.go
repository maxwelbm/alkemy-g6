package models

import "errors"

var (
	ErrSellerNotFound = errors.New("Seller not found")
)

type Seller struct {
	ID          int
	CID         string
	CompanyName string
	Address     string
	Telephone   string
}

type SellerDTO struct {
	ID          int    `json:"id,omitempty"`
	CID         string `json:"cid"`
	CompanyName string `json:"company_name"`
	Address     string `json:"address"`
	Telephone   string `json:"telephone"`
}

type SellersService interface {
	GetAll() (sellers []Seller, err error)
	GetById(id int) (seller Seller, err error)
	GetByCid(cid int) (seller Seller, err error)
	Create(seller SellerDTO) (sellerReturn Seller, err error)
	Update(id int, seller SellerDTO) (sellerReturn Seller, err error)
	Delete(id int) (err error)
}

type SellersRepository interface {
	GetAll() (sellerMap []Seller, err error)
	GetById(id int) (seller Seller, err error)
	GetByCid(cid int) (seller Seller, err error)
	Create(seller SellerDTO) (sellerReturn Seller, err error)
	Update(id int, seller SellerDTO) (sellerReturn Seller, err error)
	Delete(id int) (err error)
}
