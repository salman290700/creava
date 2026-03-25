package user

import (
	"context"
	"database/sql"
	"gotweet/internal/model"
	"time"
)

type UserRepository interface {
	GetUserByEmailOrUsername(ctx context.Context, email, username string) (*model.UserModel, error)
	CreateUser(ctx context.Context, model *model.UserModel) (int64, error)
	GetRefreshtoken(ctx context.Context, id int64, now time.Time) (*model.RefreshTokenModel, error)
	CreateRefreshToken(ctx context.Context, id int64, refresh_token string) error
	GetRefreshtokenByToken(ctx context.Context, token string) (*model.RefreshTokenModel, error)
	StoreRefreshToken(ctx context.Context, refresh_token *model.RefreshTokenModel, userID int64) error
	GetUserbyId(ctx context.Context, id int64) (*model.UserModel, error)
	DeleteRefreshToken(ctx context.Context, userID int64) error
}

type userRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}
