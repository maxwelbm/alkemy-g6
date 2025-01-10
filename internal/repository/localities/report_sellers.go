package localities_repository

import (
	"database/sql"
	"errors"

	"github.com/maxwelbm/alkemy-g6/internal/models"
)

func (r *LocalityRepository) ReportSellers(id int) (report models.LocalitySellersReport, err error) {
	// selects locality info and seller count
	query := `
		SELECT l.id, l.locality_name, COUNT(s.id) AS sellers_count
		FROM localities AS l
		LEFT JOIN sellers AS s ON l.id = s.locality_id
		WHERE l.id = ?
		GROUP BY l.id;
	`
	row := r.db.QueryRow(query, id)
	// scans row into report
	err = row.Scan(&report.ID, &report.LocalityName, &report.SellersCount)

	// check for errors
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = models.ErrLocalityNotFound
			return
		}
		return
	}
	return
}
