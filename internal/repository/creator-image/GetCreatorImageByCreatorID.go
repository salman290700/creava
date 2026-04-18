package creatorimage

import (
	"context"
	"database/sql"
	"gotweet/internal/model"
)

func (r *creatorImageRepository) GetCreatorImageByCreatorID(ctx context.Context, creatorID int64) (*model.CreatorImage, error) {
	query := `
		SELECT id, creator_id, image_url, created_at
		FROM creator_images
		WHERE creator_id = ?
		ORDER BY created_at DESC
		LIMIT 1
	`

	row := r.db.QueryRowContext(ctx, query, creatorID)

	var creatorImage model.CreatorImage
	err := row.Scan(
		&creatorImage.ID,
		&creatorImage.CreatorID,
		&creatorImage.ImageUrl,
		&creatorImage.CreatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // or return custom error if you prefer
		}
		return nil, err
	}
	return &creatorImage, nil
}
