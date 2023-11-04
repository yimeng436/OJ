package utils

import "time"

func TimeToString(tem time.Time) string {
	return tem.Format("2006-01-02 15:04:05")
}
