package address

import "context"

func (r *addressRepository) CreateAddress(ctx context.Context, province_code, regency_code, district_code, sub_district_code, village_code, address string, postal_code int32) (int64, error) {
	query := `INSERT INTO addresses province_code, regency_code, district_code, sub_district_code, village_code, address, postal_code VALUES (?, ?, ?, ?, ?, ?, ?)`
	res, err := r.db.ExecContext(ctx, query, province_code, regency_code, district_code, sub_district_code, village_code, address, postal_code)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}
