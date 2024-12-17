package repository

import "frescos/models"

type SectionRepository struct {
	dbPath string
	db     map[int]models.Section
}
