package repository

import models "github.com/maxwelbm/alkemy-g6/internal/models/sections"

type Sections struct {
	db map[int]models.Section
}

func NewSections(db map[int]models.Section) *Sections {
	defaultDb := make(map[int]models.Section)
	if db != nil {
		defaultDb = db
	}
	return &Sections{db: defaultDb}
}
