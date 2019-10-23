package response

import "zeus/manager/task/model"

type CreateTaskResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type GetTaskListResponse struct {
	Code    int             `json:"code"`
	Message string          `json:"message"`
	Data    GetTaskListData `json:"data"`
}

type GetTaskListData struct {
	Total    int              `json:"total"`
	TaskList []model.TaskInfo `json:"task_list"`
}
