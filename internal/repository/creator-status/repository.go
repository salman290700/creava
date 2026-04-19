package creatorstatus

import (
	"context"
	"database/sql"
	"gotweet/internal/model"
)

type CreatorStatusRepository interface {
	// CreateCreatorStatus(ctx context.Context, creator_id int64)
	GetCreatorStatus(ctx context.Context, creator_status_id int64) (*model.CreatorStatus, error)
}

type creatorStatusRepository struct {
	db *sql.DB
}

func NewCreatorStatusRepository(db *sql.DB) CreatorStatusRepository {
	return &creatorStatusRepository{
		db: db,
	}
}
