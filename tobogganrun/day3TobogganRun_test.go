package tobogganrun

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_tree_count_matches_expected(t *testing.T) {

	testData := []struct {
		name        string
		treeCount   int
		right       int
		down        int
		mountainMap []string
	}{
		{"Empty map", 0, 3, 1, []string{""}},
		{"First two rows", 1, 3, 1, []string{"..##.......", "#...#...#..", ".#....#..#."}},
		{"First four rows", 1, 3, 1, []string{"..##.......", "#...#...#..", ".#....#..#.", "..#.#...#.#"}},
		{"First five rows", 2, 3, 1, []string{"..##.......", "#...#...#..", ".#....#..#.", "..#.#...#.#", ".#...##..#."}},
	}

	for _, td := range testData {
		t.Run(td.name, func(t *testing.T) {
			result := CountTrees(td.right, td.down, td.mountainMap)

			assert.Equal(t, td.treeCount, result)
		})
	}
}

func Test_tree_count_is_expected_for_sample_input(t *testing.T) {

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

	testData := []struct {
		name      string
		treeCount int
		right     int
		down      int
	}{
		{"Right 1, down 1 is 2", 2, 1, 1},
		{"Right 3, down 1 is 7", 7, 3, 1},
		{"Right 5, down 1 is 3", 3, 5, 1},
		{"Right 7, down 1 is 4", 4, 7, 1},
		{"Right 1, down 2 is 2", 2, 1, 2},
	}

	for _, td := range testData {
		t.Run(td.name, func(t *testing.T) {
			result := CountTrees(td.right, td.down, mountainMap)

			assert.Equal(t, td.treeCount, result)
		})
	}
}

func Test_day3_part1_answer(t *testing.T) {

	result := CountTrees(3, 1, day3Part1TestInput)

	assert.Equal(t, 272, result)
}

func Test_day3_part2_answer(t *testing.T) {

	testData := []struct {
		name  string
		right int
		down  int
	}{
		{"Right 1, down 1", 1, 1},
		{"Right 3, down 1", 3, 1},
		{"Right 5, down 1", 5, 1},
		{"Right 7, down 1", 7, 1},
		{"Right 1, down 2", 1, 2},
	}

	multipliedTotal := 1

	for _, td := range testData {
		t.Run(td.name, func(t *testing.T) {
			result := CountTrees(td.right, td.down, day3Part1TestInput)

			multipliedTotal *= result
		})
	}

	assert.Equal(t, 3898725600, multipliedTotal)
}
