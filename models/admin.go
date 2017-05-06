package models

import (
	"crypto/md5"
	"fmt"

	"github.com/astaxie/beego/orm"
)

//admin表
type Admin struct {
	Id       int
	Name     string
	Password string
}

func (u *Admin) TableName() string {
	//	return "bbs_admin"
	return "admin"
}

func GetOneAdmin(name, pwd string) (*Admin, error) {

	o := orm.NewOrm()
	admin := new(Admin)
	//	qs := o.QueryTable("bbs_admin") 等于 qs := o.QueryTable(admin)
	qs := o.QueryTable(admin)

	byte_pwd := []byte(pwd)
	has_pwd := md5.Sum(byte_pwd)
	pwd = fmt.Sprintf("%x", has_pwd) //将[]byte转成16进制

	err := qs.Filter("name", name).Filter("password", pwd).One(admin)
	if err != nil {
		return nil, err
	}
	return admin, err
}

// 添加一条数据
func Add(name, pwd string) (int64, error) {

	admin := new(Admin)

	byte_pwd := []byte(pwd)
	has_pwd := md5.Sum(byte_pwd)
	pwd = fmt.Sprintf("%x", has_pwd) //将[]byte转成16进制

	admin.Name = name
	admin.Password = pwd
	id, err := orm.NewOrm().Insert(admin)
	return id, err
}

// 查询所有数据
func GetAllAdmin() ([]*Admin, error) {
	o := orm.NewOrm()
	admin := make([]*Admin, 0)
	qs := o.QueryTable("bbs_admin")
	_, err := qs.All(&admin) // 返回总数量和错误信息

	return admin, err
}
func init() {
	// 需要在 init 中注册定义的 model
	//	orm.RegisterModel(new(Admin))

	// 表前缀设置
	orm.RegisterModelWithPrefix("bbs_", new(Admin))
}
