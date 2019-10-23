package model

import (
	"github.com/astaxie/beego/orm"
	"zeus/manager/common/model"
)

type TaskInfo struct {
	model.BaseModel
	UserId     int    `json:"user_id"`
	Content    string `json:"content"`
	FinishTime int    `json:"finish_time"`
}

func init() {
	model.Init()
	orm.RegisterModel(new(TaskInfo))
}

func (TaskInfo) TableName() string {
	return model.CompletedTableName("task_info")
}
