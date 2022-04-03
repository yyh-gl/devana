package ddd

import (
	"fmt"

	"github.com/yyh-gl/devana/common"
)

const (
	analyticsMethodName = "d/d/d"

	// 理想値
	ideal float64 = 1

	// TODO: 特定条件に一致するコミットの数を設定
	deployNumPerDay float64 = 1
	// TODO: 手動打ち込み
	devMemberNum float64 = 3
)

func Do(repo *common.GitRepository) error {
	logs, err := repo.FetchLogs()
	if err != nil {
		return err
	}

	for _, l := range logs {
		fmt.Println("========================")
		fmt.Printf("%+v\n", l)
		fmt.Println("========================")
	}

	records := common.Records{
		{"ideal", common.ConvertToString(ideal)},
		{"result", common.ConvertToString(deployNumPerDay / devMemberNum)},
	}
	common.OutputResult(analyticsMethodName, records)
	return nil
}
