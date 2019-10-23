package models

import "github.com/astaxie/beego/orm"

type TaskInfo struct {
	Id         int    `json:"id"`
	UserId     int    `json:"user_id"`
	Content    string `json:"content"`
	FinishTime int    `json:"finish_time"`
	CreateTime int    `json:"create_time"`
	UpdateTime int    `json:"update_time"`
	IsDeleted  int    `json:"is_deleted"`
}

func init() {
	Init()
	orm.RegisterModel(new(TaskInfo))

}

func (TaskInfo) TableName() string {
	return TableName("task_info")
}
