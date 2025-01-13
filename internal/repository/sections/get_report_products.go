package sections_repository

import (
	"database/sql"

	"github.com/maxwelbm/alkemy-g6/internal/models"
)

func (r *SectionRepository) GetReportProducts(sectionId int) (reportProducts []models.ProductReport, err error) {
	var query string
	var rows *sql.Rows

	// Refazer a query
	query = ` 
		SELECT s.id, s.section_number, COUNT(pb.id) AS products_count
		FROM sections s 
		LEFT JOIN product_batches pb ON s.id = pb.section_id
		WHERE (? = 0 OR s.id = ?) 
		GROUP BY s.id
	`

	rows, err = r.DB.Query(query, sectionId, sectionId)

	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var productReport models.ProductReport
		err = rows.Scan(
			&productReport.SectionId,
			&productReport.SectionNumber,
			&productReport.ProductsCount,
		)
		if err != nil {
			return
		}
		reportProducts = append(reportProducts, productReport)
	}

	if len(reportProducts) == 0 {
		err = models.ErrSectionNotFound
		return
	}

	if err = rows.Err(); err != nil {
		return
	}

	return
}
