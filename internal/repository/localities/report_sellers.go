package localitiesrp

import (
	"github.com/maxwelbm/alkemy-g6/internal/models"
)

func (r *LocalityRepository) ReportSellers(id int) (reports []models.LocalitySellersReport, err error) {
	// selects locality info and seller count
	query := `
		SELECT l.id, l.locality_name, COUNT(s.id) AS sellers_count
		FROM localities AS l
		LEFT JOIN sellers AS s ON l.id = s.locality_id
		WHERE (? = 0 OR l.id = ?)
		GROUP BY l.id
	`
	rows, err := r.db.Query(query, id, id)

	if err != nil {
		return reports, err
	}

	defer rows.Close()
	// scans row into report
	for rows.Next() {
		var report models.LocalitySellersReport
		if err = rows.Scan(&report.ID, &report.LocalityName, &report.SellersCount); err != nil {
			return reports, err
		}

		reports = append(reports, report)
	}

	if len(reports) == 0 {
		err = models.ErrLocalityNotFound
		return reports, err
	}

	if err = rows.Err(); err != nil {
		return reports, err
	}

	return reports, nil
}
