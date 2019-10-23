package service

import (
	"zeus/manager/task/dao"
	"zeus/manager/task/model"
)

type TaskInfoService interface {
	InsertInfo(info model.TaskInfo) int
	GetTodayList(status int) (int, []model.TaskInfo)
}

type taskInfoService struct {
	TaskInfoDao dao.TaskInfoDao
}

func NewTaskInfoService() TaskInfoService {
	return &taskInfoService{
		TaskInfoDao: dao.NewTaskInfoDao(),
	}
}

func (svc taskInfoService) InsertInfo(info model.TaskInfo) int {
	return svc.TaskInfoDao.InsertInfo(info)
}

func (svc taskInfoService) GetTodayList(status int) (int, []model.TaskInfo) {
	return svc.TaskInfoDao.GetTodayList(status)
}
