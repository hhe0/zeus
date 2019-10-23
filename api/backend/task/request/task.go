package request

type CreateTaskRequest struct {
	Content string `json:"content"`
}

type GetTaskListRequest struct {
	Status int `form:"status"`
}
