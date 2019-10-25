package dao

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"log"
	enumCommon "zeus/common/enum"
	"zeus/common/util"
	"zeus/manager/task/enum"
	"zeus/manager/task/model"
)

type TaskInfoDao interface {
	InsertInfo(userId int, content string) int
	GetTodayList(status int) (int, []model.TaskInfo)
	UpdateStatus(id, status int)
	DeleteInfo(id int)
}

type taskInfoDao struct {
	DB     orm.Ormer
	Logger *log.Logger
}

func NewTaskInfoDao() TaskInfoDao {
	return &taskInfoDao{
		DB:     orm.NewOrm(),
		Logger: logs.GetLogger("db"),
	}
}

func (dao taskInfoDao) InsertInfo(userId int, content string) int {
	taskInfo := model.TaskInfo{
		UserId:  userId,
		Content: content,
		Status:  enum.TaskStatusNo,
	}
	taskInfo.CreateTime = util.GetCurrentTimestampInt()
	taskInfo.UpdateTime = util.GetCurrentTimestampInt()

	insertId, err := dao.DB.Insert(&taskInfo)
	if err != nil {
		dao.Logger.Println(err)
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
			dao.Logger.Println(err)
		}
	} else {
		taskTotal, err = dao.DB.QueryTable(model.TaskInfo{}.TableName()).Filter("create_time__gt", todayMinTimestamp).Filter("create_time__lt", todayMaxTimstamp).Filter("status", status).Filter("is_deleted", 0).All(&taskList)
		if err != nil {
			dao.Logger.Println(err)
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
			dao.Logger.Println(err)
		}
	}
}

func (dao taskInfoDao) DeleteInfo(id int) {
	var taskInfo model.TaskInfo

	taskInfo.Id = id
	if dao.DB.Read(&taskInfo) == nil {
		taskInfo.IsDeleted = enumCommon.IsDeletedYes
		if _, err := dao.DB.Update(&taskInfo); err == nil {
			dao.Logger.Println(err)
		}
	}
}
