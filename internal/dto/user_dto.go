package dto

type (
	RegisterRequest struct {
		Email           string `json:"email" validate:"required,email"`
		Username        string `json:"username" validate:"required,min=3"`
		Password        string `json:"password" validate:"required,min=8"`
		PasswordConfirm string `json:"password_confirm" validate:"required,min=8"`
	}

	RegisterResponse struct {
		ID int64 `json:"id"`
	}

	LoginRequst struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=8"`
	}

	LoginResponse struct {
		Token        string `json:"token"`
		RefreshToken string `json:"refresh_token"`
	}

	RefreshTokenRequest struct {
		RefreshToken string `json:"refresh_token" validate:"required"`
	}

	RefreshTokenResponse struct {
		RefreshToken string `json:"refresh_token" validate:"required"`
		Token        string `json:"token" validate:"required"`
	}
)
