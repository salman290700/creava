package user

import (
	"context"
	"database/sql"
	"fmt"
	"gotweet/internal/model"
)

func (r *userRepository) GetUserByEmailOrUsername(ctx context.Context, email, username string) (*model.UserModel, error) {
	query := `SELECT id, email, username, password, created_at, updated_at
							FROM users 
							WHERE email = ? 
							OR username = ?`
	row := r.db.QueryRowContext(ctx, query, email, username)
	fmt.Println(row)
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
