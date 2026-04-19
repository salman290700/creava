package email

import (
	"context"
	"gotweet/internal/model"
)

func (r *emailRepository) GetEmailString(ctx context.Context, id int64) (*model.Email, error) {
	query := `SELECT * from emails WHERE id = ? LIMIT 1`
	row := r.db.QueryRowContext(ctx, query, id)
	var model model.Email
	if err := row.Scan(&model.ID, &model.Email, &model.Verified, &model.CreatedAt, &model.UpdatedAt); err != nil {
		return nil, err
	}
	return &model, nil
}
