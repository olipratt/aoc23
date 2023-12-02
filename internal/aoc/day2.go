package aoc

import (
	"log"
	"strconv"
	"strings"
)

// Calculate which lines in the input are possible games with:
// 12 red cubes, 13 green cubes, and 14 blue cubes.
func Day2(input string) []int {
	lines := strings.Split(input, "\n")

	result := []int{}

	for _, line := range lines {
		game := parseBagGame(line)
		if game.possible() {
			result = append(result, game.id)
		}
	}

	return result
}

// Calculate the power of each game.
// Power is defined as the counts of the minimal set of cubes to play a game,
// multiplied together.
func Day2b(input string) []int {
	lines := strings.Split(input, "\n")

	result := []int{}

	for _, line := range lines {
		game := parseBagGame(line)
		result = append(result, game.power())
	}

	return result
}

type BagRound struct {
	red   int
	green int
	blue  int
}

// A round is possible if it uses no more than:
// 12 red, 13 green, 14 blue.
func (r BagRound) possible() bool {
	//nolint:gomnd // magic numbers from problem.
	return (r.red <= 12) && (r.green <= 13) && (r.blue <= 14)
}

// Parse something that looks like:
// "1 red, 2 green, 6 blue".
func parseBagRound(input string) BagRound {
	round := BagRound{red: 0, green: 0, blue: 0}

	// First get each colour statement.
	entryStrings := strings.Split(input, ",")
	for _, entryString := range entryStrings {
		// Now split into count and colour.
		entryString = strings.TrimSpace(entryString)
		splitEntry := strings.Split(entryString, " ")
		//nolint:gomnd // string must split in 2.
		if len(splitEntry) != 2 {
			log.Fatal("Extra ' ' in string")
		}

		// Parse the count.
		count, err := strconv.Atoi(splitEntry[0])
		if err != nil {
			log.Fatalf("Failed to parse as integer: %s", splitEntry[0])
		}

		// Check the colour and set the count for that colour.
		switch splitEntry[1] {
		case "red":
			round.red = count
		case "green":
			round.green = count
		case "blue":
			round.blue = count
		}
	}

	return round
}

type BagGame struct {
	id     int
	rounds []BagRound
}

func (g BagGame) possible() bool {
	for _, round := range g.rounds {
		if !round.possible() {
			return false
		}
	}

	return true
}

// Power is defined as the counts of the minimal set of cubes to play a game,
// multiplied together.
func (g BagGame) power() int {
	minimalRound := BagRound{red: 0, green: 0, blue: 0}

	for _, round := range g.rounds {
		if round.red > minimalRound.red {
			minimalRound.red = round.red
		}
		if round.green > minimalRound.green {
			minimalRound.green = round.green
		}
		if round.blue > minimalRound.blue {
			minimalRound.blue = round.blue
		}
	}

	return minimalRound.red * minimalRound.green * minimalRound.blue
}

// Parse something that looks like:
// "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green".
func parseBagGame(input string) BagGame {
	splitString := strings.Split(input, ":")
	//nolint:gomnd // string must split in 2.
	if len(splitString) != 2 {
		log.Fatalf("Extra ':' in string")
	}

	// Parse the initial ID number for the game.
	gameIDStrings := strings.Split(splitString[0], " ")
	gameIDString := gameIDStrings[len(gameIDStrings)-1]
	gameID, err := strconv.Atoi(gameIDString)
	if err != nil {
		log.Fatalf("Failed to parse as integer: %s", gameIDString)
	}

	// Parse all the rounds.
	roundStrings := strings.Split(splitString[1], ";")
	rounds := make([]BagRound, 0, len(roundStrings))
	for _, roundString := range roundStrings {
		rounds = append(rounds, parseBagRound(roundString))
	}

	return BagGame{id: gameID, rounds: rounds}
}
