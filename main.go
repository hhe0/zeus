package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "zeus/routers"
)

func init() {
	// TODO: 处理错误日志
	_ = orm.RegisterDriver("mysql", orm.DRMySQL)
	_ = orm.RegisterDataBase("default", "mysql", "root:root@/orm_test?charset=utf8")
}

func main() {
	beego.Run()
}
