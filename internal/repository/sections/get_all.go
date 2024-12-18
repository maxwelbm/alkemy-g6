package section_repository

import models "github.com/maxwelbm/alkemy-g6/internal/models/sections"

func (r *Sections) GetAll() (sec map[int]models.Section, err error) {
	sec = make(map[int]models.Section)

	// copy db
	for key, value := range r.db {
		sec[key] = value
	}

	return
}
