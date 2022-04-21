package postgres

import "github.com/jmoiron/sqlx"

type mediaRepo struct {
	db *sqlx.DB
}

// New Media Repo ...
func NewMediaRepo(db *sqlx.DB) *mediaRepo {
	return &mediaRepo{db: db}
}
