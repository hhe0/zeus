package routers

import (
	"github.com/astaxie/beego"
	"zeus/controllers/backend"
)

func init() {
	// TODOList
	beego.Router("/", &backend.TODOListController{})
}
