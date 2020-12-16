package tobogganrun

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_when_map_is_empty_tree_count_is_zero(t *testing.T) {

	mountainMap := []string{""}
	result := CountTrees(mountainMap)

	assert.Equal(t, 0, result)
}

func Test_tree_count_is_one_for_first_two_rows(t *testing.T) {

	mountainMap := []string{"..##.......", "#...#...#..", ".#....#..#."}
	result := CountTrees(mountainMap)

	assert.Equal(t, 1, result)
}
func Test_tree_count_is_one_for_first_four_rows(t *testing.T) {

	mountainMap := []string{"..##.......", "#...#...#..", ".#....#..#.", "..#.#...#.#"}
	result := CountTrees(mountainMap)

	assert.Equal(t, 1, result)
}
func Test_tree_count_is_two_for_first_five_rows(t *testing.T) {

	mountainMap := []string{"..##.......", "#...#...#..", ".#....#..#.", "..#.#...#.#", ".#...##..#."}
	result := CountTrees(mountainMap)

	assert.Equal(t, 2, result)
}
func Test_tree_count_is_seven_for_sample_input(t *testing.T) {

	mountainMap := []string{
		"..##.......",
		"#...#...#..",
		".#....#..#.",
		"..#.#...#.#",
		".#...##..#.",
		"..#.##.....",
		".#.#.#....#",
		".#........#",
		"#.##...#...",
		"#...##....#",
		".#..#...#.#",
	}
	result := CountTrees(mountainMap)

	assert.Equal(t, 7, result)
}

func Test_day3_part1_answer(t *testing.T) {

	result := CountTrees(day3Part1TestInput)

	assert.Equal(t, 272, result)
}
