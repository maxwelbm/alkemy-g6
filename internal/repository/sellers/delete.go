package sellers_repository

func (r *SellersDefault) Delete(id int) (err error) {
	// delete seller from database
	query := "DELETE FROM sellers WHERE id = ?"
	_, err = r.DB.Exec(query, id)
	if err != nil {
		return
	}
	return
}
