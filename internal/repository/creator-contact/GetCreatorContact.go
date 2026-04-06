package creatorContact

import (
	"context"
	"database/sql"
	"gotweet/internal/model"
)

func (r *creatorContactRepository) GetCreatorContact(ctx context.Context, creator_id int64, phone_number string) (int64, error) {
	query := `SELECT id from creator_contacts where creator_id = ? AND phone_number = ?`
	res := r.db.QueryRowContext(ctx, query, creator_id, phone_number)
	var response model.ID
	if err := res.Err(); err != nil {
		if err == sql.ErrNoRows {
			return 0, nil
		}
		return 0, err
	}
	res.Scan(&response.ID)
	return response.ID, nil
}
