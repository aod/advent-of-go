package main

import (
	"fmt"

	"github.com/aoktayd/adventofgode/internal/input"
)

const unitDifference byte = 32

func react(polymer []byte) ([]byte, bool) {
	hasReacted := false
	reactionAt := 0

	for i := 0; i < len(polymer)-1; i++ {
		firstUnit := polymer[i]
		neighbouringUnit := polymer[i+1]

		if firstUnit+unitDifference == neighbouringUnit || neighbouringUnit+unitDifference == firstUnit {
			hasReacted = true
			reactionAt = i
			break
		}
	}

	polymer = append(polymer[:reactionAt], polymer[reactionAt+2:]...)

	return polymer, hasReacted
}

func chainReaction(polymer []byte) []byte {
	polymerCopy := append([]byte{}, polymer...)
	for {
		newPolymer, hasReacted := react(polymerCopy)
		if !hasReacted {
			break
		}
		polymerCopy = newPolymer
	}
	return polymerCopy
}

func main() {
	input := input.Read(input.File{Puzzle: input.Puzzle{Year: 2018, Day: 5}, FileType: ".txt"})
	inputLength := len(input)

	part1 := chainReaction(input)

	const baseUnit = 65
	const unitTypes = 25
	smallestPolymer := inputLength

	for i := 0; i <= unitTypes; i++ {
		unitToRemove := byte(baseUnit + i)
		var polymer []byte
		for _, unit := range input {
			if unit == unitToRemove || unit == unitToRemove+unitDifference {
				continue
			}
			polymer = append(polymer, unit)
		}

		if inputLength == len(polymer) {
			continue
		}

		polymer = chainReaction(polymer)
		polymerLength := len(polymer)

		if polymerLength < smallestPolymer {
			smallestPolymer = polymerLength
		}
	}

	fmt.Println("Part 1: ", len(part1))
	fmt.Println("Part 2: ", smallestPolymer)
}
