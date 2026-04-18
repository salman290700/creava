package creatordonation

import (
	"context"
	"database/sql"
	"gotweet/internal/model"
)

type CreatorDonationRepository interface {
	CreateCreatorDonation(ctx context.Context, data *model.CreatorDonation) (int64, error)
	GetDonationById(ctx context.Context, id int64) (*model.CreatorDonation, error)
}

type creatorDonationRepository struct {
	db *sql.DB
}

func NewCreatorBalanceRepository(db *sql.DB) CreatorDonationRepository {
	return &creatorDonationRepository{
		db: db,
	}
}
