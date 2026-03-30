package socket

import (
	"gotweet/internal/middleware"
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
	socket_router := h.api.Group("/ws")
	socket_router.Use(middleware.AuthMiddleware(secreKey))
	socket_router.POST("/ws/create_room", h.CreateRoom)
	socket_router.GET("/ws/join_room/:room_id", h.JoinRoom)
	socket_router.GET("/ws/get_clients/:room_id", h.GetClients)
}
