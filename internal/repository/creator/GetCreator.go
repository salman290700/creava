package creator

import (
	"context"
	"database/sql"
	"gotweet/internal/model"
)

func (r *creatorRepository) GetCreatorByID(ctx context.Context, id int64) (*model.Creator, error) {
	query := `
		SELECT id, name, username, password, version, created_at
		FROM creators
		WHERE id = ?
		LIMIT 1
	`

	row := r.db.QueryRowContext(ctx, query, id)

	var creator model.Creator
	err := row.Scan(
		&creator.ID,
		&creator.Name,
		&creator.Username,
		&creator.Password,
		&creator.Version,
		&creator.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // or return custom error like ErrCreatorNotFound
		}
		return nil, err
	}

	return &creator, nil
}
