package user

import "context"

func (r *userRepository) CreateRefreshToken(ctx context.Context, id int64, refresh_token string) error {
	query := `INSERT INTO refresh_token (refresh_token, user_id) values (?, ?)`
	_, err := r.db.ExecContext(ctx, query, refresh_token, id)
	if err != nil {
		return err
	}
	return nil
}
