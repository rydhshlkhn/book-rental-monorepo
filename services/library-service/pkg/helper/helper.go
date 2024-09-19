package helper

import (
	"fmt"
	"time"
)

func TruncateToDate(t time.Time) time.Time {
	return t.Truncate(24 * time.Hour)
}

func ParseDate(dateStr string) (time.Time, error) {
	// Define the layout corresponding to the input format
	layout := "2006-01-02 15:04:05"

	// Parse the date string
	parsedTime, err := time.Parse(layout, dateStr)
	if err != nil {
		return time.Time{}, fmt.Errorf("parsing date: %v", err)
	}

	return parsedTime, nil
}
