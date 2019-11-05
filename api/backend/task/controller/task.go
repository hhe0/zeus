package controller

import (
	"encoding/json"
	"github.com/astaxie/beego/validation"
	"zeus/api/backend/task/request"
	"zeus/api/backend/task/response"
	"zeus/common/apicode"
	"zeus/common/controller"
	"zeus/manager/task/service"
)

type TaskController struct {
	controller.BaseController
}

func (ctrl *TaskController) CreateTask() {
	var (
		req       request.CreateTaskRequest
		validator validation.Validation
	)
	_ = json.Unmarshal(ctrl.Ctx.Input.RequestBody, &req)
	if checkPass, err := validator.Valid(&req); err != nil || !checkPass {
		ctrl.Error(apicode.InputError)
	}

	// TODO: user_id 需要在 token 中进行获取
	//service.NewTaskInfoService().InsertInfo(1, req.Content)

	ctrl.Success()
}

func (ctrl *TaskController) GetTaskList() {
	status, err := ctrl.GetInt("status")
	if err != nil {
		ctrl.Error(apicode.InputError)
	}

	total, data := service.NewTaskInfoService().GetTodayList(status)

	ctrl.Success(response.GetTaskListResponse{
		Total:    total,
		TaskList: data,
	})
}

func (ctrl *TaskController) UpdateTaskStatus() {
	var req request.UpdateTaskStatusRequest
	if err := json.Unmarshal(ctrl.Ctx.Input.RequestBody, &req); err != nil {
		ctrl.Error(apicode.InputError)
	}

	service.NewTaskInfoService().UpdateStatus(req.Id, req.Status)

	ctrl.Success()
}

func (ctrl *TaskController) DeleteTask() {
	var req request.DeleteTaskRequest
	if err := json.Unmarshal(ctrl.Ctx.Input.RequestBody, &req); err != nil {
		ctrl.Error(apicode.InputError)
	}

	service.NewTaskInfoService().DeleteInfo(req.Id)

	ctrl.Success()
}
