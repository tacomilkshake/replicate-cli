package tables

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

func CreateTable() *tablewriter.Table {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	return table
}
