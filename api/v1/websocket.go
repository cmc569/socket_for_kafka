package v1

import (
	"api/util/general"
	"api/util/websocket"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Connection(context *gin.Context) {
	var uuid string
	uuid = general.CreateUUid()
	pool := websocket.GetPool()
	fmt.Println("WebSocket Endpoint Hit")
	conn, err := websocket.WsUpgrade.Upgrade(context.Writer, context.Request, nil)
	if err != nil {
		fmt.Fprintf(context.Writer, "%+v\n", err)
	}
	client := &websocket.Client{
		ID:   uuid,
		Conn: conn,
		Pool: pool,
	}
	pool.Register <- client
	client.Read()
}
