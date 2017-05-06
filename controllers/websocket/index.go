package websocket

import (
	"beego_yc/models"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
)

type IndexController struct {
	beego.Controller
}

var Hub = models.NewHub()

func init() {
	go Hub.Start()
}

// 建立websocket链接
func (this *IndexController) Join() {

	//beego.BConfig.WebConfig.AutoRender = false

	store_id := this.GetString("store_id")
	if store_id == "" {
		fmt.Println("store_id不能为空")
	}
	id, _ := strconv.Atoi(store_id)

	ws, err := websocket.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil, 1024, 1024)

	if err != nil {
		fmt.Println("链接报错")
	}

	// 将连接保存struct
	client := &models.Client{
		Ws:       ws,
		Local_id: id,
	}

	defer func() {
		Hub.DeleteClient <- client
	}()

	//	// 数据库查询
	order_msg, err := models.Query(id)
	fmt.Println(order_msg)

	if err != nil {
		if ws.WriteJSON("{\"data\":["+""+"]}") != nil {
			return
		}
	}

	//map转json
	str, err := json.Marshal(order_msg)

	if err != nil {
		if ws.WriteJSON("{\"data\":["+""+"]}") != nil {
			return
		}
	}

	// 发送消息到客户端
	if ws.WriteJSON("{\"data\":"+string(str)+"}") != nil {
		return
	}

	// 将client写入channel中（1channe）

	Hub.AddClient <- client

	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			//if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
			fmt.Println("链接已经关闭")
			//}
			return
		}

		//将msg保存在channel中（2channel）
		Hub.Msg <- msg
	}

}

func (this *IndexController) Push() {
	beego.BConfig.WebConfig.AutoRender = false

	store_id := this.GetString("store_id")
	if store_id == "" {
		this.Ctx.WriteString("store_id不能为空")
	}
	id, _ := strconv.Atoi(store_id)

	msg := this.GetString("msg")
	if msg == "" {
		this.Ctx.WriteString("msg不能为空")
	}

	num, err := models.WebsocketOrderMsgInsert(id, msg)
	if err != nil && num == 0 {
		this.Ctx.WriteString("数据保存失败")
	}

	//判断map是否有连接
	if len(Hub.Clients) == 0 {
		this.Ctx.WriteString("客户端没有连接，消息推送失败")
	}

	map_msg := make(map[string]interface{})
	map_msg["id"] = store_id
	map_msg["msg"] = msg

	str, err := json.Marshal(map_msg)

	for k, v := range Hub.Clients {
		if v == id {
			//发送消息到客户端
			err := k.Ws.WriteJSON("{\"data\":[" + string(str) + "]}")
			if err != nil {
				this.Ctx.WriteString("客户端没有连接，消息推送失败")
			} else {
				this.Ctx.WriteString("success")
			}
		}

	}
}
