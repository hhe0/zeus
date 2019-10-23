package enum

const (
	TaskStatusAll = iota // 所有的状态
	TaskStatusNo         // 未完成
	TaskStatusYes        // 已完成
)

func GetAllStatus() []int {
	return []int{
		TaskStatusNo,
		TaskStatusYes,
	}
}
