package service

import (
	"zeus/manager/task/dao"
	"zeus/manager/task/model"
)

type TaskInfoService interface {
	InsertInfo(info model.TaskInfo) int
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
