package creatordata

import (
	"context"
)

func (r *creatorDataRepository) CreateCreatorData(ctx context.Context, creator_id int64, name string, email_id int64) (int64, int64, int64, error) {
	query := `INSERT INTO creatorData (creator_id, name, email) VALUES (?, ?, ?)`
	res, err := r.db.ExecContext(ctx, query, creator_id, name, email_id)
	if err != nil {
		return 0, 0, 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, 0, 0, err
	}
	return id, creator_id, email_id, nil
}
