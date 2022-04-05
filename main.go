package main

import (
	"context"
	"fmt"
	"github.com/yyh-gl/devana/ddd"
	"github.com/yyh-gl/devana/pr_lead_time"
	"os"

	"github.com/yyh-gl/devana/common"
)

func main() {
	ctx := context.Background()

	// TODO: トークンの受け取り方を変更
	if len(os.Args) < 4 {
		msg := fmt.Sprintf(`以下のとおり引数を指定してください。
$ devana <url> <since> <until> <token>
=> url:   調査対象のGitリポジトリURL【必須】
   since: 調査対象期間の開始日（例：2022-04-01）【必須】
   until: 調査対象期間の開始日（例：2022-09-30）【必須】
   token: プライベートリポジトリアクセス用のトークン（Personal access tokens）【任意】
`)
		fmt.Println(msg)
		os.Exit(1)
	}
	url := os.Args[1]
	since := common.ConvertToSinceDatetime(os.Args[2])
	until := common.ConvertToUntilDatetime(os.Args[3])
	var token string
	if len(os.Args) >= 5 {
		token = os.Args[4]
	}

	gitClient, err := common.NewGitClient(url, token)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	gitHubClient, err := common.NewGitHubClient(ctx, url, token)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	cond := common.NewConditions(3, since, until)

	fmt.Println(fmt.Sprintf("調査期間: %v ~ %v", since, until))
	analyzers := []common.Analyzer{
		ddd.NewAnalyzer(gitClient, *cond),
		pr_lead_time.NewAnalyzer(gitHubClient, *cond),
	}
	for _, a := range analyzers {
		records, err := a.Do(ctx)
		if err != nil {
			fmt.Println(err)
		}

		common.OutputResult(a.Name(), records)
	}
}
