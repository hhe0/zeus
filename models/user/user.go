package user

import "github.com/astaxie/beego/orm"

type User struct {
	Id            int `orm:"column(Id)"`
	LastLoginTime int `orm:"column(last_login_time)"`
	CreateTime    int `orm:"column(create_time)"`
	UpdateTime    int `orm:"column(update_time)"`
}

func InsertInfo(user *User) User {
	o := orm.NewOrm()
	o.Using("default")
	o.Insert(user) //插入数据库

	return *user
}
