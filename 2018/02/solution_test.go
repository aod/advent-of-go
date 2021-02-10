package day02

import (
	"fmt"
	"testing"

	"github.com/aoktayd/aoc/input"
)

var puzzleInput = input.File().StringArray()

func TestPart1(t *testing.T) {
	t.Run("Sample input", func(t *testing.T) {
		testCases := []struct {
			input  []string
			answer int
		}{
			{
				input: []string{
					"abcdef",
					"bababc",
					"abbcde",
					"abcccd",
					"aabcdd",
					"abcdee",
					"ababab",
				},
				answer: 12,
			},
		}

		for i, tC := range testCases {
			t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
				if output := part1(tC.input); output != tC.answer {
					t.Errorf("Part 1 failed, expected %d, got %d", tC.answer, output)
				}
			})
		}
	})

	t.Run("Puzzle input", func(t *testing.T) {
		if output, answer := part1(puzzleInput), 7192; output != answer {
			t.Errorf("Expected %d, got %d", answer, output)
		}
	})
}

func BenchmarkPart1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		part1(puzzleInput)
	}
}

func TestPart2(t *testing.T) {
	t.Run("Sample input", func(t *testing.T) {
		testCases := []struct {
			input  []string
			answer string
		}{
			{
				input: []string{
					"abcde",
					"fghij",
					"klmno",
					"pqrst",
					"fguij",
					"axcye",
					"wvxyz",
				},
				answer: "fgij",
			},
		}

		for i, tC := range testCases {
			t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
				if output := part2(tC.input); output != tC.answer {
					t.Errorf("Part 2 failed, expected %s, got %s", tC.answer, output)
				}
			})
		}
	})

	t.Run("Puzzle input", func(t *testing.T) {
		if output, answer := part2(puzzleInput), "mbruvapghxlzycbhmfqjonsie"; output != answer {
			t.Errorf("Expected %s, got %s", answer, output)
		}
	})
}

func BenchmarkPart2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		part2(puzzleInput)
	}
}
