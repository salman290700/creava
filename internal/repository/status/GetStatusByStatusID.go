package status

import (
	"context"
	"database/sql"
	"gotweet/internal/model"
)

func (r *statusRepository) GetStatusByStatusID(ctx context.Context, id int64) (*model.Status, error) {
	query := `
		SELECT id, status, created_at, updated_at
		FROM status
		WHERE id = ?
		LIMIT 1
	`

	row := r.db.QueryRowContext(ctx, query, id)

	var status model.Status
	err := row.Scan(
		&status.ID,
		&status.Status,
		&status.CreatedAt,
		&status.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &status, nil
}
