package creatorbalance

import (
	"context"
	"gotweet/internal/model"
)

func (r *creatorBalanceRepository) GetCreatorBalanceByCreatorID(ctx context.Context, creator_id int64) (*model.CreatorBalance, error) {
	query := `SELECT id, creator_id, balance, pending_balance, withdrawn_balance, currency, version, created_at, updated_at WHERE CREATOR_ID = ?`
	row := r.db.QueryRowContext(ctx, query, creator_id)

	var response model.CreatorBalance
	row.Scan(&response.ID, &response.CreatorID, &response.Balance, &response.PendingBalance, &response.WithdrawnBalance, &response.Currency, &response.Version, &response.CreatedAt, &response.UpdatedAt)
	return &response, nil
}
