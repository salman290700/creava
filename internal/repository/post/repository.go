package post

import (
	"context"
	"database/sql"
	"gotweet/internal/model"
)

type PostRepository interface {
	CreatePost(ctx context.Context, model *model.PostModel) (int64, error)
}

type postRepository struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) PostRepository {
	return &postRepository{
		db: db,
	}
}
