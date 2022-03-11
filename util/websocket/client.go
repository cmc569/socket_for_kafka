package websocket

import (
	"api/util/log"
	"github.com/gorilla/websocket"
)

type Client struct {
	ID   string
	Conn *websocket.Conn
	Pool *Pool
}

type Message struct {
	Type int         `json:"type"`
	Body interface{} `json:"body"`
}

func (client *Client) Read() {
	defer func() {
		client.Pool.Unregister <- client
		log.Error(client.Conn.Close())
	}()
	for {
		if messageType, p, err := client.Conn.ReadMessage(); err != nil {
			log.Error(err)
			return
		}else {
			client.Pool.SendMessage <- Message{Type: messageType, Body: string(p)}
		}
	}
}
