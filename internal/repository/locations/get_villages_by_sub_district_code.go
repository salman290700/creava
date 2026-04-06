package locations

import (
	"context"
	"gotweet/internal/model"
	"time"
)

func (r *locationsRepository) GetVillagesBySubdistrictId(ctx context.Context, subdistrict_id string) ([]model.Villages, error) {
	var (
		data     model.Villages
		response []model.Villages
	)
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	query := `SELECT id, code, name, sub_district_code from villages where sub_district_code = ?`
	rows, err := r.db.QueryContext(ctx, query, subdistrict_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&data.Id, &data.Code, &data.Name, &data.SubDistrictCode)
		response = append(response, data)
	}
	return response, nil
}
