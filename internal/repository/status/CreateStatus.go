package status

import "context"

func (r *statusRepository) CreateStatus(ctx context.Context, status string) (int64, error) {
	query := `INSERT INTO statuses (status) VALUES (?)`
	res, err := r.db.ExecContext(ctx, query, status)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}
