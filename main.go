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

	if err := ddd.Do(repo); err != nil {
		log.Fatal(err)
	}
}
