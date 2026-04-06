package creator

import (
	"context"
	"database/sql"
	"gotweet/internal/model"
)

type CreatorRepository interface{
	CreateCreator(ctx context.Context, data *model.Creator) (int64, error)
}

type creatorRepository struct {
	db *sql.DB
}

func NewCreatorRepository(db *sql.DB) CreatorRepository {
	return &creatorRepository{
		db: db,
	}
}
