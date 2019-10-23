package util

import "time"

// 获取当前时间戳 - int64
func GetCurrentTimestampInt64() int64 {
	return time.Now().Unix()
}

// 获取当前时间戳 - int
func GetCurrentTimestampInt() int {
	return int(time.Now().Unix())
}

// 获取当前日期 - time.Time
func GetCurrentTime() time.Time {
	return time.Now()
}

// 获取当天的最小时间戳
func GetTodayMinTimestamp() int {
	t := time.Now()
	year, month, day := t.Date()

	return int(time.Date(year, month, day, 0, 0, 0, 0, t.Location()).Unix())
}

// 获取当天的最大时间戳
func GetTodayMaxTimestamp() int {
	t := time.Now()
	year, month, day := t.Date()

	return int(time.Date(year, month, day+1, 0, 0, 0, 0, t.Location()).Unix() - 1)
}
