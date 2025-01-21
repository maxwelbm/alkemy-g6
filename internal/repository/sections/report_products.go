package sectionsrp

import (
	"database/sql"

	"github.com/maxwelbm/alkemy-g6/internal/models"
)

func (r *SectionRepository) ReportProducts(sectionID int) (reportProducts []models.ProductReport, err error) {
	var query string

	var rows *sql.Rows

	query = ` 
		SELECT s.id, s.section_number, COUNT(pb.id) AS products_count
		FROM sections s 
		LEFT JOIN product_batches pb ON s.id = pb.section_id
		WHERE (? = 0 OR s.id = ?) 
		GROUP BY s.id
	`

	rows, err = r.db.Query(query, sectionID, sectionID)

	if err != nil {
		return reportProducts, err
	}

	defer rows.Close()

	for rows.Next() {
		var productReport models.ProductReport
		err = rows.Scan(
			&productReport.SectionID,
			&productReport.SectionNumber,
			&productReport.ProductsCount,
		)

		if err != nil {
			return reportProducts, err
		}

		reportProducts = append(reportProducts, productReport)
	}

	if len(reportProducts) == 0 {
		err = models.ErrSectionNotFound
		return reportProducts, err
	}

	if err = rows.Err(); err != nil {
		return reportProducts, err
	}

	return reportProducts, nil
}
