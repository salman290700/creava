package email

import (
	"context"
	"database/sql"
)

type EmailRepository interface {
	CreateEmail(ctx context.Context, email string, verified bool) (int64, error)
}

type emailRepository struct {
	db *sql.DB
}

func NewEmailRepository(db *sql.DB) EmailRepository {
	return &emailRepository{
		db: db,
	}
}
