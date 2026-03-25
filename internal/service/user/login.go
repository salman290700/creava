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

	"golang.org/x/crypto/bcrypt"
)

func (s *userService) Login(ctx context.Context, loginDto *dto.LoginRequst) (string, string, int, error) {
	// Check user
	user, err := s.userRepo.GetUserByEmailOrUsername(ctx, loginDto.Email, "")
	if err != nil {
		return "", "", http.StatusBadRequest, err
	}

	if user == nil {
		return "", "", http.StatusBadRequest, errors.New("Wrong email or Password")
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginDto.Password)); err != nil {
		return "", "", http.StatusNotFound, errors.New("wrong email or password")
	}

	if user == nil {
		return "", "", http.StatusBadRequest, errors.New("user is not found")
	}
	// Generate Access Token
	token, err := jwt.CreateToken(user.ID, user.Username, s.cfg.SecretJwt)
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
