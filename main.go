package main

import (
	"log"

	"github.com/yyh-gl/devana/common"
	"github.com/yyh-gl/devana/ddd"
)

func main() {
	repo, err := common.NewGitRepository("https://github.com/yyh-gl/devana")
	if err != nil {
		log.Fatal(err)
	}

	//t1 := time.Date(2021, 4, 3, 0, 0, 0, 0, time.Local)
	//t2 := time.Date(2022, 4, 3, 23, 59, 59, 99999, time.Local)
	cond := common.NewConditions(3, nil, nil)

	ddd := ddd.NewDDDAnalyzer(repo, *cond)
	if err := ddd.Do(); err != nil {
		log.Fatal(err)
	}
}
