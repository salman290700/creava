package creatoremail

import (
	"context"
	"gotweet/internal/model"
)

func (r *creatorEmailRepository) GetCreatorEmail(ctx context.Context, creator_id int64) (*model.CreatorEmail, error) {
	query := `SELECT * FROM creator_emails where creator_id = ? ORDER BY created_at DESC LIMIT 1`
	row := r.db.QueryRowContext(ctx, query, creator_id)
	var creatorEmail model.CreatorEmail
	err := row.Scan(
		&creatorEmail.ID,
		&creatorEmail.Email,
		&creatorEmail.CreatorId,
		&creatorEmail.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &creatorEmail, nil
}
