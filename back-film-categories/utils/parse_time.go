package utils

import "time"

func ParseTime(t string) time.Time {
	parsedTime, err := time.Parse(time.RFC3339, t)
	if err != nil {
		return time.Time{}
	}
	return parsedTime
}
