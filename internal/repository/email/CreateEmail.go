package email

import "context"

func (r *emailRepository) CreateEmail(ctx context.Context, email string, verified bool) (int64, error) {
	query := `INSERT INTO emails (email, verified) VALUES (?, ?)`
	res, err := r.db.ExecContext(ctx, query, email, verified)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, nil
	}
	return id, nil
}
