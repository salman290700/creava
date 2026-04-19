package image

import (
	"context"
	"database/sql"
	"gotweet/internal/model"
)

type ImageRepository interface {
	CreateImage(ctx context.Context, image string) (int64, error)
	GetImageByID(ctx context.Context, id int64) (*model.Image, error)
}

type imageRepository struct {
	db *sql.DB
}

func NewImageReopsitory(db *sql.DB) ImageRepository {
	return &imageRepository{
		db: db,
	}
}
