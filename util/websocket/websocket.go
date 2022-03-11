package websocket

import (
	"github.com/gorilla/websocket"
	"net/http"
)

var WsUpgrade = websocket.Upgrader{
	ReadBufferSize:  4096,
	WriteBufferSize: 4096,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}
