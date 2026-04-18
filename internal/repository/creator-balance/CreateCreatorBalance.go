package creatorbalance

import (
	"context"
	"time"
)

func (r *creatorBalanceRepository) CreateCreatorBalance(ctx context.Context, creator_id int64, balance float64, pending_balance float64, withdrawn_balance float64, currency string, version int64, created_at time.Time) (int64, error) {
	query := `INSERT INTO creator_balance (creator_id, balance, pending_balance, withdrawn_balance, currency, version, created_at) VALUES (?, ?, ?, ?, ?, ?, ?)`
	res, err := r.db.ExecContext(ctx, query, creator_id, balance, pending_balance, withdrawn_balance, currency, version, created_at)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	return id, nil
}
