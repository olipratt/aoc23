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

// Calculate the number of scratch cards in total (original + won).
func Day4b(input string) int {
	lines := strings.Split(input, "\n")

	scratchCards := make([]scratchCard, 0, len(lines))
	for _, line := range lines {
		scratchCards = append(scratchCards, parseScratchCard(line))
	}

	counter := createScratchCardCounter(scratchCards)
	cardTotal := SumArray(counter.cardCounts)

	return cardTotal
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

func (s scratchCard) winsCount() int {
	winnersCount := 0
	for _, num := range s.drawnNumbers {
		if slices.Contains(s.winningNumbers, num) {
			winnersCount++
		}
	}
	return winnersCount
}

func (s scratchCard) score() int {
	winnersCount := s.winsCount()

	if winnersCount == 0 {
		return 0
	}

	//nolint:gomnd // Score is a power of 2.
	return intPow(2, uint(winnersCount-1))
}

type scratchCardCounter struct {
	scratchCards []scratchCard
	cardCounts   []int
}

func createScratchCardCounter(scratchCards []scratchCard) scratchCardCounter {
	// Assumes card IDs go from 1 - #(cards), and cards are provided in order
	// in the input array.

	// Initialise the counter to 1 for each card we have.
	// There is no zero-th card, so leave that entry 0.
	initialCounts := make([]int, len(scratchCards)+1)
	for i := 1; i < len(initialCounts); i++ {
		initialCounts[i] = 1
	}

	counter := scratchCardCounter{scratchCards: scratchCards, cardCounts: initialCounts}
	counter.populateCounts()

	return counter
}

func (c *scratchCardCounter) populateCounts() {
	// Assume the input is set up such that we won't ever have cards added
	// beyond the end of the input array, since handling that is not defined.
	for _, card := range c.scratchCards {
		wins := card.winsCount()
		for i := 0; i < wins; i++ {
			c.cardCounts[card.id+1+i] += c.cardCounts[card.id]
		}
	}
}
