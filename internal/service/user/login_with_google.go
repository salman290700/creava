package user

import (
	"context"
	"gotweet/internal/dto"
	"gotweet/internal/model"
	"gotweet/pkg/jwt"
	refreshtoken "gotweet/refresh_token"
	"net/http"
	"time"
)

func (s *userService) LoginWithGoogle(ctx context.Context, req *dto.UserResGoogle) (string, string, int, error) {
	// Check user
	userModel := &model.UserModel{
		Email:     req.Email,
		Username:  req.Email,
		Password:  "",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	var userId int64
	user, err := s.userRepo.GetUserByEmailOrUsername(ctx, req.Email, "")
	if err != nil {
		return "", "", http.StatusBadRequest, err
	}

	if user == nil {
		userId, err = s.userRepo.CreateUser(ctx, userModel)
		if err != nil {
			return "", "", http.StatusInternalServerError, err
		}
	} else {
		userId = user.ID
	}

	

	// Generate Access Token
	token, err := jwt.CreateToken(userId, user.Username, s.cfg.SecretJwt)
	if err != nil {
		return "", "", http.StatusInternalServerError, err
	}

	// Get Refresh Token if exists
	now := time.Now()
	refresh_token_obj, err := s.userRepo.GetRefreshtoken(ctx, user.ID, now)
	if err != nil {
		return "", "", http.StatusInternalServerError, err
	}
	if refresh_token_obj != nil {
		return token, refresh_token_obj.RefreshToken, http.StatusOK, nil
	}
	// generate & store refresh token
	refresh_token, err := refreshtoken.GenerateRefreshToken()
	if err != nil {
		return "", "", http.StatusBadRequest, err
	}

	err = s.userRepo.StoreRefreshToken(ctx, &model.RefreshTokenModel{
		UserID:       user.ID,
		RefreshToken: refresh_token,
		CreatedAt:    now,
		UpdatedAt:    now,
		ExpiredAt:    time.Now().Add(7 * 24 * time.Hour),
	}, user.ID)
	if err != nil {
		return "", "", http.StatusInternalServerError, err
	}
	return token, refresh_token, http.StatusAccepted, nil
}
