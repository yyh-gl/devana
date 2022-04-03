package ddd

import (
	"github.com/yyh-gl/devana/common"
)

// 理想値
const ideal float64 = 0.1

type DDDAnalyzer struct {
	repo       *common.GitRepository
	conditions common.Conditions
}

func NewDDDAnalyzer(repo *common.GitRepository, cond common.Conditions) common.Analyzer {
	return &DDDAnalyzer{repo: repo, conditions: cond}
}

func (d DDDAnalyzer) Name() string {
	return "d/d/d"
}

func (d DDDAnalyzer) Do() (common.Records, error) {
	tags, err := d.repo.FetchTags(d.conditions.Since, d.conditions.Until)
	if err != nil {
		return nil, err
	}

	deployTotalCount := float64(len(tags))
	businessDayCount := float64(common.CountBusinessDay(d.conditions.Since, d.conditions.Until))
	deployCountPerDay := deployTotalCount / businessDayCount

	records := common.Records{
		{"ideal", common.ConvertToString(ideal)},
		{"result", common.ConvertToString(deployCountPerDay / float64(d.conditions.DevelopmentMemberNum))},
	}
	return records, nil
}
