package aoc

import "strings"

func isDigit(char rune) bool {
	return strings.ContainsRune("0123456789", char)
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
