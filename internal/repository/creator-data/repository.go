package creatordata

import (
	"context"
	"database/sql"
	"gotweet/internal/model"
	"time"
)

type CreatorDataRepository interface {
	// creator_data_id, creator_id, email, error
	CreateCreatorData(ctx context.Context, creator_id int64, name string, email_id int64) (int64, int64, int64, error)
	// Get Creator Data
	GetCreatorData(ctx context.Context, email string) (*model.CreatorData, error)
	StoreRefreshToken(ctx context.Context, refresh_token *model.RefreshTokenCreatorModel, userID int64) error
	GetRefreshtoken(ctx context.Context, id int64, now time.Time) (*model.RefreshTokenCreatorModel, error)
	GetCreatorDataByCreatorID(ctx context.Context, creator_id int64) (*model.CreatorData, error)
}

type creatorDataRepository struct {
	db *sql.DB
}

func NewCreatorDataRepository(db *sql.DB) CreatorDataRepository {
	return &creatorDataRepository{
		db: db,
	}
}
