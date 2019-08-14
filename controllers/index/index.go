package index

import (
	"github.com/astaxie/beego"
	"zeus/common/apiresponse"
	"zeus/controllers/index/response"
)

type Index struct {
	beego.Controller
}

func (c *Index) Get() {
	c.Data["json"] = apiresponse.SuccessResponse{
		Data: response.GetIndexResponse{
			Message: "Hello, Beego!",
		},
	}

	c.ServeJSON()
}
