package sectionsrp

import "github.com/maxwelbm/alkemy-g6/internal/models"

func (r *SectionRepository) Delete(id int) (err error) {
	var exists bool
	err = r.db.QueryRow("SELECT EXISTS(SELECT 1 FROM sections WHERE id = ?)", id).Scan(&exists)

	if err != nil {
		return
	}

	if !exists {
		err = models.ErrSectionNotFound
		return
	}

	query := "DELETE FROM sections WHERE id = ?"
	_, err = r.db.Exec(query, id)

	return
}
