package creatordonation

import (
	"context"
	"gotweet/internal/model"
)

func (r *creatorDonationRepository) GetDonationById(ctx context.Context, id int64) (*model.CreatorDonation, error) {
	query := `SELECT * FROM creator_donation WHERE id = ?`
	row := r.db.QueryRowContext(ctx, query, id)
	var response model.CreatorDonation
	err := row.Scan(&response.ID, &response.SenderCreatorID, &response.ReceiverCreatorID, &response.TransactionID, &response.Amount, &response.Currency, &response.DonorName, &response.DonorEmail, &response.DonorMessage, &response.Message, &response.PaymentMethod, &response.ExternalReference, &response.Status, &response.Version, &response.CreatedAt, &response.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
