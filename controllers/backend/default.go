package backend

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"zeus/models"
)

type TODOListController struct {
	beego.Controller
}

func (c *TODOListController) Get() {
	o := orm.NewOrm()
	todoList := models.TaskInfo{
		UserId:     1,
		Content:    "这是一条todolist",
		FinishTime: 1565590544,
		CreateTime: 1565590544,
		UpdateTime: 1565590544,
	}

	id, err := o.Insert(&todoList)
	if err != nil {
		fmt.Println(err)
	}
	c.Data["json"] = id
	c.ServeJSON()
}
