package creatoraddress

import "context"

func (r *creatorAddressRepository) CreateCreatorAddress(ctx context.Context, creator_id int64, address, province, regency, district, sub_district, village string) (int64, error) {
	query := `INSERT INTO creator_address )creator_id, address, province, regency, district, sub_district, village) VALUES (?, ?, ?, ?, ?, ?, ?)`
	res, err := r.db.ExecContext(ctx, query, creator_id, address, province, regency, district, sub_district, village)
	if err != nil {
		return 0, nil
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, nil
	}
	return id, nil
}
