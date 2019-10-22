package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

type TODOList struct {
	Name string
}

func (c *MainController) Get() {
	data := &TODOList{
		Name: "何宏",
	}
	c.Data["json"] = data
	c.ServeJSON()
}
