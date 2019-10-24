package controller

import (
	"github.com/astaxie/beego"
)

type TestController struct {
	beego.Controller
}

type TestHelloWorldResponse struct {
	Code    int
	Message string
	Data    interface{}
}

func (ctrl *TestController) TestHelloWorld() {
	ctrl.Data["json"] = TestHelloWorldResponse{
		Code:    0,
		Message: "成功",
		Data:    "Hello World",
	}
	ctrl.ServeJSON()
}
