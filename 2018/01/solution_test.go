package day01

import (
	"fmt"
	"testing"

	"github.com/aoktayd/aoc/input"
)

var puzzleInput = input.File().IntArray()

func TestPart1(t *testing.T) {
	t.Run("Sample input", func(t *testing.T) {
		testCases := []struct {
			input  []int
			answer int
		}{
			{
				input:  []int{1, 1, 1},
				answer: 3,
			},
			{
				input:  []int{1, 1, -2},
				answer: 0,
			},
			{
				input:  []int{-1, -2, -3},
				answer: -6,
			},
		}

		for i, tC := range testCases {
			t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
				if output := part1(tC.input); output != tC.answer {
					t.Errorf("part1(%#v) returned %d, expected %d", tC.input, output, tC.answer)
				}
			})
		}
	})

	t.Run("Puzzle input", func(t *testing.T) {
		if output, answer := part1(puzzleInput), 402; output != answer {
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
			input  []int
			answer int
		}{
			{
				input:  []int{1, -1},
				answer: 0,
			},
			{
				input:  []int{3, 3, 4, -2, -4},
				answer: 10,
			},
			{
				input:  []int{-6, 3, 8, 5, -6},
				answer: 5,
			},
			{
				input:  []int{7, 7, -2, -7, -4},
				answer: 14,
			},
		}

		for i, tC := range testCases {
			t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
				if output := part2(tC.input); output != tC.answer {
					t.Errorf("part2(%#v) returned %d, expected %d", tC.input, output, tC.answer)
				}
			})
		}
	})

	t.Run("Puzzle input", func(t *testing.T) {
		if output, answer := part2(puzzleInput), 481; output != answer {
			t.Errorf("Expected %d, got %d", answer, output)
		}
	})
}

func BenchmarkPart2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		part2(puzzleInput)
	}
}
