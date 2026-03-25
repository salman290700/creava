package user

import (
	"context"
	"database/sql"
	"gotweet/internal/model"
	"time"
)

func (r *userRepository) GetRefreshtoken(ctx context.Context, id int64, now time.Time) (*model.RefreshTokenModel, error) {
	query := `SELECT id, user_id, refresh_token, created_at, updated_at, expired_at FROM refresh_token WHERE user_id = ? AND expired_at >= ?`
	row := r.db.QueryRowContext(ctx, query, id, now)
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
