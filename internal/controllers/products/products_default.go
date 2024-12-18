package controller

import (
	"errors"
	"fmt"

	models "github.com/maxwelbm/alkemy-g6/internal/models/products"
)

type ProductsDefault struct {
	sv models.ProductService
}

func NewProductsDefault(sv models.ProductService) *ProductsDefault {
	return &ProductsDefault{sv: sv}
}

type ProductResJSON struct {
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}

type ProductFullJSON struct {
	ID             int     `json:"id"`
	ProductCode    string  `json:"product_code"`
	Description    string  `json:"description"`
	Height         float64 `json:"height"`
	Length         float64 `json:"length"`
	Width          float64 `json:"width"`
	Weight         float64 `json:"weight"`
	ExpirationRate float64 `json:"expiration_rate"`
	FreezingRate   float64 `json:"freezing_rate"`
	RecomFreezTemp float64 `json:"recommended_freezing_temp"`
	ProductTypeID  int     `json:"product_type_id"`
	SellerID       int     `json:"seller_id"`
}

type NewProductAttributesJSON struct {
	ProductCode    *string  `json:"product_code"`
	Description    *string  `json:"description"`
	Height         *float64 `json:"height"`
	Length         *float64 `json:"length"`
	Width          *float64 `json:"width"`
	Weight         *float64 `json:"weight"`
	ExpirationRate *float64 `json:"expiration_rate"`
	FreezingRate   *float64 `json:"freezing_rate"`
	RecomFreezTemp *float64 `json:"recommended_freezing_temp"`
	ProductTypeID  *int     `json:"product_type_id"`
	SellerID       *int     `json:"seller_id,omitempty"`
}

func (p *NewProductAttributesJSON) validate() (err error) {
	var validationErrors []string
	var nilPointerErrors []string

	// Check for nil pointers and collect their errors
	if p.ProductCode == nil {
		nilPointerErrors = append(nilPointerErrors, "error: attribute ProductCode cannot be nil")
	} else if *p.ProductCode == "" {
		validationErrors = append(validationErrors, "error: attribute ProductCode cannot be empty")
	}

	if p.Description == nil {
		nilPointerErrors = append(nilPointerErrors, "error: attribute Description cannot be nil")
	} else if *p.Description == "" {
		validationErrors = append(validationErrors, "error: attribute Description cannot be empty")
	}

	if p.Height == nil {
		nilPointerErrors = append(nilPointerErrors, "error: attribute Height cannot be nil")
	} else if *p.Height <= 0 {
		validationErrors = append(validationErrors, "error: attribute Height cannot be negative")
	}

	if p.Length == nil {
		nilPointerErrors = append(nilPointerErrors, "error: attribute Length cannot be nil")
	} else if *p.Length <= 0 {
		validationErrors = append(validationErrors, "error: attribute Length cannot be negative")
	}

	if p.Width == nil {
		nilPointerErrors = append(nilPointerErrors, "error: attribute Width cannot be nil")
	} else if *p.Width <= 0 {
		validationErrors = append(validationErrors, "error: attribute Width cannot be negative")
	}

	if p.Weight == nil {
		nilPointerErrors = append(nilPointerErrors, "error: attribute Weight cannot be nil")
	} else if *p.Weight <= 0 {
		validationErrors = append(validationErrors, "error: attribute Weight cannot be negative")
	}

	if p.ExpirationRate == nil {
		nilPointerErrors = append(nilPointerErrors, "error: attribute ExpirationRate cannot be nil")
	} else if *p.ExpirationRate < 0 {
		validationErrors = append(validationErrors, "error: attribute ExpirationRate cannot be negative")
	}

	if p.FreezingRate == nil {
		nilPointerErrors = append(nilPointerErrors, "error: attribute FreezingRate cannot be nil")
	} else if *p.FreezingRate < 0 {
		validationErrors = append(validationErrors, "error: attribute FreezingRate cannot be negative")
	}

	if p.RecomFreezTemp == nil {
		nilPointerErrors = append(nilPointerErrors, "error: attribute RecomFreezTemp cannot be nil")
	}

	if p.ProductTypeID == nil {
		nilPointerErrors = append(nilPointerErrors, "error: attribute ProductTypeID cannot be nil")
	} else if *p.ProductTypeID <= 0 {
		validationErrors = append(validationErrors, "error: attribute ProductTypeID must be positive")
	}

	if p.SellerID != nil && *p.SellerID < 0 {
		validationErrors = append(validationErrors, "error: attribute SellerID must be non-negative")
	}

	// Combine all errors before returning
	if len(nilPointerErrors) > 0 || len(validationErrors) > 0 {
		var allErrors []string
		allErrors = append(allErrors, nilPointerErrors...)
		allErrors = append(allErrors, validationErrors...)

		err = errors.New(fmt.Sprintf("validation errors: %v", allErrors))
	}
	return
}
