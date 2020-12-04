package integers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_when_no_match_nothing_is_found(t *testing.T) {

	i, j := indexOfFirstToSum(2020, []int{1, 2, 3})

	assert.Equal(t, -1, i)
	assert.Equal(t, -1, j)
}

func Test_when_sum_is_found_indices_returned(t *testing.T) {

	testData := []struct {
		name          string
		sum           int
		slice         []int
		first, second int
	}{
		{"first plus second", 3, []int{1, 2, 3}, 0, 1},
		{"second plus third", 5, []int{1, 2, 3}, 1, 2},
		{"non-consecutive indexes", 4, []int{1, 2, 3}, 0, 2},
		{"no sum found", 99, []int{1, 2, 3}, -1, -1},
		{"actual test input", 2020, adventTestValues, -1, -1},
	}

	for _, td := range testData {
		t.Run(td.name, func(t *testing.T) {
			i, j := indexOfFirstToSum(td.sum, td.slice)

			assert.Equal(t, td.first, i)
			assert.Equal(t, td.second, j)
		})
	}
}

func Test_the_answer(t *testing.T) {
	i, j := indexOfFirstToSum(2020, adventTestValues)
	assert.Equal(t, 0, adventTestValues[i]*adventTestValues[j])
}
