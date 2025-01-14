package sectionsrp

import "database/sql"

type SectionRepository struct {
	DB *sql.DB
}

func NewSectionsRepository(DB *sql.DB) *SectionRepository {
	repo := &SectionRepository{
		DB: DB,
	}
	return repo
}
