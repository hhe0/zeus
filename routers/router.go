package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	controllerEmail "zeus/api/backend/email/controller"
	"zeus/api/backend/task/controller"
)

func init() {
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "content-type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
		AllowOrigins:     []string{"http://localhost:*", "http://127.0.0.1:*"},
	}))

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
}
