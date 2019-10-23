package routers

import (
	"github.com/astaxie/beego"
	"zeus/api/backend/task/controller"
)

func init() {
	// 任务清单
	beego.Router("/api/backend/task-list", &controller.TaskController{}, "post:CreateTask") // 新增任务
}
