package creatorstatus

import (
	"context"
	"gotweet/internal/model"
)

func (r *creatorStatusRepository) GetCreatorStatus(ctx context.Context, creator_status_id int64) (*model.CreatorStatus, error) {
	query := `SELECT * from creator_statuses WHERE id = ?`
	res := r.db.QueryRowContext(ctx, query, creator_status_id)
	var response model.CreatorStatus
	res.Scan(&response.ID, &response.CreatorID, &response.Status, &response.CreatedAt)
	return &response, nil
}
