package controller

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"zeus/api/backend/task/request"
	"zeus/api/backend/task/response"
	"zeus/manager/task/model"
	"zeus/manager/task/service"
)

type TaskController struct {
	beego.Controller
}

func (ctr *TaskController) CreateTask() {
	var req request.CreateTaskRequest
	if err := json.Unmarshal(ctr.Ctx.Input.RequestBody, &req); err != nil {
		fmt.Println("error")
	}

	service.NewTaskInfoService().InsertInfo(model.TaskInfo{
		// TODO: userId登录后做特殊处理
		UserId:  1,
		Content: req.Content,
	})

	ctr.Data["json"] = response.CreateTaskResponse{
		Code:    0,
		Message: "新增成功",
	}
	ctr.ServeJSON()
}
