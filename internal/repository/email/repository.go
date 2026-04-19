package email

import (
	"context"
	"database/sql"
	"gotweet/internal/model"
)

type EmailRepository interface {
	CreateEmail(ctx context.Context, email string, verified bool) (int64, error)
	GetEmailString(ctx context.Context, creator_id int64) (*model.Email, error)
}

type emailRepository struct {
	db *sql.DB
}

func NewEmailRepository(db *sql.DB) EmailRepository {
	return &emailRepository{
		db: db,
	}
}
