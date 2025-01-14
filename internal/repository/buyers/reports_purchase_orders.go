package buyersrp

import (
	"github.com/maxwelbm/alkemy-g6/internal/models"
)

func (r *BuyerRepository) ReportPurchaseOrders(id int) (reports []models.BuyerPurchaseOrdersReport, err error) {
	query := `
		SELECT b.id, b.card_number_id, b.first_name, b.last_name, COUNT(p.id) purchase_orders_count
		FROM buyers b
		LEFT JOIN purchase_orders p ON b.id = p.buyer_id
		WHERE (? = 0 OR b.id = ?)
		GROUP BY b.id
	`
	rows, err := r.db.Query(query, id, id)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var report models.BuyerPurchaseOrdersReport
		err = rows.Scan(&report.ID, &report.CardNumberID, &report.FirstName, &report.LastName, &report.PurchaseOrdersCount)
		if err != nil {
			return
		}
		reports = append(reports, report)
	}
	if len(reports) == 0 {
		err = models.ErrBuyerNotFound
		return
	}
	if err = rows.Err(); err != nil {
		return
	}

	return
}
