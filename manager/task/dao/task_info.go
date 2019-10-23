package dao

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"zeus/common/util"
	"zeus/manager/task/enum"
	"zeus/manager/task/model"
)

type TaskInfoDao interface {
	InsertInfo(info model.TaskInfo) int
	GetTodayList(status int) (int, []model.TaskInfo)
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
		Status:  enum.TaskStatusNo,
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

func (dao taskInfoDao) GetTodayList(status int) (int, []model.TaskInfo) {
	var (
		taskList  []model.TaskInfo
		taskTotal int64
		err       error
	)

	// TODO: 待优化
	if status == enum.TaskStatusAll {
		taskTotal, err = dao.DB.QueryTable(model.TaskInfo{}.TableName()).All(&taskList)
		if err != nil {
			// TODO: 记录日志
			fmt.Println(err)
		}
	} else {
		taskTotal, err = dao.DB.QueryTable(model.TaskInfo{}.TableName()).Filter("status", status).All(&taskList)
		if err != nil {
			// TODO: 记录日志
			fmt.Println(err)
		}
	}

	return int(taskTotal), taskList
}
