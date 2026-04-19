package creatorprofile

import (
	"context"
	"gotweet/internal/model"
	"net/http"
)

func (s *creatorProfileService) GetCreatorProfile(ctx context.Context, creator_id int64) (*model.CreatorProfiling, int, error) {
	res, err := s.creatorDataRepo.GetCreatorDataByCreatorID(ctx, creator_id)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	email_id := res.Email
	// address_id := res.Address
	creator_phone_number_id := res.PhoneNumber
	creator := res.Creator
	creator_status_id := res.Status
	creator_created_at := res.CreatedAt
	creator_updated_at := res.UpdatedAt
	creator_image_url := res.ImageUrl

	// Get Creator Email
	creatorEmail, err := s.creatorEmailRepo.GetCreatorEmail(ctx, email_id)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	// Get Email
	email, err := s.emailRepo.GetEmailString(ctx, creatorEmail.Email)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	// Get Creator Contact
	creator_contact, err := s.creatorContactRepo.GetCreatorContactByCreatorID(ctx, creator_phone_number_id)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}
	// Get Contact
	contact, err := s.contactRepo.GetContactById(ctx, creator_contact.ContactID)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}
	// Get Creator Image
	creator_image, err := s.creatorImageRepo.GetCreatorImageByCreatorID(ctx, creator_image_url)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	// Get Image Image repo has not build yet
	image, err := s.imageRepo.GetImageByID(ctx, creator_image.ImageUrl)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}
	// Get Creator
	creator_obj, err := s.creatorRepo.GetCreatorByID(ctx, creator_id)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	// Get Creator Status
	creator_status, err := s.creatorStatusRepo.GetCreatorStatus(ctx, creator_status_id)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}
	// Get Status
	status, err := s.statusRepo.GetStatusByStatusID(ctx, creator_status.Status)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}
	// Get Profile Picture
	return &model.CreatorProfiling{
		CreatorID:   creator,
		Email:       email.Email,
		PhoneNumber: contact.PhoneNumber,
		Image_url:   image.Image,
		Username:    creator_obj.Username,
		Status:      status.Status,
		CreatedAt:   creator_created_at,
		UpdatedAt:   creator_updated_at,
	}, http.StatusOK, nil
}
