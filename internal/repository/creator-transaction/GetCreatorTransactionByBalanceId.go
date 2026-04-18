package creatortransaction

import (
	"context"
	"gotweet/internal/model"
)

func (r *creatorTransactionRepository) GetCreatorTransactionByBalanceId(ctx context.Context, balance_id int64) (*model.CreatorTransaction, error) {
	query := `SELECT * FROM creator_transaction WHERE balance_id = ?`
	row := r.db.QueryRowContext(ctx, query, balance_id)
	var response model.CreatorTransaction
	if err := row.Scan(&response.ID, &response.Creator_id, &response.Balance_id, &response.Transaction_type, &response.Amount, &response.Currency, &response.Reference_id, &response.Description, &response.Status, &response.Version, &response.Created_at, &response.Updated_at); err != nil {
		return nil, err
	}

	return &response, nil
}
