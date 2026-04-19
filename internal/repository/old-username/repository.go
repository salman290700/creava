package oldusername

import (
	"context"
	"database/sql"
)

type OldUsernameRepository interface {
	CreateOldUsername(ctx context.Context, username string, creator_id int64) (int64, error)
}

type oldUsernameRepository struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) OldUsernameRepository {
	return &oldUsernameRepository{
		db: db,
	}
}
