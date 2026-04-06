package locations

import (
	"context"
	"fmt"
	"gotweet/internal/model"
	"time"
)

func (r *locationsRepository) GetRegenciesByProvinciesId(ctx context.Context, provincy_id string) ([]model.Regencies, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	fmt.Println(provincy_id)
	fmt.Println("In Repo")
	query := `SELECT id, code, name, province_code FROM regencies where province_code = ?`
	rows, err := r.db.QueryContext(ctx, query, provincy_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var regencies []model.Regencies
	for rows.Next() {
		var regency model.Regencies

		err := rows.Scan(&regency.Id, &regency.Code, &regency.Name, &regency.ProvincesCode)
		if err != nil {
			return nil, err
		}
		regencies = append(regencies, regency)
	}

	return regencies, nil
}
