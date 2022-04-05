package ddd

import (
	"context"
	"github.com/yyh-gl/devana/common"
)

// 理想値
const ideal float64 = 0.1

type Analyzer struct {
	client     *common.GitClient
	conditions common.Conditions
}

func NewAnalyzer(client *common.GitClient, cond common.Conditions) common.Analyzer {
	return &Analyzer{client: client, conditions: cond}
}

func (a Analyzer) Name() string {
	return "d/d/d"
}

func (a Analyzer) Do(_ context.Context) (common.Records, error) {
	tags, err := a.client.FetchTags(a.conditions.Since, a.conditions.Until)
	if err != nil {
		return nil, err
	}

	deployTotalCount := float64(len(tags))
	businessDayCount := float64(common.CountBusinessDay(a.conditions.Since, a.conditions.Until))
	deployCountPerDay := deployTotalCount / businessDayCount

	records := common.Records{
		{"ideal", common.ConvertToString(ideal)},
		{"result", common.ConvertToString(deployCountPerDay / float64(a.conditions.DevelopmentMemberNum))},
	}
	return records, nil
}
