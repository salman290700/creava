package contact

import (
	"context"
	"database/sql"
	"gotweet/internal/model"
)

type ContactRepository interface {
	CreateContact(ctx context.Context, phone_number string) (int64, error)
	GetContactById(ctx context.Context, contact_id int64) (*model.Contact, error)
}

type contactRepository struct {
	db *sql.DB
}

func NewContactRepository(db *sql.DB) ContactRepository {
	return &contactRepository{
		db: db,
	}
}
