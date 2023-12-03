package aoc

import (
	"errors"
	"log"
	"slices"
	"strconv"
	"strings"
)

var errNoPartNum = errors.New("no partNumber in this space")
var errNotGearRatio = errors.New("this point isn't a gear ratio")

// Calculate which numbers in the multiline input string are adjacent to a
// (non-'.') symbol. Includes diagonally adjacent.
func Day3(input string) []int {
	schematic := newEngineSchematic(input)
	result := schematic.validPartNumbers()
	return result
}

// Find all gear ratios.
// A gear ratio is when two numbers are connected by a '*' character.
func Day3b(input string) []int {
	schematic := newEngineSchematic(input)
	result := schematic.gearRatios()
	return result
}

type partNumber struct {
	value  int
	line   int
	index  int
	length int
}

func (n partNumber) occupiesCell(p Point) bool {
	return p.y == n.line && p.x >= n.index && p.x < (n.index+n.length)
}

type engineSchematic struct {
	schematic []string
	parts     []partNumber
}

func newEngineSchematic(input string) engineSchematic {
	lines := strings.Split(input, "\n")

	s := engineSchematic{schematic: lines, parts: nil}
	s.populatePartNumbers()

	return s
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

func (s *engineSchematic) populatePartNumbers() {
	parts := make([]partNumber, 0)

	for index := range s.schematic {
		partNumbers := s.parseSchematicLine(index)
		parts = append(parts, partNumbers...)
	}

	s.parts = parts
}

func (s engineSchematic) isSymbolAt(p Point) bool {
	return strings.ContainsRune("*#+$@=/-_|\\%&", rune(s.schematic[p.y][p.x]))
}

func (s engineSchematic) isValid(partNum partNumber) bool {
	// Find all surrounding cells.
	surroundingCells := make([]Point, 0)
	for i := 0; i < partNum.length; i++ {
		adjacentPoints := s.surroundingCells(Point{partNum.index + i, partNum.line})
		for _, point := range adjacentPoints {
			// Ignore duplicates and cells inside the part number.
			if !slices.Contains(surroundingCells, point) && !partNum.occupiesCell(point) {
				surroundingCells = append(surroundingCells, point)
			}
		}
	}

	// Check if any surrounding cell is a symbol.
	for _, point := range surroundingCells {
		if s.isSymbolAt(point) {
			return true
		}
	}

	return false
}

// Get the numbers that are valid part numbers only.
func (s engineSchematic) validPartNumbers() []int {
	numbers := make([]int, 0)

	for _, partNum := range s.parts {
		if s.isValid(partNum) {
			numbers = append(numbers, partNum.value)
		}
	}

	return numbers
}

func (s engineSchematic) surroundingCells(input Point) []Point {
	// Offsets to all surrounding points.
	offsets := []Point{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	result := make([]Point, 0)

	schematicHeight := len(s.schematic)
	schematicWidth := len(s.schematic[0])
	// Calculate all valid surrounding points.
	for _, offset := range offsets {
		pt := input.Add(offset)
		// The new point is only valid if it's within the schematic.
		if pt.x >= 0 && pt.x < schematicWidth && pt.y >= 0 && pt.y < schematicHeight {
			result = append(result, pt)
		}
	}

	return result
}

// Given a cell in the schematic, find the partNumber occupying it, if any.
func (s engineSchematic) partNumOccupyingCell(input Point) (partNumber, error) {
	for _, p := range s.parts {
		if p.occupiesCell(input) {
			return p, nil
		}
	}

	return partNumber{}, errNoPartNum
}

// Get the gear ratio at the given cell.
// This is the two adjacent numbers multiplied together.
func (s engineSchematic) gearRatio(input Point) (int, error) {
	inputs := make([]partNumber, 0)

	for _, cell := range s.surroundingCells(input) {
		partNum, err := s.partNumOccupyingCell(cell)
		// Be careful to avoid duplicates.
		if err == nil && !slices.Contains(inputs, partNum) {
			inputs = append(inputs, partNum)
			// If we have found both adjacent numbers, we're done.
			//nolint:gomnd // Two adjacent cells needed.
			if len(inputs) == 2 {
				return inputs[0].value * inputs[1].value, nil
			}
		}
	}

	log.Printf("Failed to find ratio at %v", input)
	return 0, errNotGearRatio
}

// Get all gear ratios in the schematic.
func (s engineSchematic) gearRatios() []int {
	ratios := make([]int, 0)

	for rowIdx, line := range s.schematic {
		for colIdx, char := range line {
			if char == '*' {
				ratio, err := s.gearRatio(Point{colIdx, rowIdx})
				if err == nil {
					ratios = append(ratios, ratio)
				}
			}
		}
	}

	return ratios
}
