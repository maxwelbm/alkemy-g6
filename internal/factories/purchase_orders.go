package factories

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/pkg/randstr"
)

type PurchaseOrdersFactory struct {
	db *sql.DB
}

func NewPurchaseOrdersFactory(db *sql.DB) *PurchaseOrdersFactory {
	return &PurchaseOrdersFactory{db: db}
}

func defaultPurchaseOrders() models.PurchaseOrders {
	return models.PurchaseOrders{
		OrderNumber:     randstr.Alphanumeric(8),
		OrderDate:       randstr.RandDate(),
		TrackingCode:    randstr.Alphanumeric(8),
		BuyerID:         1,
		ProductRecordID: 1,
	}
}

func (f *PurchaseOrdersFactory) Create(purchaseOrders models.PurchaseOrders) (record models.PurchaseOrders, err error) {
	populatePurchaseOrdersParams(&purchaseOrders)

	if err = f.checkBuyerExists(purchaseOrders.BuyerID); err != nil {
		return purchaseOrders, err
	}

	if err = f.checkProductRecordExists(purchaseOrders.ProductRecordID); err != nil {
		return purchaseOrders, err
	}

	query := `
		INSERT INTO purchase_orders 
			(
			%s
			order_number,
			order_date, 
			tracking_code, 
			buyer_id,
			product_record_id
			) 
		VALUES (%s?, ?, ?, ?, ?)
	`

	switch purchaseOrders.ID {
	case 0:
		query = fmt.Sprintf(query, "", "")
	default:
		query = fmt.Sprintf(query, "id,", strconv.Itoa(purchaseOrders.ID)+",")
	}

	_, err = f.db.Exec(query,
		purchaseOrders.OrderNumber,
		purchaseOrders.OrderDate,
		purchaseOrders.TrackingCode,
		purchaseOrders.BuyerID,
		purchaseOrders.ProductRecordID,
	)

	return purchaseOrders, err
}

// Refazer
func populatePurchaseOrdersParams(purchaseOrders *models.PurchaseOrders) {
	defaultPurchaseOrders := defaultPurchaseOrders()
	if purchaseOrders == nil {
		purchaseOrders = &defaultPurchaseOrders
	}

	if purchaseOrders.OrderNumber == "" {
		purchaseOrders.OrderNumber = defaultPurchaseOrders.OrderNumber
	}

	if purchaseOrders.OrderDate == "" {
		purchaseOrders.OrderDate = defaultPurchaseOrders.OrderDate
	}

	if purchaseOrders.TrackingCode == "" {
		purchaseOrders.TrackingCode = defaultPurchaseOrders.TrackingCode
	}

	if purchaseOrders.BuyerID == 0 {
		purchaseOrders.BuyerID = defaultPurchaseOrders.BuyerID
	}

	if purchaseOrders.ProductRecordID == 0 {
		purchaseOrders.ProductRecordID = defaultPurchaseOrders.ProductRecordID
	}
}

func (f *PurchaseOrdersFactory) checkBuyerExists(buyerID int) (err error) {
	var count int
	err = f.db.QueryRow(`SELECT COUNT(*) FROM buyers WHERE id = ?`, buyerID).Scan(&count)

	if err != nil {
		return
	}

	if count > 0 {
		return
	}

	err = f.createBuyer()

	return
}

func (f *PurchaseOrdersFactory) createBuyer() (err error) {
	buyerFactory := NewBuyerFactory(f.db)
	_, err = buyerFactory.Create(models.Buyer{})

	return
}

func (f *PurchaseOrdersFactory) checkProductRecordExists(productRecordID int) (err error) {
	var count int
	err = f.db.QueryRow(`SELECT COUNT(*) FROM product_records WHERE id = ?`, productRecordID).Scan(&count)

	if err != nil {
		return
	}

	if count > 0 {
		return
	}

	err = f.createProductRecord()

	return
}

func (f *PurchaseOrdersFactory) createProductRecord() (err error) {
	productRecordFactory := NewProductRecodsFactory(f.db)
	_, err = productRecordFactory.Create(models.ProductRecord{})

	return
}
