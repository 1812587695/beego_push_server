package main

import (
	_ "beego_yc/routers"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"

	"github.com/astaxie/beego"
)

func init() {
	// 注册驱动
	orm.RegisterDriver("mysql", orm.DRMySQL)
	// 注册默认数据库
	// 我的mysql的root用户密码为root，打算把数据表建立在名为beego数据库里
	// 备注：此处第一个参数必须设置为“default”（因为我现在只有一个数据库），否则编译报错说：必须有一个注册DB的别名为 default
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(192.168.210.19:3306)/yctp?charset=utf8")
}

func main() {

	// 开启 orm 调试模式：开发过程中建议打开，release时需要关闭
	orm.Debug = true

	beego.Run()
}
