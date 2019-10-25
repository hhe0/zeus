package cache

import (
	"zeus/manager/task/dao"
	"zeus/manager/task/model"
)

type TaskInfoCache interface {
	InsertInfo(userId int, content string) int
	GetTodayList(status int) (int, []model.TaskInfo)
	UpdateStatus(id, status int)
	DeleteInfo(id int)
}

type taskInfoCache struct {
	TaskInfoDao dao.TaskInfoDao
}

func NewTaskInfoCache() TaskInfoCache {
	return &taskInfoCache{
		TaskInfoDao: dao.NewTaskInfoDao(),
	}
}

func (cache taskInfoCache) InsertInfo(userId int, content string) int {
	return cache.TaskInfoDao.InsertInfo(userId, content)
}

func (cache taskInfoCache) GetTodayList(status int) (int, []model.TaskInfo) {
	return cache.TaskInfoDao.GetTodayList(status)
}

func (cache taskInfoCache) UpdateStatus(id, status int) {
	cache.TaskInfoDao.UpdateStatus(id, status)
}

func (cache taskInfoCache) DeleteInfo(id int) {
	cache.TaskInfoDao.DeleteInfo(id)
}
