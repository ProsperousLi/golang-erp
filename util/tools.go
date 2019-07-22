package util

import (
	"time"

	"erpweb/logs"

	"github.com/satori/go.uuid"
)

//时间戳格式化
func Str2TimeStamp(param string) int64 {
	timeLayout := "2006-01-02 15:04:05"
	the_time, err := time.ParseInLocation(timeLayout, param, time.Local)
	if err != nil {
		logs.FileLogs.Error("Str2TimeStamp failed : %s", err)
	}
	unix_time2 := the_time.Unix()

	return unix_time2
}

//时间戳格式化
func Str2Time(param string) time.Time {
	timeLayout := "2006-01-02 15:04:05"
	the_time, err := time.ParseInLocation(timeLayout, param, time.Local)
	if err != nil {
		logs.FileLogs.Error("Str2TimeStamp failed : %s", err)
	}

	return the_time
}

//uuid
func GetToken() string {
	// 创建
	u1, _ := uuid.NewV4()
	logs.FileLogs.Info("UUIDv4: %s\n", u1)

	return u1.String()
}
