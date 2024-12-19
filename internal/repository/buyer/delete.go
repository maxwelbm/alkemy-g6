package repository

func (r *BuyerRepository) Delete(id int) (err error) {
	delete(r.Buyers, id)
	return nil
}
