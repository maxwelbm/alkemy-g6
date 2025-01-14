package sectionsrp

import "database/sql"

type SectionRepository struct {
	db *sql.DB
}

func NewSectionsRepository(db *sql.DB) *SectionRepository {
	return &SectionRepository{
		db: db,
	}
}
