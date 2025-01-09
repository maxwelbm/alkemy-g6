package repository

import (
	"database/sql"

	"github.com/maxwelbm/alkemy-g6/internal/models"
)

func (p *Products) GetReportRecords(id int) (list []models.ProductReportRecords, err error) {
	var query string
	var rows *sql.Rows

	if id == 0 {
		query = "SELECT p.`id`, p.`description`, COALESCE(COUNT(pr.`id`), 0) AS records_count FROM products p LEFT JOIN product_records pr ON p.`id` = pr.`product_id` GROUP BY p.`id`, p.`description`"
		rows, err = p.DB.Query(query)
	} else {
		query = "SELECT p.`id`, p.`description`, COALESCE(COUNT(pr.`id`), 0) AS records_count FROM products p LEFT JOIN product_records pr ON p.`id` = pr.`product_id` WHERE p.`id` = ? GROUP BY p.`id`, p.`description`"
		rows, err = p.DB.Query(query, id)
	}

	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var productRecord models.ProductReportRecords
		err = rows.Scan(
			&productRecord.ProductId,
			&productRecord.Description,
			&productRecord.RecordsCount,
		)
		if err != nil {
			return
		}
		list = append(list, productRecord)
	}

	if err = rows.Err(); err != nil {
		return
	}

	return
}
