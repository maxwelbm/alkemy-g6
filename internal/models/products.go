package models

import "errors"

var (
	ErrProductNotFound   = errors.New("Product not found")
)

type Product struct {
	// ID is the unique identifier of the product
	ID int
	// ProductCode is the unique code of the product
	ProductCode string
	// Description is the description of the product
	Description string
	// Height is the height of the product
	Height float64
	// Length is the length of the product
	Length float64
	// Width is the width of the product
	Width float64
	// Weight is the weight of the product
	Weight float64
	// ExpirationRate is the rate at which the product expires
	ExpirationRate float64
	// FreezingRate is the rate at which the product should be frozen
	FreezingRate float64
	// RecomFreezTemp is the recommended freezing temperature for the product
	RecomFreezTemp float64
	// ProductTypeID is the unique identifier of the product type
	ProductTypeID int
	// SellerID is the unique identifier of the seller
	SellerID int
}

type ProductDTO struct {
	ID             int
	ProductCode    string
	Description    string
	Height         float64
	Length         float64
	Width          float64
	Weight         float64
	ExpirationRate float64
	FreezingRate   float64
	RecomFreezTemp float64
	ProductTypeID  int
	SellerID       int
}

type ProductService interface {
	GetAll() (list []Product, err error)
	GetById(id int) (prod Product, err error)
	Create(prod ProductDTO) (newProd Product, err error)
	Update(id int, prod ProductDTO) (updatedProd Product, err error)
	Delete(id int) (err error)
}

type ProductRepository interface {
	GetAll() (list []Product, err error)
	GetById(id int) (prod Product, err error)
	Create(prod ProductDTO) (newProd Product, err error)
	Update(id int, prod ProductDTO) (updatedProd Product, err error)
	Delete(id int) (err error)
}