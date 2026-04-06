package creatoremail

import (
	"context"
	"database/sql"
)

type CreatorEmailRepository interface {
	CreateCreatorEmail(ctx context.Context, creator_id int64, email string) (int64, error)
}

type creatorEmailRepository struct {
	db *sql.DB
}

func NewCreatorEmailRepository(db *sql.DB) CreatorEmailRepository {
	return &creatorEmailRepository{
		db: db,
	}
}
