package routers

import (
	"github.com/astaxie/beego"
	"zeus/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
