package repository

import models "github.com/maxwelbm/alkemy-g6/internal/models/sections"

func (r *Sections) GetById(id int) (sec models.Section, err error) {
	sec, ok := r.db[id]
	if !ok {
		err = ErrSectionNotFound
	}

	return
}
