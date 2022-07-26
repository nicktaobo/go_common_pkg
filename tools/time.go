package tools

import (
	"time"
)

func GetTodayString() string {
	return GetDateFormat(time.Now(), "2006-01-02")
}

func GetDateFormat(time time.Time, format string) string {
	return time.Format(format)
}
