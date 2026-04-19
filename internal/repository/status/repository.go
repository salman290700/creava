package status

import (
	"context"
	"database/sql"
	"gotweet/internal/model"
)

type StatusRepository interface {
	CreateStatus(ctx context.Context, status string) (int64, error)
	GetStatusByStatusID(ctx context.Context, id int64) (*model.Status, error)
}

type statusRepository struct {
	db *sql.DB
}

func NewStatusRepository(db *sql.DB) StatusRepository {
	return &statusRepository{
		db: db,
	}
}
