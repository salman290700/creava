package user

import (
	"context"
	"errors"
)

func (r *userRepository) DeleteRefreshToken(ctx context.Context, userID int64) error {
	query := `delete from refresh_token where user_id = ?`
	result, err := r.db.ExecContext(ctx, query, userID)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("nothing to delete")
	}
	return nil
}
