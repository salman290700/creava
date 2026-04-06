package locations

import (
	"context"
	"database/sql"
	"gotweet/internal/model"
	"time"
)

func (r *locationsRepository) GetDistrictsByRegencyId(ctx context.Context, regency_id string) ([]model.Districts, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	query := `SELECT id, code, name, regency_code from districts where regency_id = ?`
	rows, err := r.db.QueryContext(ctx, query, regency_id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	defer rows.Close()

	var districts []model.Districts

	for rows.Next() {
		var district model.Districts
		err := rows.Scan(&district.Id, &district.Code, &district.Name, &district.RegencyCode)
		if err != nil {
			return nil, err
		}
		districts = append(districts, district)
	}
	return districts, nil
}
