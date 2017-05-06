package websocket

import (
	"fmt"

	"github.com/astaxie/beego"
)

type ClientController struct {
	beego.Controller
}

func (this *ClientController) Get() {

	store_id := this.GetString("store_id")
	fmt.Println("=============", store_id)

	this.Data["store_id"] = store_id
	this.TplName = "websocket/index/index.html"
	fmt.Println("++++++++++++++", store_id)
}
