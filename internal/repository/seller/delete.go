package repository

func (r *SellerRepository) Delete(id int) (err error) {
	delete(r.Sellers, id)
	return nil
}
