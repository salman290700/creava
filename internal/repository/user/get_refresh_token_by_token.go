package user

import (
	"context"
	"database/sql"
	"gotweet/internal/model"
)

func (r *userRepository) GetRefreshtokenByToken(ctx context.Context, token string) (*model.RefreshTokenModel, error) {
	query := `SELECT id, user_id, refresh_token, created_at, updated_at, expired_at WHERE refresh_token = ?`
	row := r.db.QueryRowContext(ctx, query, token)
	var result model.RefreshTokenModel
	err := row.Scan(&result.ID, &result.UserID, &result.RefreshToken, &result.CreatedAt, &result.UpdatedAt, &result.ExpiredAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &result, nil
}
