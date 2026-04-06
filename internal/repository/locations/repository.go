package locations

import (
	"context"
	"database/sql"
	"gotweet/internal/model"
)

type LocationsRepository interface {
	// Functions
	GetProvinces(ctx context.Context) ([]model.Provincies, error)
	GetRegenciesByProvinciesId(ctx context.Context, provincy_id string) ([]model.Regencies, error)
	GetDistrictsByRegencyId(ctx context.Context, regency_id string) ([]model.Districts, error)
	GetSubDistrictsByDistrictId(ctx context.Context, district_id string) ([]model.SubDistricts, error)
	GetVillagesBySubdistrictId(ctx context.Context, subdistrict_id string) ([]model.Villages, error)
}
type locationsRepository struct {
	db *sql.DB
}

func NewLocationsRepository(db *sql.DB) LocationsRepository {
	return &locationsRepository{
		db: db,
	}
}
