package creatorContact

import (
	"context"
)

func (r *creatorContactRepository) CreateCreatorContact(ctx context.Context, creator_id int64, phone_number string) (int64, error) {
	query := `INSERT INTO creator_contact (creator_id, phone_number) VALUES (?, ?)`
	res, err := r.db.ExecContext(ctx, query, creator_id, phone_number)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}
