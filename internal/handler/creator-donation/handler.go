package creatordonation

import (
	"gotweet/internal/middleware"
	creatordonation "gotweet/internal/service/creator-donation"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	api                    *gin.Engine
	validate               *validator.Validate
	creatorDonationService creatordonation.CreatorDonationService
}

func NewCreatorDonationHandler(api *gin.Engine, validate *validator.Validate, creatorDonationService creatordonation.CreatorDonationService) *Handler {
	return &Handler{
		api:                    api,
		validate:               validate,
		creatorDonationService: creatorDonationService,
	}
}

func (h *Handler) CreatorHandlerRouteList(secretKey string) {
	creatorDonationRoute := h.api.Group("/creator")
	creatorDonationRoute.Use(middleware.AuthMiddleware(secretKey))
	creatorDonationRoute.POST("/creato-creator-donation", h.CreateCreatorDonation)
}
