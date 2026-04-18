package creatortransaction

import (
	"context"
	"database/sql"
	"gotweet/internal/model"
	"time"
)

type CreatorTransactionRepository interface {
	CreateCreatorTransaction(ctx context.Context, creator_id int64, balance_id int64, transaction_type string, amount float64, currency string, reference_id int64, description string, status string, version int64, created_at time.Time, updated_at time.Time) (int64, error)
	GetCreatorTransaction(ctx context.Context, id int64) (*model.CreatorTransaction, error)
	GetCreatorTransactionByBalanceId(ctx context.Context, balance_id int64) (*model.CreatorTransaction, error)
	GetCreatorTransactionByCreatorId(ctx context.Context, creator_id int64) (*model.CreatorTransaction, error)
}

type creatorTransactionRepository struct {
	db *sql.DB
}

func NewCreatorTransactionRepository(db *sql.DB) CreatorTransactionRepository {
	return &creatorTransactionRepository{
		db: db,
	}
}
