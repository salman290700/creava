package creator

import (
	"context"
	"gotweet/internal/dto"
	"gotweet/internal/model"
	"gotweet/pkg/jwt"
	refreshtoken "gotweet/refresh_token"
	"net/http"
	"time"
)

func (s *creatorService) LoginWithGoogle(ctx context.Context, data *dto.UserResGoogle) (string, string, int64, error) {
	var (
		creator_id int64
		email      string
	)
	// Get CreatorData First
	res, err := s.creatorDataRepo.GetCreatorData(ctx, data.Email)
	if err != nil {
		return "", "", http.StatusBadRequest, err
	}
	// If Null Create Creator and insert email
	name := data.Given_name + data.Family_name
	if res == nil {
		creator_id_res, err := s.creatorRepo.CreateCreator(ctx, name, "")
		if err != nil {
			return "", "", http.StatusBadRequest, err
		}
		creator_id = creator_id_res
		_, err = s.emailRepo.CreateEmail(ctx, data.Email, data.Email_verified)
		if err != nil {
			return "", "", http.StatusBadRequest, err
		}
		email = data.Email
	}

	// Get creator_id, email
	if res != nil {
		creator_id = res.ID
	}
	// Create jwt token
	token, err := jwt.CreateToken(creator_id, email, s.cfg.SecretJwt)
	if err != nil {
		return "", "", http.StatusInternalServerError, err
	}
	// Create Refresh Token
	refresh_token, err := refreshtoken.GenerateRefreshToken()
	if err != nil {
		return "", "", http.StatusBadRequest, err
	}
	now := time.Now()
	err = s.creatorDataRepo.StoreRefreshToken(ctx, &model.RefreshTokenCreatorModel{
		CreatorID:    creator_id,
		RefreshToken: refresh_token,
		CreatedAt:    now,
		UpdatedAt:    now,
		ExpiredAt:    time.Now().Add(7 * 24 * time.Hour),
	}, creator_id)
	if err != nil {
		return "", "", http.StatusInternalServerError, err
	}
	return token, refresh_token, http.StatusBadRequest, nil
}
