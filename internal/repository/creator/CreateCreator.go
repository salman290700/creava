package creator

import (
	"context"
)

func (r *creatorRepository) CreateCreator(ctx context.Context, name, hashed_password string) (int64, error) {
	// Create Creator
	query := `INSERT INTO creators (name, password) VALUES (?, ?)`
	res, err := r.db.ExecContext(ctx, query, name, hashed_password)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}
