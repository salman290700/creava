package creatorbalance

import (
	"context"
	"gotweet/internal/model"
)

func (r *creatorBalanceRepository) GetCreatorBalance(ctx context.Context, id int64) (*model.CreatorBalance, error) {
	query := `SELECT id, creator_id, balance, pending_balance, withdrawn_balance, currency, version, created_at, updated_at WHERE id = ?`
	row := r.db.QueryRowContext(ctx, query, id)

	var response model.CreatorBalance
	row.Scan(&response.ID, &response.CreatorID, &response.Balance, &response.PendingBalance, &response.WithdrawnBalance, &response.Currency, &response.Version, &response.CreatedAt, &response.UpdatedAt)
	return &response, nil
}
