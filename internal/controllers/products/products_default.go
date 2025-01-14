package productsctl

import (
	"errors"
	"fmt"

	"github.com/maxwelbm/alkemy-g6/internal/models"
)

type ProductsDefault struct {
	SV models.ProductService
}

func NewProductsController(sv models.ProductService) *ProductsDefault {
	return &ProductsDefault{SV: sv}
}

type ProductResJSON struct {
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}

type ReportRecordsResJSON struct {
	Data any `json:"data,omitempty"`
}

type ProductFullJSON struct {
	ID             int     `json:"id"`
	ProductCode    string  `json:"product_code"`
	Description    string  `json:"description"`
	Height         float64 `json:"height"`
	Length         float64 `json:"length"`
	Width          float64 `json:"width"`
	NetWeight      float64 `json:"net_weight"`
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
	NetWeight      *float64 `json:"net_weight"`
	ExpirationRate *float64 `json:"expiration_rate"`
	FreezingRate   *float64 `json:"freezing_rate"`
	RecomFreezTemp *float64 `json:"recommended_freezing_temp"`
	ProductTypeID  *int     `json:"product_type_id"`
	SellerID       *int     `json:"seller_id,omitempty"`
}

type ReportRecordFullJSON struct {
	ProductId    int    `json:"product_id"`
	Description  string `json:"description"`
	RecordsCount int    `json:"records_count"`
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

	if p.NetWeight == nil {
		nilPointerErrors = append(nilPointerErrors, "error: attribute NetWeight cannot be nil")
	} else if *p.NetWeight <= 0 {
		validationErrors = append(validationErrors, "error: attribute NetWeight cannot be negative")
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

type UpdateProductAttributesJSON struct {
	ProductCode    *string  `json:"product_code,omitempty"`
	Description    *string  `json:"description,omitempty"`
	Height         *float64 `json:"height,omitempty"`
	Length         *float64 `json:"length,omitempty"`
	Width          *float64 `json:"width,omitempty"`
	NetWeight      *float64 `json:"net_weight,omitempty"`
	ExpirationRate *float64 `json:"expiration_rate,omitempty"`
	FreezingRate   *float64 `json:"freezing_rate,omitempty"`
	RecomFreezTemp *float64 `json:"recommended_freezing_temp,omitempty"`
	ProductTypeID  *int     `json:"product_type_id,omitempty"`
	SellerID       *int     `json:"seller_id,omitempty"`
}

func (p *UpdateProductAttributesJSON) validate() (err error) {
	var validationErrors []string

	if p.ProductCode != nil && *p.ProductCode == "" {
		validationErrors = append(validationErrors, "error: attribute ProductCode cannot be empty")
	}

	if p.Description != nil && *p.Description == "" {
		validationErrors = append(validationErrors, "error: attribute Description cannot be empty")
	}

	if p.Height != nil && *p.Height <= 0 {
		validationErrors = append(validationErrors, "error: attribute Height cannot be negative")
	}

	if p.Length != nil && *p.Length <= 0 {
		validationErrors = append(validationErrors, "error: attribute Length cannot be negative")
	}

	if p.Width != nil && *p.Width <= 0 {
		validationErrors = append(validationErrors, "error: attribute Width cannot be negative")
	}

	if p.NetWeight != nil && *p.NetWeight <= 0 {
		validationErrors = append(validationErrors, "error: attribute NetWeight cannot be negative")
	}

	if p.ExpirationRate != nil && *p.ExpirationRate < 0 {
		validationErrors = append(validationErrors, "error: attribute ExpirationRate cannot be negative")
	}

	if p.FreezingRate != nil && *p.FreezingRate < 0 {
		validationErrors = append(validationErrors, "error: attribute FreezingRate cannot be negative")
	}

	if p.ProductTypeID != nil && *p.ProductTypeID <= 0 {
		validationErrors = append(validationErrors, "error: attribute ProductTypeID must be positive")
	}

	if p.SellerID != nil && *p.SellerID < 0 {
		validationErrors = append(validationErrors, "error: attribute SellerID must be non-negative")
	}

	if p.SellerID != nil && *p.SellerID == 0 {
		validationErrors = append(validationErrors, "error: cannot nullify SellerID")
	}

	if len(validationErrors) > 0 {
		err = errors.New(fmt.Sprintf("validation errors: %v", validationErrors))
	}
	return
}
