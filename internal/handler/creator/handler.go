package creator

import (
	"gotweet/internal/service/creator"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	api            *gin.Engine
	validate       *validator.Validate
	creatorService creator.CreatorService
}

func NewCreatorHandler(
	api *gin.Engine,
	validate *validator.Validate,
	creatorService creator.CreatorService,
) *Handler {
	return &Handler{
		api:            api,
		validate:       validate,
		creatorService: creatorService,
	}
}

func (h *Handler) CreatorHandlerRouteList(secretKey string) {
	creatorRoute := h.api.Group("/creator")
	creatorRoute.POST("/login-google-auth", h.LoginWithGoogle)
}
