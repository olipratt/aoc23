package aoc

import (
	"log"
	"strings"
)

// Calculate the numbers created by finding the first and last digits and
// concatenating them in each line of the given string.
func Day1(input string) []int {
	lines := strings.Split(input, "\n")

	result := []int{}

	for _, l := range lines {
		firstDigit := firstDigitInString(l)
		lastDigit := firstDigitInString(reverseString(l))

		result = append(result, firstDigit*10+lastDigit)
	}

	return result
}

func firstDigitInString(input string) int {
	for _, c := range input {
		if strings.ContainsRune("0123456789", c) {
			return int(c) - int('0')
		}
	}

	log.Fatalf("No int in string %v", input)
	return -1
}

func reverseString(input string) string {
	length := len(input)
	result := make([]rune, length)
	for i, c := range input {
		result[length-1-i] = c
	}
	return string(result)
}
