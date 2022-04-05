package pr_lead_time

import (
	"context"
	"github.com/yyh-gl/devana/common"
	"sort"
)

// 理想値
const ideal float64 = 0.1

type Analyzer struct {
	client     *common.GitHubClient
	conditions common.Conditions
}

func NewAnalyzer(client *common.GitHubClient, cond common.Conditions) common.Analyzer {
	return &Analyzer{client: client, conditions: cond}
}

func (a Analyzer) Name() string {
	return "PR Lead Time"
}

func (a Analyzer) Do(ctx context.Context) (common.Records, error) {
	prs, err := a.client.FetchPRs(ctx, a.conditions.Since, a.conditions.Until)
	if err != nil {
		return nil, err
	}

	takenDayTotal := 0.0
	takenDays := make([]float64, len(prs))
	for i, pr := range prs {
		takenDay := pr.ClosedAt.Sub(pr.CreatedAt).Seconds() / 86400
		takenDayTotal += takenDay
		takenDays[i] = takenDay
	}
	sort.Float64s(takenDays)

	records := common.Records{
		{"ave", common.ConvertToString(takenDayTotal / float64(len(prs)))},
		{"med", common.ConvertToString(takenDays[len(takenDays)/2])},
	}
	return records, nil
}
