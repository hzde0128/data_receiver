package datetool

import (
	"time"
)

const (
	LAYOUT_TIME = "2006-01-02 15:04:05"
	LAYOUT_DATE = "2006-01-02"
)

func GetCurrentTime() string {
	return time.Now().Format(LAYOUT_TIME)
}

func GetCurrentDate() string {
	return time.Now().Format(LAYOUT_DATE)
}
