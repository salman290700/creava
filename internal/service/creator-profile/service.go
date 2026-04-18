package creatorprofile

import (
	"context"
	"gotweet/internal/config"
	"gotweet/internal/model"
	"gotweet/internal/repository/address"
	"gotweet/internal/repository/creator"
	creatorRepository "gotweet/internal/repository/creator"
	creatoraddress "gotweet/internal/repository/creator-address"
	creatorbalance "gotweet/internal/repository/creator-balance"
	creatorContact "gotweet/internal/repository/creator-contact"
	creatordata "gotweet/internal/repository/creator-data"
	creatoremail "gotweet/internal/repository/creator-email"
	creatorimage "gotweet/internal/repository/creator-image"
	"gotweet/internal/repository/email"
	"gotweet/internal/repository/status"
)

type CreatorProfileService interface {
	GetCreatorProfile(ctx context.Context, cretor_id int64) (*model.CreatorProfiling, error)
}

type creatorProfileService struct {
	statusRepo         status.StatusRepository
	addressRepo        address.AddressRepository
	emailRepo          email.EmailRepository
	creatorContactRepo creatorContact.CreatorContactRepository
	creatorRepo        creatorRepository.CreatorRepository
	creatorAddressRepo creatoraddress.CreatorAddressRepository
	creatorEmailRepo   creatoremail.CreatorEmailRepository
	creatorDataRepo    creatordata.CreatorDataRepository
	creatorbalancerepo creatorbalance.CreatorBalanceRepository
	creatorImageRepo   creatorimage.CreatorImageRepository
	cfg                *config.Config
}

func NewCreatorProfileService(
	statusRepo status.StatusRepository,
	addressRepo address.AddressRepository,
	emailRepo email.EmailRepository,
	creatorContactRepo creatorContact.CreatorContactRepository,
	creatorRepo creator.CreatorRepository,
	creatorAddressRepo creatoraddress.CreatorAddressRepository,
	creatorEmailRepo creatoremail.CreatorEmailRepository,
	creatorDataRepo creatordata.CreatorDataRepository,
	creatorbalancerepo creatorbalance.CreatorBalanceRepository,
	creatorImageRepo creatorimage.CreatorImageRepository,
	cfg *config.Config,
) CreatorProfileService {
	return &creatorProfileService{
		statusRepo:         statusRepo,
		addressRepo:        addressRepo,
		emailRepo:          emailRepo,
		creatorContactRepo: creatorContactRepo,
		creatorRepo:        creatorRepo,
		creatorAddressRepo: creatorAddressRepo,
		creatorEmailRepo:   creatorEmailRepo,
		creatorDataRepo:    creatorDataRepo,
		creatorbalancerepo: creatorbalancerepo,
		creatorImageRepo:   creatorImageRepo,
		cfg:                cfg,
	}
}
