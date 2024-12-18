package repository

import models "github.com/maxwelbm/alkemy-g6/internal/models/sections"

func (r *Sections) GetAll() (sec []models.Section, err error) {
	for _, value := range r.db {
		sec = append(sec, value)
	}

	return
}
