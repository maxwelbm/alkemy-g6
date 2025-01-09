package sections_repository

func (r *SectionRepository) Delete(id int) (err error) {
	query := "DELETE FROM sections WHERE id = ?"
	_, err = r.DB.Exec(query, id)
	if err != nil {
		return
	}

	return nil
}
