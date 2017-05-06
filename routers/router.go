package routers

import (
	"beego_yc/controllers"
	"beego_yc/controllers/websocket"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.Router("/ws", &websocket.ClientController{}, "GET:Get")
	beego.Router("/ws/join", &websocket.IndexController{}, "GET:Join")
	beego.Router("/ws/push", &websocket.IndexController{}, "GET:Push")
}
