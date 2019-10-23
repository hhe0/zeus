package request

type CreateTaskRequest struct {
	Content string `json:"content"`
}

type GetTaskListRequest struct {
	Status int `form:"status"`
}

type UpdateTaskStatusRequest struct {
	Id     int `json:"id"`
	Status int `json:"status"`
}

type DeleteTaskRequest struct {
	Id int `json:"id"`
}
