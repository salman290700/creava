package socket

import "github.com/gorilla/websocket"

// Client represen websocket client
type Client struct {
	Conn     *websocket.Conn
	Message  chan *Message
	ID       string `json:"id"`
	RoomId   string `json:"room_id"`
	Username string `json:"username`
}

type Message struct {
	Content  string `json:"content"`
	RoomID   string `json:"roomId"`
	Username string `json:"username"`
}
