package controllers

import (
	"github.com/astaxie/beego"
)

type HelloController struct {
	beego.Controller
}

type person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type JsonResponse struct {
	Data []person `json:"data"`
}

func (c *HelloController) Get() {
	var jsonResponse JsonResponse

	var persons []person
	persons = append(persons, person{"awks", 23})
	persons = append(persons, person{"ruby", 25})
	jsonResponse.Data = persons
	c.Data["json"] = &jsonResponse

	c.ServeJSON()
}
