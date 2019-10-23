package model

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type BaseModel struct {
	Id         int `json:"id"`
	CreateTime int `json:"create_time"`
	UpdateTime int `json:"update_time"`
	IsDeleted  int `json:"is_deleted" orm:"default(0)"`
}

func Init() {
	dbhost := beego.AppConfig.String("dbhost")
	dbport := beego.AppConfig.String("dbport")
	dbuser := beego.AppConfig.String("dbuser")
	dbpassword := beego.AppConfig.String("dbpassword")
	dbname := beego.AppConfig.String("dbname")
	dsn := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8mb4&loc=Asia%2FShanghai"

	orm.RegisterDataBase("default", "mysql", dsn)
}

func CompletedTableName(str string) string {
	return beego.AppConfig.String("dbprefix") + str
}
