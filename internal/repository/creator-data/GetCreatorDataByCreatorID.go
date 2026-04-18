package creatordata

import (
	"context"
	"database/sql"
	"gotweet/internal/model"
)

func (r *creatorDataRepository) GetCreatorDataByCreatorID(ctx context.Context, id int64) (*model.CreatorData, error) {
	query := `SELECT id, creator, email, phone_number, address, version, status, created_at, udpated_at WHERE creator = ?`
	res := r.db.QueryRowContext(ctx, query, id)
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
