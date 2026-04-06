package creator

import (
	"context"
	"database/sql"
)

type CreatorRepository interface {
	CreateCreator(ctx context.Context, name, hashed_password string) (int64, error)
}

type creatorRepository struct {
	db *sql.DB
}

func NewCreatorRepository(db *sql.DB) CreatorRepository {
	return &creatorRepository{
		db: db,
	}
}
