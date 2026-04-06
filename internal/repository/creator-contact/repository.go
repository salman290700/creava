package creatorContact

import (
	"context"
	"database/sql"
)

type CreatorContactRepository interface {
	CreateCreatorContact(ctx context.Context, creator_id int64, phone_number string) (int64, error)
	GetCreatorContact(ctx context.Context, creator_id int64, phone_number string) (int64, error)
}

type creatorContactRepository struct {
	db *sql.DB
}

func NewCreatorContactRepository(db *sql.DB) CreatorContactRepository {
	return &creatorContactRepository{
		db: db,
	}
}
