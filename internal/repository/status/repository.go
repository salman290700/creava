package status

import (
	"context"
	"database/sql"
)

type StatusRepository interface {
	CreateStatus(ctx context.Context, status string) (int64, error)
}

type statusRepository struct {
	db *sql.DB
}

func NewStatusRepository(db *sql.DB) StatusRepository {
	return &statusRepository{
		db: db,
	}
}
