package creatoremail

import (
	"context"
)

func (r *creatorEmailRepository) CreateCreatorEmail(ctx context.Context, creator_id int64, email string) (int64, error) {
	query := `INSERT INTO creator_email (creator_id, email) (?, ?)`
	res, err := r.db.ExecContext(ctx, query, creator_id, email)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}
