package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func Init() {
	dbhost := beego.AppConfig.String("dbhost")
	dbport := beego.AppConfig.String("dbport")
	dbuser := beego.AppConfig.String("dbuser")
	dbpassword := beego.AppConfig.String("dbpassword")
	dbname := beego.AppConfig.String("dbname")
	dsn := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8mb4&loc=Asia%2FShanghai"
	fmt.Println(dsn)

	err := orm.RegisterDataBase("default", "mysql", dsn)
	fmt.Println(err)
}

func TableName(str string) string {
	return beego.AppConfig.String("dbprefix") + str
}
