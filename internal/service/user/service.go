package user

import (
	"context"
	"gotweet/internal/config"
	"gotweet/internal/dto"
	"gotweet/internal/repository/user"
)

type UserService interface {
	Register(ctx context.Context, req *dto.RegisterRequest) (int64, int, error)
	Login(ctx context.Context, req *dto.LoginRequst) (string, string, int, error)
	RefreshToken(ctx context.Context, req *dto.RefreshTokenRequest, userID int64) (string, string, int, error)
	LoginWithGoogle(ctx context.Context, req *dto.UserResGoogle) (string, string, int, error)
}

type userService struct {
	cfg      *config.Config
	userRepo user.UserRepository
}

func NewService(cfg *config.Config, userREpo user.UserRepository) UserService {
	return &userService{
		cfg:      cfg,
		userRepo: userREpo,
	}
}
