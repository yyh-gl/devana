package common

import (
	"time"

	"github.com/yut-kt/goholiday"
)

func CountBusinessDay(since, until time.Time) int {
	hours := until.Sub(since).Hours()
	days := int(hours/24 + 1)

	count := 0
	for i := 0; i < days; i++ {
		datetime := since.Add(time.Duration(i) * 24 * time.Hour)
		if goholiday.IsBusinessDay(datetime) {
			count += 1
		}
	}
	return count
}

func ConvertToSinceDatetime(src string) time.Time {
	layout := "2006-01-02"
	t, _ := time.Parse(layout, src)
	t = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)
	return t
}

func ConvertToUntilDatetime(src string) time.Time {
	layout := "2006-01-02"
	t, _ := time.Parse(layout, src)
	t = time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 999999999, time.Local)
	return t
}
