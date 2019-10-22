package models

type TODOList struct {
	Id         int    `json:"id"`
	UserId     int    `json:"user_id"`
	Content    string `json:"content"`
	FinishTime int    `json:"finish_time"`
	CreateTime int    `json:"create_time"`
	UpdateTime int    `json:"update_time"`
	IsDeleted  int    `json:"is_deleted"`
}

func (m *TODOList) TableName() string {
	return TableName("todo_list")
}
