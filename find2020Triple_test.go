package integers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_when_no_match_triples_are_negative(t *testing.T) {

	i, j, k := indexOfTripleToSum(10, []int{1, 2, 3})

	assert.Equal(t, -1, i)
	assert.Equal(t, -1, j)
	assert.Equal(t, -1, k)
}

func Test_when_sum_is_found_three_indices_returned(t *testing.T) {

	testData := []struct {
		name                 string
		sum                  int
		slice                []int
		first, second, third int
	}{
		{"first plus second plus third ", 6, []int{1, 2, 3}, 0, 1, 2},
		{"first plus third plus fourth ", 8, []int{1, 2, 3, 4}, 0, 2, 3},
		{"second plus third plus fourth ", 9, []int{1, 2, 3, 4}, 1, 2, 3},
	}

	for _, td := range testData {
		t.Run(td.name, func(t *testing.T) {
			i, j, k := indexOfTripleToSum(td.sum, td.slice)

			assert.Equal(t, td.first, i)
			assert.Equal(t, td.second, j)
			assert.Equal(t, td.third, k)
		})
	}
}

// func Test_the_triple_answer(t *testing.T) {
// 	i, j, k := indexOfTripleToSum(2020, adventTestValues)
// 	assert.Equal(t, 0, adventTestValues[i]*adventTestValues[j]*adventTestValues[k])
// }
