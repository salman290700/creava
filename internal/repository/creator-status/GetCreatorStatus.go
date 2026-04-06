package creatorstatus

import (
	"context"
	"gotweet/internal/model"
)

func (r *creatorStatusRepository) GetCreatorStatus(ctx context.Context, creator_id int64) (*model.CreatorStatus, error) {
	query := `SELECT creator_status.id, creator_status.creator_id, status.status, creator_status.created_at FROM creator_status 
	INNER JOIN status ON status.id = creators.id
	where crestor_status.creator_id = ?`
	res := r.db.QueryRowContext(ctx, query, creator_id)
	var response model.CreatorStatus
	res.Scan(&response.ID, &response.CreatorID, &response.Status, &response.CreatedAt)

	return &response, nil
}
