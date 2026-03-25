package user

import (
	"context"
	"errors"
	"gotweet/internal/dto"
	"gotweet/internal/model"
	"gotweet/pkg/jwt"
	refreshtoken "gotweet/refresh_token"
	"net/http"
	"time"
)

func (s *userService) RefreshToken(ctx context.Context, req *dto.RefreshTokenRequest, userID int64) (string, string, int, error) {
	// Check user exists
	userExists, err := s.userRepo.GetUserbyId(ctx, userID)
	if err != nil {
		return "", "", http.StatusInternalServerError, err
	}

	if userExists == nil {
		return "", "", http.StatusInternalServerError, err
	}
	// Get Refresh_token by user_id

	refresh_token, err := s.userRepo.GetRefreshtoken(ctx, userID, time.Now())
	if err != nil {
		return "", "", http.StatusInternalServerError, err
	}

	if refresh_token == nil {
		return "", "", http.StatusUnauthorized, errors.New("refresh token was expired")
	}

	// check refresh token is match with request body
	if req.RefreshToken != refresh_token.RefreshToken {
		return "", "", http.StatusUnauthorized, errors.New("unauthorized access")
	}
	user, err := s.userRepo.GetUserbyId(ctx, refresh_token.UserID)
	// Generate new token

	token, err := jwt.CreateToken(refresh_token.UserID, user.Username, s.cfg.SecretJwt)
	if err != nil {
		return "", "", http.StatusInternalServerError, err
	}

	// delete old refresh token & generate new refresh token
	err = s.userRepo.DeleteRefreshToken(ctx, userID)
	if err != nil {
		return "", "", http.StatusInternalServerError, err
	}

	new_refresh_token, err := refreshtoken.GenerateRefreshToken()
	if err != nil {
		return "", "", http.StatusInternalServerError, err
	}
	now := time.Now()

	err = s.userRepo.StoreRefreshToken(ctx, &model.RefreshTokenModel{
		UserID:       userID,
		RefreshToken: new_refresh_token,
		CreatedAt:    now,
		UpdatedAt:    now,
		ExpiredAt:    time.Now().Add(7 * 24 * time.Hour),
	}, userID)
	if err != nil {
		return "", "", http.StatusInternalServerError, err
	}

	return token, refresh_token.RefreshToken, http.StatusOK, nil
}
