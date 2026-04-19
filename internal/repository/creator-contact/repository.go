package creatorContact

import (
	"context"
	"database/sql"
	"gotweet/internal/model"
)

type CreatorContactRepository interface {
	CreateCreatorContact(ctx context.Context, creator_id int64, phone_number int64) (int64, error)
	GetCreatorContact(ctx context.Context, creator_id int64, phone_number string) (*model.CreatorContact, error)
	GetCreatorContactByCreatorID(ctx context.Context, creator_id int64) (*model.CreatorContact, error)
}

type creatorContactRepository struct {
	db *sql.DB
}

func NewCreatorContactRepository(db *sql.DB) CreatorContactRepository {
	return &creatorContactRepository{
		db: db,
	}
}
