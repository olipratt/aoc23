package aoc_test

import (
	"testing"

	"olipratt/aoc23/internal/aoc"

	"github.com/stretchr/testify/assert"
)

func TestDay2(t *testing.T) {
	t.Parallel()

	input := `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
	Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
	Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
	Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
	Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`

	values := aoc.Day2(input)
	assert.Equal(t, []int{1, 2, 5}, values)

	input = readFile(t, "input/day2.txt")
	values = aoc.Day2(input)
	total := aoc.SumArray(values)
	assert.Equal(t, 2593, total)
}

func TestDay2b(t *testing.T) {
	t.Parallel()

	input := `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
	Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
	Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
	Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
	Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`

	values := aoc.Day2b(input)
	assert.Equal(t, []int{48, 12, 1560, 630, 36}, values)

	input = readFile(t, "input/day2.txt")
	values = aoc.Day2b(input)
	total := aoc.SumArray(values)
	assert.Equal(t, 54699, total)
}
