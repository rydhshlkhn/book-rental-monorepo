package helper

import "time"

func TruncateToDate(t time.Time) time.Time {
	return t.Truncate(24 * time.Hour)
}
