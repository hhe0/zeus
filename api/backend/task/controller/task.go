package controller

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"zeus/api/backend/task/request"
	"zeus/api/backend/task/response"
	"zeus/manager/task/model"
	"zeus/manager/task/service"
)

type TaskController struct {
	beego.Controller
}

func (ctrl *TaskController) CreateTask() {
	var (
		req       request.CreateTaskRequest
		validator validation.Validation
	)
	_ = json.Unmarshal(ctrl.Ctx.Input.RequestBody, &req)
	if _, err := validator.Valid(&req); err != nil {
		// TODO: 校验失败的处理
		ctrl.Data["json"] = response.CreateTaskResponse{
			Code:    100001,
			Message: "输入参数有误x",
		}
		ctrl.ServeJSON()
	}

	service.NewTaskInfoService().InsertInfo(model.TaskInfo{
		// TODO: userId登录后做特殊处理
		UserId:  1,
		Content: req.Content,
	})

	ctrl.Data["json"] = response.CreateTaskResponse{
		Code:    0,
		Message: "成功",
	}
	ctrl.ServeJSON()
}

func (ctrl *TaskController) GetTaskList() {
	status, err := ctrl.GetInt("status")
	if err != nil {
		ctrl.Data["json"] = response.CreateTaskResponse{
			Code:    100001,
			Message: "输入参数有误",
		}
		ctrl.ServeJSON()
	}

	total, data := service.NewTaskInfoService().GetTodayList(status)
	ctrl.Data["json"] = response.GetTaskListResponse{
		Code:    0,
		Message: "成功",
		Data: response.GetTaskListData{
			Total:    total,
			TaskList: data,
		},
	}
	ctrl.ServeJSON()
}

func (ctrl *TaskController) UpdateTaskStatus() {
	var req request.UpdateTaskStatusRequest
	if err := json.Unmarshal(ctrl.Ctx.Input.RequestBody, &req); err != nil {
		ctrl.Data["json"] = response.CreateTaskResponse{
			Code:    100001,
			Message: "输入参数有误",
		}
		ctrl.ServeJSON()
	}

	service.NewTaskInfoService().UpdateStatus(req.Id, req.Status)
	ctrl.Data["json"] = response.CreateTaskResponse{
		Code:    0,
		Message: "成功",
	}
	ctrl.ServeJSON()
}

func (ctrl *TaskController) DeleteTask() {
	var req request.DeleteTaskRequest
	if err := json.Unmarshal(ctrl.Ctx.Input.RequestBody, &req); err != nil {
		ctrl.Data["json"] = response.CreateTaskResponse{
			Code:    100001,
			Message: "输入参数有误",
		}
		ctrl.ServeJSON()
	}

	service.NewTaskInfoService().DeleteInfo(req.Id)
	ctrl.Data["json"] = response.CreateTaskResponse{
		Code:    0,
		Message: "成功",
	}
	ctrl.ServeJSON()
}
