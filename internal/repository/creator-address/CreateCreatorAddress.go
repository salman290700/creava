package creatoraddress

import "context"

func (r *creatorAddressRepository) CreateCreatorAddress(ctx context.Context, creator_id int64, address_id int64) (int64, error) {
	// Check address
	address_id, err := r.GetCreatorAddressbyCreatorIdAndAddressId(ctx, creator_id, address_id)
	if err != nil {
		return 0, err
	}

	if address_id != 0 {
		return 0, nil
	}
	query := `INSERT INTO creator_addresses (creator_id, address_id) VALUES (?, ?)`
	res, err := r.db.ExecContext(ctx, query, creator_id, address_id)
	if err != nil {
		return 0, nil
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, nil
	}
	return id, nil
}
