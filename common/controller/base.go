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
	var response http.APIResponse
	response = http.APIResponse{
		Code:    apicode.APISuccessCode.Code,
		Message: apicode.APISuccessCode.Message,
	}

	if len(data) > 0 {
		response.Data = data[0]
	} else {
		response.Data = map[string]interface{}{}
	}

	ctrl.Data["json"] = response
	ctrl.ServeJSON()
}

func (ctrl *BaseController) Error(code apicode.APICode) {
	ctrl.Data["json"] = http.APIResponse{
		Code:    code.Code,
		Message: code.Message,
		Data:    map[string]interface{}{},
	}
	ctrl.ServeJSON()
}
