package socket

import (
	socket "gotweet/websocket"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (h *Handler) JoinRoom(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}

	roomID := c.Param("room_id")
	clientID := c.Param("client_id")
	username := c.Param("username")

	cl := &socket.Client{
		Conn:     conn,
		Message:  make(chan *socket.Message),
		ID:       clientID,
		RoomId:   roomID,
		Username: username,
	}

	m := &socket.Message{
		Content:  "A new User has joined this room",
		RoomID:   roomID,
		Username: username,
	}

	h.hub.Register <- cl
	h.hub.Broadcast <- m
	go cl.WriteMessage()
	cl.ReadMessage(h.hub)
}


