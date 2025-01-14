package localitiesrp

import "database/sql"

type LocalityRepository struct {
	db *sql.DB
}

func NewLocalityRepository(db *sql.DB) *LocalityRepository {
	return &LocalityRepository{db: db}
}
