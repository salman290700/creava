package creator

import (
	"context"
	"gotweet/internal/config"
	"gotweet/internal/dto"
	"gotweet/internal/repository/address"
	"gotweet/internal/repository/creator"
	creatoraddress "gotweet/internal/repository/creator-address"
	creatorContact "gotweet/internal/repository/creator-contact"
	creatordata "gotweet/internal/repository/creator-data"
	creatoremail "gotweet/internal/repository/creator-email"
	"gotweet/internal/repository/email"
	"gotweet/internal/repository/status"
)

type CreatorService interface {
	LoginWithGoogle(ctx context.Context, token *dto.UserResGoogle) (string, string, int64, error)
}

type creatorService struct {
	statusRepo         status.StatusRepository
	addressRepo        address.AddressRepository
	emailRepo          email.EmailRepository
	creatorContactRepo creatorContact.CreatorContactRepository
	creatorRepo        creator.CreatorRepository
	creatorAddressRepo creatoraddress.CreatorAddressRepository
	creatorEmailRepo   creatoremail.CreatorEmailRepository
	creatorDataRepo    creatordata.CreatorDataRepository
	cfg                *config.Config
}

func NewCreatorService(
	statusRepo status.StatusRepository,
	addressRepo address.AddressRepository,
	emailRepo email.EmailRepository,
	creatorContactRepo creatorContact.CreatorContactRepository,
	creatorRepo creator.CreatorRepository,
	creatorAddressRepo creatoraddress.CreatorAddressRepository,
	creatorEmailRepo creatoremail.CreatorEmailRepository,
	creatorDataRepo creatordata.CreatorDataRepository,
	cfg *config.Config,
) CreatorService {
	return &creatorService{
		statusRepo:         statusRepo,
		addressRepo:        addressRepo,
		emailRepo:          emailRepo,
		creatorContactRepo: creatorContactRepo,
		creatorRepo:        creatorRepo,
		creatorAddressRepo: creatorAddressRepo,
		creatorEmailRepo:   creatorEmailRepo,
		creatorDataRepo:    creatorDataRepo,
		cfg:                cfg,
	}
}
