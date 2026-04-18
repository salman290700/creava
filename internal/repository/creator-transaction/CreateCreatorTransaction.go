package creatortransaction

import (
	"context"
	"time"
)

func (r *creatorTransactionRepository) CreateCreatorTransaction(ctx context.Context, creator_id int64, balance_id int64, transaction_type string, amount float64, currency string, reference_id int64, description string, status string, version int64, created_at time.Time, updated_at time.Time) (int64, error) {
	query := `INSERT INTO creator_transaction (creator_id, balance_id, transaction_type, amount, amount, reference_id, description, status, version, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	row, err := r.db.ExecContext(ctx, query, creator_id, balance_id, transaction_type, amount, currency, reference_id, description, status, version, created_at, updated_at)
	if err != nil {
		return 0, err
	}

	id, err := row.LastInsertId()
	return id, nil
}
