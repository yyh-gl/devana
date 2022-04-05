package common

import (
	"context"
	"time"
)

type Analyzer interface {
	Name() string
	Do(context.Context) (Records, error)
}

type Conditions struct {
	DevelopmentMemberNum int
	Since                time.Time
	Until                time.Time
}

func NewConditions(devMemberNum int, since, until time.Time) *Conditions {
	return &Conditions{
		DevelopmentMemberNum: devMemberNum,
		Since:                since,
		Until:                until,
	}
}
