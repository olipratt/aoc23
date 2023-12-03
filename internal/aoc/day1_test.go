package aoc_test

import (
	"testing"

	"olipratt/aoc23/internal/aoc"

	"github.com/stretchr/testify/assert"
)

func TestDay1(t *testing.T) {
	t.Parallel()

	input := `1abc2
	pqr3stu8vwx
	a1b2c3d4e5f
	treb7uchet`

	values := aoc.Day1(input, false)
	assert.Equal(t, []int{12, 38, 15, 77}, values)

	input = readFile(t, "input/day1.txt")
	values = aoc.Day1(input, false)
	total := 0
	for _, value := range values {
		total += value
	}
	assert.Equal(t, 54697, total)
}
func TestDay1b(t *testing.T) {
	t.Parallel()

	input := `two1nine
	eightwothree
	abcone2threexyz
	xtwone3four
	4nineeightseven2
	zoneight234
	7pqrstsixteen`

	values := aoc.Day1(input, true)
	assert.Equal(t, []int{29, 83, 13, 24, 42, 14, 76}, values)

	input = readFile(t, "input/day1.txt")
	values = aoc.Day1(input, true)
	total := 0
	for _, value := range values {
		total += value
	}
	assert.Equal(t, 54885, total)
}
