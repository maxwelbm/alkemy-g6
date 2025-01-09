package sections_repository

import "github.com/maxwelbm/alkemy-g6/internal/models"

func (r *SectionRepository) Delete(id int) (err error) {
	var exists bool
	err = r.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM sellers WHERE id = ?)", id).Scan(&exists)
	if err != nil {
		return
	}
	if !exists {
		err = models.ErrSectionNotFound
		return
	}

	query := "DELETE FROM sections WHERE id = ?"
	_, err = r.DB.Exec(query, id)
	if err != nil {
		return
	}

	return nil
}
