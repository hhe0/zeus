package routers

import (
	"github.com/astaxie/beego"
	controllerEmail "zeus/api/backend/email/controller"
	"zeus/api/backend/task/controller"
	controllerUser "zeus/api/backend/user/controller"
)

func init() {
	// 新增任务
	beego.Router("/api/backend/task", &controller.TaskController{}, "post:CreateTask")
	// 任务列表
	beego.Router("/api/backend/task/list", &controller.TaskController{}, "get:GetTaskList")
	// 更改任务状态
	beego.Router("/api/backend/task/status", &controller.TaskController{}, "patch:UpdateTaskStatus")
	// 删除任务
	beego.Router("/api/backend/task", &controller.TaskController{}, "delete:DeleteTask")

	// 用户发送邮件提醒
	beego.Router("/api/backend/email/send", &controllerEmail.EmailController{}, "POST:SendEmail")

	// TODO: 用户登录
	beego.Router("/api/backend/user/login", &controllerUser.UserController{}, "POST:UserLogin")
	// TODO: 获取用户信息
	beego.Router("/api/backend/user/info", &controllerUser.UserController{}, "GET:GetUserInfo")
}
