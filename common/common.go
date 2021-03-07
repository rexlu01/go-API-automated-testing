package common

import (
	"os"
	"time"
)

// 时间差，纳秒
func DiffNano(startTime time.Time) (diff int64) {
	diff = int64(time.Since(startTime))
	return
}

//获取当前目录
func GetLocalPath() (pwd string) {
	str, err := os.Getwd()
	if err != nil {
		return ""
	}

	return str
}
