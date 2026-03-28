package socket

import "github.com/gorilla/websocket"

// Client represen websocket client
type Client struct {
	Conn *websocket.Conn
	Send chan []byte
	ID   string
}


