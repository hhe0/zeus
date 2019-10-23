package service

import (
	"zeus/manager/task/cache"
	"zeus/manager/task/model"
)

type TaskInfoService interface {
	InsertInfo(info model.TaskInfo) int
	GetTodayList(status int) (int, []model.TaskInfo)
	UpdateStatus(id, status int)
	DeleteInfo(id int)
}

type taskInfoService struct {
	TaskInfoCache cache.TaskInfoCache
}

func NewTaskInfoService() TaskInfoService {
	return &taskInfoService{
		TaskInfoCache: cache.NewTaskInfoCache(),
	}
}

func (svc taskInfoService) InsertInfo(info model.TaskInfo) int {
	return svc.TaskInfoCache.InsertInfo(info)
}

func (svc taskInfoService) GetTodayList(status int) (int, []model.TaskInfo) {
	return svc.TaskInfoCache.GetTodayList(status)
}

func (svc taskInfoService) UpdateStatus(id, status int) {
	svc.TaskInfoCache.UpdateStatus(id, status)
}

func (svc taskInfoService) DeleteInfo(id int) {
	svc.TaskInfoCache.DeleteInfo(id)
}
