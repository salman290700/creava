package dto

import "time"

type (
	CreateDtoReq struct {
		Name  string `json:"name" validate:"required"`
		Email string `json:"email" validate:"required"`
	}

	CreatorDonationRequest struct {
		SenderCreatorID   int64     `json:"sender_creator_id" validate:"required"`
		ReceiverCreatorID int64     `json:"receiver_creator_id" validate:"required"`
		TransactionID     int64     `json:"transaction_id" validate:"required"`
		Amount            float64   `json:"amount" validate:"required"`
		Currency          string    `json:"currency" validate:"required"`
		DonorName         string    `json:"donor_name"`
		DonorEmail        string    `json:"donor_email"`
		DonorMessage      string    `json:"donor_message"`
		Message           string    `json:"message"`
		PaymentMethod     string    `json:"payment_method"`
		ExternalReference int64     `json:"external_reference"`
		Status            string    `json:"status"`
		Version           int64     `json:"version"`
		CreatedAt         time.Time `json:"created_at"`
		UpdatedAt         time.Time `json:"updated_at"`
	}

	CreatorDonationRespose struct {
		ID int64 `json:"id"`
	}
)
