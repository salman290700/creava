package creatorimage

import (
	"context"
	"database/sql"
	"gotweet/internal/model"
	"time"
)

type CreatorImageRepository interface {
	CreateCreatorImage(ctx context.Context, creator_id int64, image_url string, created_at time.Time) (int64, error)
	GetCreatorImageByCreatorID(ctx context.Context, creator_id int64) (*model.CreatorImage, error)
}

type creatorImageRepository struct {
	db *sql.DB
}

func NewCreatorEmailRepository(db *sql.DB) CreatorImageRepository {
	return &creatorImageRepository{
		db: db,
	}
}
