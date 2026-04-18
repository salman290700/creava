package creatordonation

import (
	"context"
	"gotweet/internal/model"
)

func (r *creatorDonationRepository) CreateCreatorDonation(ctx context.Context, data *model.CreatorDonation) (int64, error) {
	query := `INSERT INTO creator_donation (sender_creator_id, receiver_creator_id, transaction_id, amount, currency, donor_name, donor_email, donor_message, message, payment_method, external_reference, status, version, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	res, err := r.db.ExecContext(ctx, query, data.SenderCreatorID, data.ReceiverCreatorID, data.TransactionID, data.Amount, data.Currency, data.DonorName, data.DonorEmail, data.DonorMessage, data.Message, data.PaymentMethod, data.ExternalReference, data.Status, data.Version, data.CreatedAt, data.UpdatedAt)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, nil
	}
	return id, nil
}
