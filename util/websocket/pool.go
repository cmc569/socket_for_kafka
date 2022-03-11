package websocket

import (
	"api/util/log"
	"encoding/json"
	"fmt"
)

type Pool struct {
	Register    chan *Client
	Unregister  chan *Client
	ClientMap   map[string]*Client
	Broadcast   chan Message
	SendMessage chan Message
}

var pool *Pool

func NewPool() *Pool {
	pool = &Pool{
		Register:    make(chan *Client),
		Unregister:  make(chan *Client),
		ClientMap:   make(map[string]*Client),
		Broadcast:   make(chan Message),
		SendMessage: make(chan Message),
	}
	return pool
}

func GetPool() *Pool {
	return pool
}

func (pool *Pool) Start() {
	for {
		select {
		case client := <-pool.Register:
			pool.ClientMap[client.ID] = client
			fmt.Println(client.ID + " Register")
			fmt.Println("Size of Connection Pool: ", len(pool.ClientMap))
			jsonMap := map[string]interface{}{
				"from":    "socket_server",
				"project": "order_system",
				"action":  "register",
				"status":  true,
				"message": "register successfully",
				"uuid":    client.ID,
			}
			log.Error(client.Conn.WriteJSON(jsonMap))
			break
		case client := <-pool.Unregister:
			delete(pool.ClientMap, client.ID)
			fmt.Println(client.ID + " Unregister")
			fmt.Println("Size of Connection Pool: ", len(pool.ClientMap))
			break
		case message := <-pool.SendMessage:
			type ReceivedData struct {
				Uuid    string `json:"uuid"`
				Status  bool   `json:"status"`
				Message string `json:"message"`
				From    string `json:"from"`
				Project string `json:"project"`
				Action  string `json:"action"`
			}
			var data ReceivedData
			fmt.Println("send message")
			s, _ := message.Body.(string)
			fmt.Println(s)
			if err := json.Unmarshal([]byte(s), &data); err != nil {
				log.Error(err)
				return
			}
			if toClient, isOk := pool.ClientMap[data.Uuid]; isOk {
				log.Error(toClient.Conn.WriteJSON(data))
			}
			break
		case message := <-pool.Broadcast:
			fmt.Println("Sending message to all clients in Pool")
			for _, client := range pool.ClientMap {
				if err := client.Conn.WriteJSON(message); err != nil {
					fmt.Println(err)
					return
				}
			}
		}

	}
}
