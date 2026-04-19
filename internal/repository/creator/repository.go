package creator

import (
	"context"
	"database/sql"
	"gotweet/internal/model"
)

type CreatorRepository interface {
	CreateCreator(ctx context.Context, name, hashed_password string) (int64, error)
	GetCreatorByID(ctx context.Context, creator_id int64) (*model.Creator, error)
}

type creatorRepository struct {
	db *sql.DB
}

func NewCreatorRepository(db *sql.DB) CreatorRepository {
	return &creatorRepository{
		db: db,
	}
}
