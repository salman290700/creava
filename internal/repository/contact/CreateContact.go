package contact

import "context"

func (r *contactRepository) CreateContact(ctx context.Context, phone_number string) (int64, error) {
	query := `INSERT INTO contacts (phone_number) VALUES (?)`
	res, err := r.db.ExecContext(ctx, query, phone_number)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}
