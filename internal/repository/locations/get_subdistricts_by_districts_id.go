package locations

import (
	"context"
	"gotweet/internal/model"
	"time"
)

func (r locationsRepository) GetSubDistrictsByDistrictId(ctx context.Context, district_id string) ([]model.SubDistricts, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	query := `SELECT id, code, name, district_code from subdistricts WHERE district_id = ?`
	rows, err := r.db.QueryContext(ctx, query, district_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var data []model.SubDistricts

	for rows.Next() {
		var item model.SubDistricts
		err := rows.Scan(&item.Id, &item.Code, &item.Name, &item.DistrictCode)
		if err != nil {
			return nil, err
		}
		data = append(data, item)
	}
	return data, nil
}
