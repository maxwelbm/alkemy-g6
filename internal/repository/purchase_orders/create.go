package purchaseordersrp

import (
	"github.com/maxwelbm/alkemy-g6/internal/models"
)

func (r *PurchaseOrdersRepository) Create(purchaseOrdersDTO models.PurchaseOrdersDTO) (po models.PurchaseOrders, err error) {
	query := "INSERT INTO purchase_orders(`order_number`,`order_date`,`tracking_code`,`buyer_id`,`product_record_id`) VALUES(?, ?, ?, ?, ?)"

	result, err := r.DB.Exec(query, purchaseOrdersDTO.OrderNumber, purchaseOrdersDTO.OrderDate, purchaseOrdersDTO.TrackingCode, purchaseOrdersDTO.BuyerID, purchaseOrdersDTO.ProductRecordID)
	if err != nil {
		return
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return
	}

	query = "SELECT `id`,`order_number`, DATE_FORMAT(`order_date`, '%Y-%m-%d') AS order_date ,`tracking_code`,`buyer_id`,`product_record_id` FROM purchase_orders WHERE `id`=?"
	err = r.DB.QueryRow(query, lastInsertID).Scan(&po.ID, &po.OrderNumber, &po.OrderDate, &po.TrackingCode, &po.BuyerID, &po.ProductRecordID)

	return
}
