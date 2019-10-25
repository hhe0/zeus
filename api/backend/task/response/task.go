package response

import "zeus/manager/task/model"

type GetTaskListResponse struct {
	Total    int              `json:"total"`
	TaskList []model.TaskInfo `json:"task_list"`
}
