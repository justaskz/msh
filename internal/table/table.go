package table

import "fmt"

func getColumnWidths(rows [][]string) []int {
	if len(rows) == 0 {
		return nil
	}
	cols := len(rows[0])
	widths := make([]int, cols)

	for _, row := range rows {
		for i, col := range row {
			if len(col) > widths[i] {
				widths[i] = len(col)
			}
		}
	}
	return widths
}

func PrintAlignedTable(rows [][]string) {
	colWidths := getColumnWidths(rows)

	for _, row := range rows {
		for i, col := range row {
			fmt.Printf("%-*s ", colWidths[i], col)
		}
		fmt.Println()
	}
}
