package creatordonation

import (
	"context"
	"gotweet/internal/config"
	"gotweet/internal/dto"
	creatorbalance "gotweet/internal/repository/creator-balance"
	creatordonation "gotweet/internal/repository/creator-donation"
	creatortransaction "gotweet/internal/repository/creator-transaction"

	"github.com/gin-gonic/gin"
)

type CreatorDonationService interface {
	CreateCreatorDonation(ctx context.Context, request *dto.CreatorDonationRequest, creator_id int64) (int64, int, error)
}

type creatorDonationService struct {
	// Creatorbalance Repository
	creatorBalanceRepo creatorbalance.CreatorBalanceRepository
	// CratorTransaction Repository
	creatorTransactionRepo creatortransaction.CreatorTransactionRepository
	// CreatorDonationRepository
	creatorDonationRepos creatordonation.CreatorDonationRepository
	// Engine
	api *gin.Engine
	//
	cfg *config.Config
}

func NewCreatorDonationService(creatorBalanceRepo creatorbalance.CreatorBalanceRepository, creatorTransactionRepo creatortransaction.CreatorTransactionRepository, creatorDonationRepo creatordonation.CreatorDonationRepository, api *gin.Engine, cfg *config.Config) CreatorDonationService {
	return &creatorDonationService{
		creatorBalanceRepo:     creatorBalanceRepo,
		creatorTransactionRepo: creatorTransactionRepo,
		creatorDonationRepos:   creatorDonationRepo,
		api:                    api,
		cfg:                    cfg,
	}
}
