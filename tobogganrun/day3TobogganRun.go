package tobogganrun

// CountTrees through the route
func CountTrees(right int, down int, mountainMap []string) int {
	treeCount := 0
	rightCounter := 1
	rowIndex := down

	for rowIndex < len(mountainMap) {
		currentRow := mountainMap[rowIndex]
		columnIndex := right * rightCounter

		for columnIndex >= len(currentRow) {
			currentRow += currentRow
		}

		if currentRow[columnIndex] == '#' {
			treeCount++
		}

		rowIndex += down
		rightCounter++
	}

	return treeCount
}
