package aoc_test

import (
	"testing"

	"olipratt/aoc23/internal/aoc"

	"github.com/stretchr/testify/assert"
)

func TestDay3(t *testing.T) {
	t.Parallel()

	input := `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`

	values := aoc.Day3(input)
	assert.Equal(t, []int{467, 35, 633, 617, 592, 755, 664, 598}, values)

	input = readFile(t, "input/day3.txt")
	values = aoc.Day3(input)
	total := sumArray(values)
	assert.Equal(t, 531932, total)
}

func TestDay3b(t *testing.T) {
	t.Parallel()

	input := `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`

	values := aoc.Day3b(input)
	assert.Equal(t, []int{16345, 451490}, values)

	input = readFile(t, "input/day3.txt")
	values = aoc.Day3b(input)
	total := sumArray(values)
	assert.Equal(t, 73646890, total)
}
