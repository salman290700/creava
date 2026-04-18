package creatorbalance

import (
	"context"
	"database/sql"
	"gotweet/internal/model"
	"time"
)

type CreatorBalanceRepository interface {
	CreateCreatorBalance(ctx context.Context, creator_id int64, balance float64, pending_balance float64, withdrawn_balance float64, currency string, version int64, created_at time.Time) (int64, error)
	GetCreatorBalanceByCreatorID(ctx context.Context, creator_id int64) (*model.CreatorBalance, error)
	GetCreatorBalance(ctx context.Context, id int64) (*model.CreatorBalance, error)
}

type creatorBalanceRepository struct {
	db *sql.DB
}

func NewCreatorBalanceRepository(db *sql.DB) CreatorBalanceRepository {
	return &creatorBalanceRepository{
		db: db,
	}
}
