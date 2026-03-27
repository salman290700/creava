package user

import (
	"gotweet/internal/middleware"
	"gotweet/internal/service/user"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	api         *gin.Engine
	validate    *validator.Validate
	userService user.UserService
}

func NewHandler(api *gin.Engine, validate *validator.Validate, userService user.UserService) *Handler {
	return &Handler{
		api:         api,
		validate:    validate,
		userService: userService,
	}
}

func (h *Handler) RouteList(secreKey string) {
	authRoute := h.api.Group("/auth")
	authRoute.POST("/register", h.Register)
	authRoute.POST("/login", h.Login)
	authRoute.POST("/oauth", h.LoginWithGoogle)

	refreshRouter := h.api.Group("/auth")
	refreshRouter.Use(middleware.AuthRefreshToken(secreKey))
	refreshRouter.POST("/refresh_token")
}
