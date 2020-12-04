package integers

// Find2020Rhs returns true if the lhs and an element in the rhs add to 2020
func indexOfFirstToSum(sum int, input []int) (int, int) {
	for index := 0; index < len(input); index++ {
		current := input[index]
		theRest := input[index+1:]
		for i, v := range theRest {
			if current+v == sum {
				return index, i + index + 1
			}
		}
	}
	return -1, -1
}

// Find2020Pair returns true if two elements in the input add to 2020
// func Find2020Pair(input []int) bool {

// 	return false
// }
