package user

import (
	"context"
	"database/sql"
	"gotweet/internal/model"
)

func (r *userRepository) GetUserbyId(ctx context.Context, id int64) (*model.UserModel, error) {
	query := `SELECT id, email, username, password, created_at, updated_at
							FROM users 
							WHERE email id = id`
	row := r.db.QueryRowContext(ctx, query, id)
	var result model.UserModel
	err := row.Scan(&result.ID, &result.Email, &result.Username, &result.Password, &result.CreatedAt, &result.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}
	return &result, nil
}
