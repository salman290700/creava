package creatoraddress

import (
	"context"
	"database/sql"
)

type CreatorAddressRepository interface {
	CreateCreatorAddress(ctx context.Context, creator_id int64, address, province, regency, district, sub_district, village string) (int64, error)
}

type creatorAddressRepository struct {
	db *sql.DB
}

func NewCreatorAddressRepository(db *sql.DB) CreatorAddressRepository {
	return &creatorAddressRepository{
		db: db,
	}
}
