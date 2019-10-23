package dao

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"zeus/common/util"
	"zeus/manager/task/model"
)

type TaskInfoDao interface {
	InsertInfo(info model.TaskInfo) int
}

type taskInfoDao struct {
	DB orm.Ormer
}

func NewTaskInfoDao() TaskInfoDao {
	return &taskInfoDao{
		DB: orm.NewOrm(),
	}
}

func (dao taskInfoDao) InsertInfo(info model.TaskInfo) int {
	taskInfo := model.TaskInfo{
		UserId:  info.UserId,
		Content: info.Content,
	}
	taskInfo.CreateTime = util.GetCurrentTimestampInt()
	taskInfo.UpdateTime = util.GetCurrentTimestampInt()

	insertId, err := dao.DB.Insert(&taskInfo)
	if err != nil {
		// TODO: 记录日志
		fmt.Println(err)
	}

	return int(insertId)
}
