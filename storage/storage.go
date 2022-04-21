package storage

import (
	"github.com/jmoiron/sqlx"

	"github.com/asadbekGo/telegram-message-service/storage/postgres"
	"github.com/asadbekGo/telegram-message-service/storage/repo"
)

// IStorage ...
type IStorage interface {
	Message() repo.MediaStorageI
}

type storage struct {
	db        *sqlx.DB
	mediaRepo repo.MediaStorageI
}

// NewStorage ...
func NewStorage(db *sqlx.DB) *storage {
	return &storage{
		db:        db,
		mediaRepo: postgres.NewMediaRepo(db),
	}
}

func (s storage) Message() repo.MediaStorageI {
	return s.mediaRepo
}
