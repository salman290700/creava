package socket

import (
	socket "gotweet/websocket"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetClients(c *gin.Context) {
	var clients []socket.ClientRes

	roomID := c.Param("room_id")

	if _, ok := h.hub.Rooms[roomID]; !ok {
		clients = make([]socket.ClientRes, 0)
		c.JSON(http.StatusOK, clients)
	}

	for _, c := range h.hub.Rooms[roomID].Clients {
		clients = append(clients, socket.ClientRes{
			ID:       c.ID,
			Username: c.Username,
		})
	}
	c.JSON(http.StatusOK, clients)
}
