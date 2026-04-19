package creatoremail

import (
	"context"
	"database/sql"
	"gotweet/internal/model"
)

type CreatorEmailRepository interface {
	CreateCreatorEmail(ctx context.Context, creator_id int64, email string) (int64, error)
	GetCreatorEmail(ctx context.Context, creator_id int64) (*model.CreatorEmail, error)
}

type creatorEmailRepository struct {
	db *sql.DB
}

func NewCreatorEmailRepository(db *sql.DB) CreatorEmailRepository {
	return &creatorEmailRepository{
		db: db,
	}
}
