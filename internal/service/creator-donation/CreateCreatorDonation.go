package creatordonation

import (
	"context"
	"gotweet/internal/dto"
	"gotweet/internal/model"
	"net/http"
	"time"
)

func (s *creatorDonationService) CreateCreatorDonation(ctx context.Context, request *dto.CreatorDonationRequest, creator_id int64) (int64, int, error) {
	// Get Balance First
	balance, err := s.creatorBalanceRepo.GetCreatorBalanceByCreatorID(ctx, creator_id)
	if err != nil {
		return 0, http.StatusBadRequest, err
	}
	// Create Transaction Then
	now := time.Now()
	id_transaction, err := s.creatorTransactionRepo.CreateCreatorTransaction(ctx, creator_id, balance.ID, "balance", request.Amount, request.Currency, request.ExternalReference, "", request.Status, 1, now, now)
	if err != nil {
		return 0, http.StatusBadRequest, err
	}
	// Create Donation Then
	id_donation, err := s.creatorDonationRepos.CreateCreatorDonation(ctx, &model.CreatorDonation{
		SenderCreatorID:   request.SenderCreatorID,
		ReceiverCreatorID: request.ReceiverCreatorID,
		TransactionID:     id_transaction,
		Amount:            request.Amount,
		Currency:          request.Currency,
		DonorName:         request.DonorName,
		DonorEmail:        request.DonorEmail,
		DonorMessage:      request.DonorMessage,
		Message:           request.Message,
		PaymentMethod:     request.PaymentMethod,
		ExternalReference: request.ExternalReference,
		Status:            request.Status,
		CreatedAt:         now,
		UpdatedAt:         now,
	})
	return id_donation, http.StatusOK, nil
}
