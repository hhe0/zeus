package routers

import (
	"github.com/astaxie/beego"
	"zeus/api/backend/task/controller"
)

func init() {
	// 新增任务
	beego.Router("/api/backend/task", &controller.TaskController{}, "post:CreateTask")
	// 任务列表
	beego.Router("/api/backend/task-list", &controller.TaskController{}, "get:GetTaskList")
}
