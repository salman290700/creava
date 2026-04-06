package creatoraddress

import (
	"context"
	"gotweet/internal/model"
)

func (r *creatorAddressRepository) GetCreatorAddressbyCreatorIdAndAddressId(ctx context.Context, creator_id, address_id int64) (int64, error) {
	query := `SELECT address_id from creator_addresses where creator_id = ? AND address_id = ?`
	res := r.db.QueryRowContext(ctx, query, creator_id, address_id)
	var data model.ID
	res.Scan(&data.ID)

	return data.ID, nil
}
