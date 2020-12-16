package tobogganrun

// CountTrees through the route
func CountTrees(mountainMap []string) int {
	treeCount := 0
	columnIndex := 0

	for rowIndex, row := range mountainMap {
		if rowIndex == 0 {
			continue
		}

		columnIndex = 3 * rowIndex

		for columnIndex > len(row) {
			row += row
		}

		if row[columnIndex] == '#' {
			treeCount++
		}
	}

	return treeCount
}
