package ddd

import (
	"github.com/yyh-gl/devana/common"
)

const (
	analyticsMethodName = "d/d/d"

	// 理想値
	ideal float64 = 0.1
)

type DDDAnalyzer struct {
	repo       *common.GitRepository
	conditions common.Conditions
}

func NewDDDAnalyzer(repo *common.GitRepository, cond common.Conditions) common.Analyzer {
	return &DDDAnalyzer{repo: repo, conditions: cond}
}

func (d DDDAnalyzer) Do() error {
	tags, err := d.repo.FetchTags(d.conditions.Since, d.conditions.Until)
	if err != nil {
		return err
	}

	deployTotalCount := float64(len(tags))
	businessDayCount := float64(common.CountBusinessDay(d.conditions.Since, d.conditions.Until))
	deployCountPerDay := deployTotalCount / businessDayCount

	records := common.Records{
		{"ideal", common.ConvertToString(ideal)},
		{"result", common.ConvertToString(deployCountPerDay / float64(d.conditions.DevelopmentMemberNum))},
	}
	common.OutputResult(analyticsMethodName, records)
	return nil
}
