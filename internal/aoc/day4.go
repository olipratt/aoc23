package aoc

import (
	"log"
	"slices"
	"strconv"
	"strings"
)

// Calculate the points for each scratch card defined in the input.
func Day4(input string) []int {
	lines := strings.Split(input, "\n")

	result := []int{}

	for _, line := range lines {
		card := parseScratchCard(line)
		score := card.score()
		result = append(result, score)
	}

	return result
}

func parseSpaceSeparatedNumbers(input string) []int {
	result := make([]int, 0)
	for _, numStr := range strings.Split(input, " ") {
		if len(numStr) == 0 {
			// Some entries might be empty if there are consecutive spaces for
			// alignment.
			continue
		}

		value, err := strconv.Atoi(numStr)
		if err != nil {
			log.Fatalf("Failed to parse as integer: '%s'", numStr)
		}
		result = append(result, value)
	}

	return result
}

type scratchCard struct {
	id             int
	winningNumbers []int
	drawnNumbers   []int
}

func parseScratchCard(input string) scratchCard {
	splitString := strings.Split(input, ": ")
	//nolint:gomnd // string must split in 2.
	if len(splitString) != 2 {
		log.Fatalf("Extra/no ':' in string")
	}

	// Parse the ID number for the card.
	cardIDStrings := strings.Split(splitString[0], " ")
	cardIDString := cardIDStrings[len(cardIDStrings)-1]
	cardID, err := strconv.Atoi(cardIDString)
	if err != nil {
		log.Fatalf("Failed to parse as integer: '%s'", cardIDString)
	}

	// Split the remainder into winning and drawn numbers.
	winAndDrawnNums := strings.Split(splitString[1], " | ")
	//nolint:gomnd // string must split in 2.
	if len(winAndDrawnNums) != 2 {
		log.Fatalf("Extra/no ' | ' in string")
	}

	// Parse the numbers.
	winningNumbers := parseSpaceSeparatedNumbers(winAndDrawnNums[0])
	drawnNumbers := parseSpaceSeparatedNumbers(winAndDrawnNums[1])

	return scratchCard{id: cardID, winningNumbers: winningNumbers, drawnNumbers: drawnNumbers}
}

func (s scratchCard) score() int {
	winnersCount := 0
	for _, num := range s.drawnNumbers {
		if slices.Contains(s.winningNumbers, num) {
			winnersCount++
		}
	}

	if winnersCount == 0 {
		return 0
	}

	//nolint:gomnd // Score is a power of 2.
	return intPow(2, uint(winnersCount-1))
}
