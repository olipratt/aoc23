package aoc

import (
	"log"
	"strconv"
	"strings"
)

// Calculate which numbers in the multiline input string are adjacent to a
// (non-'.') symbol. Includes diagonally adjacent.
func Day3(input string) []int {
	schematic := newEngineSchematic(input)
	result := schematic.partNumbers()
	return result
}

type partNumber struct {
	value  int
	line   int
	index  int
	length int
}

type engineSchematic struct {
	schematic []string
}

func newEngineSchematic(input string) engineSchematic {
	lines := strings.Split(input, "\n")
	return engineSchematic{schematic: lines}
}

func (s engineSchematic) isSymbol(char byte) bool {
	return strings.ContainsRune("*#+$@=/-_|\\%&", rune(char))
}

func (s engineSchematic) isValid(partNum partNumber) bool {
	// Check the left chars for symbols.
	if (partNum.index) > 0 {
		// Check immediate left.
		if s.isSymbol(s.schematic[partNum.line][partNum.index-1]) {
			return true
		}
		// Check above left.
		if partNum.line > 0 && s.isSymbol(s.schematic[partNum.line-1][partNum.index-1]) {
			return true
		}
		// Check below left
		if partNum.line < len(s.schematic)-1 && s.isSymbol(s.schematic[partNum.line+1][partNum.index-1]) {
			return true
		}
	}

	// Check the right chars for symbols.
	if (partNum.index + partNum.length) < len(s.schematic[partNum.line]) {
		// Check immediate right.
		if s.isSymbol(s.schematic[partNum.line][partNum.index+partNum.length]) {
			return true
		}
		// Check above right.
		if partNum.line > 0 && s.isSymbol(s.schematic[partNum.line-1][partNum.index+partNum.length]) {
			return true
		}
		// Check below right
		if partNum.line < len(s.schematic)-1 && s.isSymbol(s.schematic[partNum.line+1][partNum.index+partNum.length]) {
			return true
		}
	}

	// Check directly above / below.
	if partNum.line > 0 {
		for offset := 0; offset < partNum.length; offset++ {
			if s.isSymbol(s.schematic[partNum.line-1][partNum.index+offset]) {
				return true
			}
		}
	}
	if partNum.line < len(s.schematic)-1 {
		for offset := 0; offset < partNum.length; offset++ {
			if s.isSymbol(s.schematic[partNum.line+1][partNum.index+offset]) {
				return true
			}
		}
	}

	return false
}

func (s engineSchematic) parseSchematicLine(lineIndex int) []partNumber {
	partNumbers := make([]partNumber, 0)

	line := s.schematic[lineIndex]
	lineLen := len(line)
	for index := 0; index < lineLen; index++ {
		if isDigit(rune(line[index])) {
			numStart := index
			var numLen int
			// Loop, increasing length, until we hit a non-digit.
			for numLen = 1; numStart+numLen < lineLen &&
				isDigit(rune(line[numStart+numLen])); numLen++ {
			}

			value, err := strconv.Atoi(line[numStart : numStart+numLen])
			if err != nil {
				log.Fatalf("Failed to parse as integer: %s", line[numStart:numStart+numLen])
			}

			partNum := partNumber{value: value, line: lineIndex, index: numStart, length: numLen}
			partNumbers = append(partNumbers, partNum)

			// This, after the loop increment, takes us to the char after the
			// character after the last digit. This is fine though because we
			// now know the char after the last digit is not a digit.
			index += numLen
		}
	}

	return partNumbers
}

// Get the numbers that are valid partnumbers only.
func (s engineSchematic) partNumbers() []int {
	numbers := make([]int, 0)

	for index := range s.schematic {
		partNumbers := s.parseSchematicLine(index)
		for _, partNum := range partNumbers {
			if s.isValid(partNum) {
				numbers = append(numbers, partNum.value)
			}
		}
	}

	return numbers
}
