package controller

import (
	"github.com/astaxie/beego"
	"zeus/common/apicode"
	"zeus/common/http"
)

type BaseController struct {
	beego.Controller
}

func (ctrl *BaseController) Success(data ...interface{}) {
	ctrl.Data["json"] = http.APIResponse{
		Code:    apicode.APISuccessCode.Code,
		Message: apicode.APISuccessCode.Message,
		Data:    data,
	}
	ctrl.ServeJSON()
}

func (ctrl *BaseController) Error(code apicode.APICode, data ...interface{}) {
	ctrl.Data["json"] = http.APIResponse{
		Code:    code.Code,
		Message: code.Message,
		Data:    data,
	}
	ctrl.ServeJSON()
}
