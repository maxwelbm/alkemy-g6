package repository

import (
	"errors"

	models "github.com/maxwelbm/alkemy-g6/internal/models/sections"
)

var (
	ErrSectionNotFound       = errors.New("Section not found")
	ErrSectionDuplicatedCode = errors.New("Section code already exists")
)

type Sections struct {
	db     map[int]models.Section
	lastId int
}

func NewSections(db map[int]models.Section) *Sections {
	defaultDb := make(map[int]models.Section)
	if db != nil {
		defaultDb = db
	}

	lastId := 0
	for _, sec := range db {
		if lastId < sec.ID {
			lastId = sec.ID
		}
	}

	return &Sections{
		db:     defaultDb,
		lastId: lastId,
	}
}
