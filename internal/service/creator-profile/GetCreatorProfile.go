package creatorprofile

import (
	"context"
	"gotweet/internal/model"
	"net/http"
)

func (s *creatorProfileService) GetCreatorProfile(ctx context.Context, cretor_id int64) (*model.CreatorProfiling, int, error) {
	res, err := s.creatorDataRepo.GetCreatorDataByCreatorID(ctx, cretor_id)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}
	// Get Profile Picture
	return nil, http.StatusOK, nil
}
