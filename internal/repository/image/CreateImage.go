package image

import (
	"context"
	"time"
)

func (r *imageRepository) CreateImage(ctx context.Context, image string) (int64, error) {
	query := `INSERT INTO image (image, created_at) VALUES (?, ?)`
	now := time.Now()
	res, err := r.db.ExecContext(ctx, query, image, now)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}
