package aoc

import "strings"

func isDigit(char rune) bool {
	return strings.ContainsRune("0123456789", char)
}
