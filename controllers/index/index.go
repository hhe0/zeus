package index

import (
	"github.com/astaxie/beego"
	"zeus/common/apiresponse"
	"zeus/controllers/index/response"
)

type Controller struct {
	beego.Controller
}

func (c *Controller) Get() {
	c.Data["json"] = apiresponse.SuccessResponse{
		Data: response.GetIndexResponse{
			Message: "Hello, Beego!",
		},
	}

	c.ServeJSON()
}
