package common

import "time"

type Analyzer interface {
	Name() string
	Do() (Records, error)
}

type Conditions struct {
	DevelopmentMemberNum int
	Since                time.Time
	Until                time.Time
}

func NewConditions(devMemberNum int, since, until *time.Time) *Conditions {
	if since == nil {
		t := time.Now().Add(-365 * 24 * time.Hour)
		since = &t
	}
	if until == nil {
		t := time.Now()
		until = &t
	}

	return &Conditions{
		DevelopmentMemberNum: devMemberNum,
		Since:                *since,
		Until:                *until,
	}
}
