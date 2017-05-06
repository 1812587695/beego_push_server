package models

import (
	"fmt"

	"github.com/gorilla/websocket"
)

type Hub struct {
	Clients      map[*Client]int
	AddClient    chan *Client
	DeleteClient chan *Client
	Msg          chan []byte
}

func NewHub() *Hub {
	return &Hub{
		Clients:      make(map[*Client]int),
		AddClient:    make(chan *Client, 1),
		DeleteClient: make(chan *Client, 1),
		Msg:          make(chan []byte, 1),
	}
}

func (hub *Hub) Start() {
	fmt.Println("这里开始了")

	for {
		select {
		// 读取入channel中的数据（1channel）
		case conn := <-hub.AddClient:
			hub.Clients[conn] = conn.Local_id // 保存map
			fmt.Println("保存map", hub.Clients)
		case conn := <-hub.DeleteClient:

			if _, ok := hub.Clients[conn]; ok {
				delete(hub.Clients, conn)
				err := conn.Ws.Close()
				if err != nil {
					fmt.Println("defer---------------------ccccccccc")
				}
				fmt.Println("defer---------------------2", err)

				//goto ForEnd
			}

		// 读取channel中的数据，这里是获取用户发送的数据（2channel）
		case msg := <-hub.Msg:

			// 这里循环map里面保存的客户端连接信息
			for k, _ := range hub.Clients {

				//发送消息到每一个连接的客户端
				if k.Ws.WriteMessage(1, msg) != nil {

					NewHub().DeleteClient <- k
				}
			}
		}
	}
	//ForEnd:
	//	return
}

type Client struct {
	Ws       *websocket.Conn
	Local_id int
}
