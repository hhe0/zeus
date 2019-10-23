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
