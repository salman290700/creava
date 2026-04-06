package locations

import (
	"context"
	"database/sql"
	"gotweet/internal/model"
	"time"
)

func (r *locationsRepository) GetProvinces(ctx context.Context) ([]model.Provincies, error) {
	// Timeout for 5s
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `SELECT id, code, name FROM provincies`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	defer rows.Close()

	var response []model.Provincies
	for rows.Next() {
		var data model.Provincies
		err = rows.Scan(&data.ID, &data.Code, &data.Name)
		if err != nil {
			return nil, err
		}
		response = append(response, data)
	}

	return response, nil
}
