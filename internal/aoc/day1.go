package aoc

import (
	"errors"
	"log"
	"strings"
)

var errNoDigit = errors.New("no digit error")

// Calculate the numbers created by finding the first and last digits and
// concatenating them in each line of the given string.
func Day1(input string, includeWords bool) []int {
	lines := strings.Split(input, "\n")

	result := []int{}

	for _, l := range lines {
		firstDigit := firstDigitInString(l, includeWords)
		lastDigit := lastDigitInString(l, includeWords)

		result = append(result, firstDigit*10+lastDigit)
	}

	return result
}

func firstDigitInString(input string, includeWords bool) int {
	for i := 0; i < len(input); i++ {
		digit, err := digitAtIndexInString(input, i, includeWords)
		if err == nil {
			return digit
		}
	}

	log.Fatalf("No int in string %v", input)
	return -1
}

func lastDigitInString(input string, includeWords bool) int {
	for i := len(input) - 1; i >= 0; i-- {
		digit, err := digitAtIndexInString(input, i, includeWords)
		if err == nil {
			return digit
		}
	}

	log.Fatalf("No int in string %v", input)
	return -1
}

func digitAtIndexInString(input string, index int, includeWords bool) (int, error) {
	// Strings for all digits with indices equal to their value.
	digitStrings := []string{"zero", "one", "two", "three", "four", "five",
		"six", "seven", "eight", "nine"}

	char := rune(input[index])

	if isDigit(char) {
		return int(char) - int('0'), nil
	}
	if includeWords {
		// Get the substring starting at this point and see if any of the
		// word for digits prefix that substring.
		substring := input[index:]
		for digit, digitWord := range digitStrings {
			if strings.HasPrefix(substring, digitWord) {
				return digit, nil
			}
		}
	}

	return 0, errNoDigit
}

//nolint:unused // May be useful later.
func reverseString(input string) string {
	length := len(input)
	result := make([]rune, length)
	for i, c := range input {
		result[length-1-i] = c
	}
	return string(result)
}
