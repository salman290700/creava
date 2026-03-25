package user

import (
	"context"
	"gotweet/internal/model"
)

func (r *userRepository) StoreRefreshToken(ctx context.Context, refresh_token *model.RefreshTokenModel, userID int64) error {
	query := `INSERT INTO refresh_token (refresh_token, user_id, created_at, updated_at, expired_at) VALUES (?, ?, ?, ?, ?)`
	_, err := r.db.ExecContext(ctx, query, refresh_token.RefreshToken, refresh_token.UserID, refresh_token.CreatedAt, refresh_token.UpdatedAt, refresh_token.ExpiredAt)
	if err != nil {
		return err
	}
	return nil
}
