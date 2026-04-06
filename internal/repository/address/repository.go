package address

import (
	"context"
	"database/sql"
)

type AddressRepository interface {
	CreateAddress(ctx context.Context, province_code, regency_code, district_code, sub_district_code, village_code, address string, postal_code int32) (int64, error)
}

type addressRepository struct {
	db *sql.DB
}

func NewAddressRepository(db *sql.DB) AddressRepository {
	return &addressRepository{
		db: db,
	}
}
