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

	LoginWithGoogleReq struct {
		Token string `json:"token", validate:"required"`
	}

	UserResGoogle struct {
		Email          string `validate:"email"`
		Email_verified bool   `validate:"email_verified"`
		Family_name    string `validate:"family_name"`
		Given_name     string `validate:"given_name"`
		Name           string `validate:"name"`
		Picture        string `validate:"picture"`
		Token          string `validate:"token"`
		Role           string `validate:"role"`
	}
)
