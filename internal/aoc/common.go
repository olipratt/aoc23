package aoc

import "strings"

func isDigit(char rune) bool {
	return strings.ContainsRune("0123456789", char)
}

// Calculate the power of an integer to some other integer.
func intPow(num int, power uint) int {
	result := 1
	for i := 0; i < int(power); i++ {
		result *= num
	}
	return result
}

type Point struct {
	x int
	y int
}

func (p Point) Add(q Point) Point {
	return Point{p.x + q.x, p.y + q.y}
}

func (p Point) Sub(q Point) Point {
	return Point{p.x - q.x, p.y - q.y}
}

func SumArray(a []int) int {
	total := 0
	for _, value := range a {
		total += value
	}
	return total
}
