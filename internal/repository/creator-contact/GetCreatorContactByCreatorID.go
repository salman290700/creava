package creatorContact

import (
	"context"
	"gotweet/internal/model"
)

func (r *creatorContactRepository) GetCreatorContactByCreatorID(ctx context.Context, creator_id int64) (*model.CreatorContact, error) {
	query := `SELECT id, contact_id, created_at from creator_contacts where creator_id = ?`
	res := r.db.QueryRowContext(ctx, query, creator_id)
	var response model.CreatorContact

	err := res.Scan(&response.ID, &response.ContactID, &response.CreatedAt)
	if err != nil {
		return nil, nil
	}
	return &response, nil
}
