package controller

import (
	"zeus/common/controller"
	"zeus/common/util"
)

type EmailController struct {
	controller.BaseController
}

func (ctrl *EmailController) SendEmail() {
	//var (
	//	req       request.CreateTaskRequest
	//	validator validation.Validation
	//)
	//_ = json.Unmarshal(ctrl.Ctx.Input.RequestBody, &req)
	//if checkPass, err := validator.Valid(&req); err != nil || !checkPass {
	//	ctrl.Error(apicode.InputError)
	//}

	err := util.SendMail("hhe_minc@126.com", "找回密码", "测试发送邮件", "")
	if err != nil {
		return
	}

	ctrl.Success()
}
