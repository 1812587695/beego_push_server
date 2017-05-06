package models

import (
	//	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
)

//admin表
type WebsocketOrderMsg struct {
	Id       int
	Local_id int
	Time     string
	Msg      string
	Status   int
}

func (m *WebsocketOrderMsg) TableName() string {
	return "sdb_websocket_order_msg"
}

// 查询所有数据
func GetAllWebsocketOrderMsg() ([]*WebsocketOrderMsg, error) {
	o := orm.NewOrm()
	msg := make([]*WebsocketOrderMsg, 0)
	qs := o.QueryTable("sdb_websocket_order_msg")
	_, err := qs.All(&msg) // 返回总数量和错误信息

	return msg, err
}

func Query(id int) ([]orm.Params, error) {
	o := orm.NewOrm()
	var maps []orm.Params
	_, err := o.Raw("select * from sdb_websocket_order_msg where local_id = ?", id).Values(&maps)

	return maps, err
}

func WebsocketOrderMsgInsert(id int, msg string) (int64, error) {
	o := orm.NewOrm()
	time := time.Now()

	result, err := o.Raw("INSERT INTO `sdb_websocket_order_msg` (`local_id`, `time`, `msg`, `status`) VALUES (?, ?, ?, '1')", id, time, msg).Exec()

	if err == nil {
		num, err := result.LastInsertId()
		return num, err
	}

	return 0, err
}

func init() {
	// 需要在 init 中注册定义的 model
	orm.RegisterModel(new(WebsocketOrderMsg))
}
