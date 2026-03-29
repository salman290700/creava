package socket

import (
	socket "gotweet/websocket"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	api *gin.Engine
	// validate    *validator.Validate
	// userService user.UserService
	hub *socket.Hub
}

func NewHandler(api *gin.Engine, hub *socket.Hub) *Handler {
	return &Handler{
		api: api,
		// validate:    validate,
		// userService: userService,
		hub: hub,
	}
}

func (h *Handler) RouteList(secreKey string) {
	h.api.POST("/ws/createRoom", h.CreateRoom)
}
