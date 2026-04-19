package status

import (
	"context"
	"time"
)

func (r *statusRepository) CreateStatus(ctx context.Context, status string) (int64, error) {
	now := time.Now()
	query := `INSERT INTO statuses (status, created_at, updated_at) VALUES (?, ?, ?)`
	res, err := r.db.ExecContext(ctx, query, status, now, now)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}
