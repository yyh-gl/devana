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
