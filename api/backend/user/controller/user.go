package controller

import (
	"zeus/common/controller"
	"zeus/common/http"
)

type UserController struct {
	controller.BaseController
}

func (ctrl *UserController) UserLogin() {
	ctrl.Data["json"] = http.APIResponse{
		Code: 20000,
		Data: "212341212",
	}
	ctrl.ServeJSON()
}
