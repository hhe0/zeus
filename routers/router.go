package routers

import (
	"github.com/astaxie/beego"
	"zeus/controllers/index"
)

func init() {
	beego.Router("/", &index.Controller{})
}
