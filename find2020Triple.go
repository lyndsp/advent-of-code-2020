package integers

func indexOfTripleToSum(sum int, input []int) (int, int, int) {
	for index := 0; index < len(input); index++ {
		current := input[index]
		theRest := input[index+1:]
		var j, k = indexOfFirstToSum(sum-current, theRest)

		if j != -1 && k != -1 {
			return index, index + j + 1, index + k + 1
		}
	}
	return -1, -1, -1
}
