package util

import (
	"time"
)

// FormatDatetime 时间格式化
func FormatDatetime(timestamp int64) string {
	return time.UnixMilli(timestamp).Format(time.DateTime)
}

func ParseDatetime(datetime string, pattern string) (time.Time, error) {
	return time.Parse(pattern, datetime)
}
