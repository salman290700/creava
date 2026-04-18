package creatorimage

import (
	"context"
	"time"
)

func (r *creatorImageRepository) CreateCreatorImage(ctx context.Context, creator_id int64, image_url string, created_at time.Time) (int64, error) {
	query := `INSERT INTO creator_image (creator_id, image_url, created_at) VALUES (?, ?, ?)`
	res, err := r.db.ExecContext(ctx, query, creator_id, image_url, created_at)
	if err != nil {
		return 0, nil
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}
