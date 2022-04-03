package common

import (
	"fmt"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
)

type Records [][]string

func OutputResult(analyticsMethodName string, records Records) {
	fmt.Printf("▼ Analytics result by %s ▼\n", analyticsMethodName)

	table := tablewriter.NewWriter(os.Stdout)
	for _, r := range records {
		table.Append(r)
	}
	table.Render()
}

// TODO: float64以外も受け取れるようにする
func ConvertToString(val float64) string {
	return strconv.FormatFloat(val, 'f', 5, 64)
}
