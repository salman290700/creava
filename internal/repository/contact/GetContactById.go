package contact

import (
	"context"
	"gotweet/internal/model"
)

func (r *contactRepository) GetContactById(ctx context.Context, contact_id int64) (*model.Contact, error) {
	query := `SELECT id, phone_number, created_at from contacts where id = ?`
	res := r.db.QueryRowContext(ctx, query, contact_id)

	var model model.Contact
	err := res.Scan(&model.ID, &model.PhoneNumber, &model.CreatedAt)
	if err != nil {
		return nil, nil
	}
	return &model, nil
}
