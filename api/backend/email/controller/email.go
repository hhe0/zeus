package controller

import (
	"encoding/json"
	"zeus/api/backend/task/request"
	"zeus/common/apicode"
	"zeus/common/controller"
	"zeus/manager/task/service"
)

type EmailController struct {
	controller.BaseController
}

func (ctrl *EmailController) SendEmail() {
	var req request.UpdateTaskStatusRequest
	if err := json.Unmarshal(ctrl.Ctx.Input.RequestBody, &req); err != nil {
		ctrl.Error(apicode.InputError)
	}

	service.NewTaskInfoService().UpdateStatus(req.Id, req.Status)

	ctrl.Success()
}
