package socket

import (
	socket "gotweet/websocket"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateRoom(c *gin.Context) {
	var req socket.CreateRoomRequest
	if err := c.Copy().ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	h.hub.Rooms[req.ID] = &socket.Room{
		ID:      req.ID,
		Name:    req.Name,
		Clients: make(map[string]*socket.Client),
	}
	c.JSON(http.StatusOK, req)
}
