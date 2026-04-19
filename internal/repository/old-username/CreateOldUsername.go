package oldusername

import "context"

func (r *oldUsernameRepository) CreateOldUsername(ctx context.Context, username string, creator_id int64) (int64, error) {
	query := `INSERT INTO old_username (username, creator_id) VALUES (?, ?)`

	res, err := r.db.ExecContext(ctx, query, username, creator_id)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}
