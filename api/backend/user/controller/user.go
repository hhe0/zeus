package controller

import (
	"zeus/api/backend/user/response"
	"zeus/common/controller"
	"zeus/common/http"
)

type UserController struct {
	controller.BaseController
}

func (ctrl *UserController) UserLogin() {
	ctrl.Data["json"] = http.APIResponse{
		Code: 20000,
		Data: "admin-token",
	}
	ctrl.ServeJSON()
}

func (ctrl *UserController) GetUserInfo() {
	//roles, name, avatar, introduction
	ctrl.Data["json"] = http.APIResponse{
		Code: 20000,
		Data: response.GetUserInfoResponse{
			Roles:        "admin",
			Introduction: "I am an editor",
			Avatar:       "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif",
			Name:         "Normal Editor",
		},
	}
	ctrl.ServeJSON()
}
