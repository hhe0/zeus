package controller

import (
	"zeus/common/controller"
	"zeus/common/util"
)

type EmailController struct {
	controller.BaseController
}

func (ctrl *EmailController) SendEmail() {
	util.SendMail(util.Mail{
		To:       "hhe_minc@126.com",
		Subject:  "找回密码",
		Body:     "测试发送邮件",
		MailType: "html",
	})

	ctrl.Success()
}
