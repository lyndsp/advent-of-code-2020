package tobogganrun

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_tree_count_matches_expected(t *testing.T) {

	testData := []struct {
		name        string
		mountainMap []string
		treeCount   int
	}{
		{"Empty map", []string{""}, 0},
		{"First two rows", []string{"..##.......", "#...#...#..", ".#....#..#."}, 1},
		{"First four rows", []string{"..##.......", "#...#...#..", ".#....#..#.", "..#.#...#.#"}, 1},
		{"First five rows", []string{"..##.......", "#...#...#..", ".#....#..#.", "..#.#...#.#", ".#...##..#."}, 2},
	}

	for _, td := range testData {
		t.Run(td.name, func(t *testing.T) {
			result := CountTrees(td.mountainMap)

			assert.Equal(t, td.treeCount, result)
		})
	}
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
