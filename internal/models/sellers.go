package models

import "errors"

var (
	ErrorNoChangesMade = errors.New("no changes made")
	ErrSellerNotFound  = errors.New("seller not found")
)

type Seller struct {
	ID          int
	CID         string
	CompanyName string
	Address     string
	Telephone   string
	LocalityID  int
}

type SellerDTO struct {
	ID          int    `json:"id,omitempty"`
	CID         string `json:"cid"`
	CompanyName string `json:"company_name"`
	Address     string `json:"address"`
	Telephone   string `json:"telephone"`
	LocalityID  int    `json:"locality_id"`
}

type SellersService interface {
	GetAll() (sellers []Seller, err error)
	GetByID(id int) (seller Seller, err error)
	GetByCid(cid int) (seller Seller, err error)
	Create(seller SellerDTO) (sellerReturn Seller, err error)
	Update(id int, seller SellerDTO) (sellerReturn Seller, err error)
	Delete(id int) (err error)
}

type SellersRepository interface {
	GetAll() (sellerMap []Seller, err error)
	GetByID(id int) (seller Seller, err error)
	GetByCid(cid int) (seller Seller, err error)
	Create(seller SellerDTO) (sellerReturn Seller, err error)
	Update(id int, seller SellerDTO) (sellerReturn Seller, err error)
	Delete(id int) (err error)
}
