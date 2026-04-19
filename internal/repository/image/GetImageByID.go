package image

import (
	"context"
	"database/sql"
	"gotweet/internal/model"
)

func (r *imageRepository) GetImageByID(ctx context.Context, id int64) (*model.Image, error) {
	query := `
		SELECT id, image, created_at
		FROM images
		WHERE id = ?
		LIMIT 1
	`

	row := r.db.QueryRowContext(ctx, query, id)

	var img model.Image
	err := row.Scan(
		&img.ID,
		&img.Image,
		&img.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // or return custom error
		}
		return nil, err
	}

	return &img, nil
}
