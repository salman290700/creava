package dto

type (
	CreateDtoReq struct {
		Name  string `json:"name" validate:"required"`
		Email string `json:"email" validate:"required"`
	}
)
