package localitiesrp

import (
	"github.com/maxwelbm/alkemy-g6/internal/models"
)

func (r *LocalityRepository) ReportCarries(id int) (reports []models.LocalityCarriesReport, err error) {
	// selects locality info and carries count
	query := `
		SELECT l.id, l.locality_name, COUNT(c.id) AS carries_count
		FROM localities AS l
		LEFT JOIN carries AS c ON l.id = c.locality_id
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
		var report models.LocalityCarriesReport
		if err = rows.Scan(&report.ID, &report.LocalityName, &report.CarriesCount); err != nil {
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
