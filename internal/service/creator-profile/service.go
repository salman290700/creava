package creatorprofile

import (
	"context"
	"gotweet/internal/config"
	"gotweet/internal/dto"
	"gotweet/internal/model"
	"gotweet/internal/repository/address"
	"gotweet/internal/repository/contact"
	"gotweet/internal/repository/creator"
	creatorRepository "gotweet/internal/repository/creator"
	creatoraddress "gotweet/internal/repository/creator-address"
	creatorbalance "gotweet/internal/repository/creator-balance"
	creatorContact "gotweet/internal/repository/creator-contact"
	creatordata "gotweet/internal/repository/creator-data"
	creatoremail "gotweet/internal/repository/creator-email"
	creatorimage "gotweet/internal/repository/creator-image"
	creatorstatus "gotweet/internal/repository/creator-status"
	"gotweet/internal/repository/email"
	"gotweet/internal/repository/image"
	oldusername "gotweet/internal/repository/old-username"
	"gotweet/internal/repository/status"
)

type CreatorProfileService interface {
	GetCreatorProfile(ctx context.Context, cretor_id int64) (*model.CreatorProfiling, int, error)
	UploadCreatorProfile(ctx context.Context, request *dto.UploadCreatorProfileDTO) (*model.CreatorProfiling, int, error)
}

type creatorProfileService struct {
	statusRepo         status.StatusRepository
	addressRepo        address.AddressRepository
	emailRepo          email.EmailRepository
	contactRepo        contact.ContactRepository
	imageRepo          image.ImageRepository
	creatorContactRepo creatorContact.CreatorContactRepository
	creatorRepo        creatorRepository.CreatorRepository
	creatorAddressRepo creatoraddress.CreatorAddressRepository
	creatorEmailRepo   creatoremail.CreatorEmailRepository
	creatorDataRepo    creatordata.CreatorDataRepository
	creatorbalancerepo creatorbalance.CreatorBalanceRepository
	creatorImageRepo   creatorimage.CreatorImageRepository
	creatorStatusRepo  creatorstatus.CreatorStatusRepository
	oldUsernameRepo    oldusername.OldUsernameRepository
	cfg                *config.Config
}

func NewCreatorProfileService(
	statusRepo status.StatusRepository,
	addressRepo address.AddressRepository,
	emailRepo email.EmailRepository,
	contactRepo contact.ContactRepository,
	imageRepo image.ImageRepository,
	creatorContactRepo creatorContact.CreatorContactRepository,
	creatorRepo creator.CreatorRepository,
	creatorAddressRepo creatoraddress.CreatorAddressRepository,
	creatorEmailRepo creatoremail.CreatorEmailRepository,
	creatorDataRepo creatordata.CreatorDataRepository,
	creatorbalancerepo creatorbalance.CreatorBalanceRepository,
	creatorImageRepo creatorimage.CreatorImageRepository,
	creatorStatusRepo creatorstatus.CreatorStatusRepository,
	oldUsernameRepo oldusername.OldUsernameRepository,
	cfg *config.Config,
) CreatorProfileService {
	return &creatorProfileService{
		statusRepo:         statusRepo,
		addressRepo:        addressRepo,
		emailRepo:          emailRepo,
		contactRepo:        contactRepo,
		imageRepo:          imageRepo,
		creatorContactRepo: creatorContactRepo,
		creatorRepo:        creatorRepo,
		creatorAddressRepo: creatorAddressRepo,
		creatorEmailRepo:   creatorEmailRepo,
		creatorDataRepo:    creatorDataRepo,
		creatorbalancerepo: creatorbalancerepo,
		creatorImageRepo:   creatorImageRepo,
		creatorStatusRepo:  creatorStatusRepo,
		oldUsernameRepo:    oldUsernameRepo,
		cfg:                cfg,
	}
}
