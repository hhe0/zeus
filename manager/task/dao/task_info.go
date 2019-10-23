package dao

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	enumCommon "zeus/common/enum"
	"zeus/common/util"
	"zeus/manager/task/enum"
	"zeus/manager/task/model"
)

type TaskInfoDao interface {
	InsertInfo(info model.TaskInfo) int
	GetTodayList(status int) (int, []model.TaskInfo)
	UpdateStatus(id, status int)
	DeleteInfo(id int)
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

	todayMinTimestamp := util.GetTodayMinTimestamp()
	todayMaxTimstamp := util.GetTodayMaxTimestamp()

	// TODO: 待优化
	if status == enum.TaskStatusAll {
		taskTotal, err = dao.DB.QueryTable(model.TaskInfo{}.TableName()).Filter("create_time__gt ", todayMinTimestamp).Filter("create_time__lt", todayMaxTimstamp).Filter("is_deleted", 0).All(&taskList)
		if err != nil {
			// TODO: 记录日志
			fmt.Println(err)
		}
	} else {
		taskTotal, err = dao.DB.QueryTable(model.TaskInfo{}.TableName()).Filter("create_time__gt", todayMinTimestamp).Filter("create_time__lt", todayMaxTimstamp).Filter("status", status).Filter("is_deleted", 0).All(&taskList)
		if err != nil {
			// TODO: 记录日志
			fmt.Println(err)
		}
	}

	return int(taskTotal), taskList
}

func (dao taskInfoDao) UpdateStatus(id, status int) {
	var taskInfo model.TaskInfo

	taskInfo.Id = id
	if dao.DB.Read(&taskInfo) == nil {
		taskInfo.Status = status
		if _, err := dao.DB.Update(&taskInfo); err == nil {
			fmt.Println(err)
		}
	}
}

func (dao taskInfoDao) DeleteInfo(id int) {
	var taskInfo model.TaskInfo

	taskInfo.Id = id
	if dao.DB.Read(&taskInfo) == nil {
		taskInfo.IsDeleted = enumCommon.IsDeletedYes
		if _, err := dao.DB.Update(&taskInfo); err == nil {
			fmt.Println(err)
		}
	}
}