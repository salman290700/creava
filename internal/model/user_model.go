package model

import "time"

type (
	UserModel struct {
		ID        int64
		Email     string
		Username  string
		Password  string
		CreatedAt time.Time
		UpdatedAt time.Time
	}

	RefreshTokenModel struct {
		ID           int64
		UserID       int64
		RefreshToken string
		CreatedAt    time.Time
		UpdatedAt    time.Time
		ExpiredAt    time.Time
	}
)
