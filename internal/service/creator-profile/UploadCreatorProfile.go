package creatorprofile

import (
	"context"
	"gotweet/internal/dto"
	"gotweet/internal/model"
	"net/http"
)

func (s *creatorProfileService) UploadCreatorProfile(ctx context.Context, request *dto.UploadCreatorProfileDTO) (*model.CreatorProfiling, int, error) {
	// Check if the Username Change
	// Get username
	creator_data, err := s.creatorDataRepo.GetCreatorDataByCreatorID(ctx, request.CreatorID)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	// Save old username with creator_id
	_, err = s.oldUsernameRepo.CreateOldUsername(ctx, creator_data.Username, creator_data.Creator)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	// Save username in creator_data
	
	return nil, http.StatusBadRequest, nil
}
