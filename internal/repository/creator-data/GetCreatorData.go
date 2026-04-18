package creatordata

import (
	"context"
	"database/sql"
	"gotweet/internal/model"
)

func (r *creatorDataRepository) GetCreatorData(ctx context.Context, email string) (*model.CreatorData, error) {
	email_query := `select id, email ,verified, created_at, updated_at from emails where email = ? LIMIT 1`
	res := r.db.QueryRowContext(ctx, email_query, email)
	var data_email model.Email
	res.Scan(&data_email.ID, &data_email.Email, &data_email.Verified, &data_email.CreatedAt, &data_email.UpdatedAt)
	if err := res.Err(); err != nil {
		return nil, nil
	}

	query := `SELECT id, creator, email, phone_number, address, version, status, created_at, udpated_at WHERE email = ?`
	res = r.db.QueryRowContext(ctx, query, data_email.ID)
	var response model.CreatorData
	err := res.Scan(&response.ID, &response.Creator, &response.Email, &response.Address, &response.Version, &response.Status, &response.Status, &response.CreatedAt, &response.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &response, nil
}
