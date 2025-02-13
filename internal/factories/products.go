package factories

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/maxwelbm/alkemy-g6/internal/models"
)

type ProductFactory struct {
	db *sql.DB
}

func NewProductFactory(db *sql.DB) *ProductFactory {
	return &ProductFactory{db: db}
}

func defaultProduct() models.Product {
	return models.Product{
		ProductCode:    RandAlphanumeric(8),
		Description:    RandChars(16),
		Height:         10,
		Length:         10,
		Width:          10,
		NetWeight:      10,
		ExpirationRate: 10,
		FreezingRate:   10,
		RecomFreezTemp: 10,
		ProductTypeID:  1,
		SellerID:       1,
	}
}

func (f ProductFactory) Build(product models.Product) models.Product {
	populateProductParams(&product)

	return product
}

func (f *ProductFactory) Create(product models.Product) (record models.Product, err error) {
	populateProductParams(&product)

	if err = f.checkSellerExists(product.SellerID); err != nil {
		return record, err
	}

	query := `
		INSERT INTO products 
			(
			%s
			product_code,
			description,
			height,
			length,
			width,
			net_weight,
			expiration_rate,
			freezing_rate,
			recommended_freezing_temperature,
			product_type_id,
			seller_id
			)
		VALUES (%s?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	switch product.ID {
	case 0:
		query = fmt.Sprintf(query, "", "")
	default:
		query = fmt.Sprintf(query, "id,", strconv.Itoa(product.ID)+",")
	}

	result, err := f.db.Exec(query,
		product.ProductCode,
		product.Description,
		product.Height,
		product.Length,
		product.Width,
		product.NetWeight,
		product.ExpirationRate,
		product.FreezingRate,
		product.RecomFreezTemp,
		product.ProductTypeID,
		product.SellerID,
	)
	if err != nil {
		return record, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return record, err
	}

	product.ID = int(id)

	return product, err
}

func populateProductParams(product *models.Product) {
	defaultProduct := defaultProduct()
	if product == nil {
		product = &defaultProduct
	}

	if product.ProductCode == "" {
		product.ProductCode = defaultProduct.ProductCode
	}

	if product.Description == "" {
		product.Description = defaultProduct.Description
	}

	if product.Height == 0 {
		product.Height = defaultProduct.Height
	}

	if product.Length == 0 {
		product.Length = defaultProduct.Length
	}

	if product.Width == 0 {
		product.Width = defaultProduct.Width
	}

	if product.NetWeight == 0 {
		product.NetWeight = defaultProduct.NetWeight
	}

	if product.ExpirationRate == 0 {
		product.ExpirationRate = defaultProduct.ExpirationRate
	}

	if product.FreezingRate == 0 {
		product.FreezingRate = defaultProduct.FreezingRate
	}

	if product.RecomFreezTemp == 0 {
		product.RecomFreezTemp = defaultProduct.RecomFreezTemp
	}

	if product.ProductTypeID == 0 {
		product.ProductTypeID = defaultProduct.ProductTypeID
	}

	if product.SellerID == 0 {
		product.SellerID = defaultProduct.SellerID
	}
}

func (f *ProductFactory) checkSellerExists(sellerID int) (err error) {
	var count int
	err = f.db.QueryRow(`SELECT COUNT(*) FROM sellers WHERE id = ?`, sellerID).Scan(&count)

	if err != nil {
		return
	}

	if count == 0 {
		err = fmt.Errorf("seller with id %d does not exist", sellerID)
	}

	if err != nil {
		err = f.createSeller()
	}

	return
}

func (f *ProductFactory) createSeller() (err error) {
	sellerFactory := NewSellerFactory(f.db)
	_, err = sellerFactory.Create(models.Seller{})

	return
}
