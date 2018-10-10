package main

import (
	"fmt"

	"github.com/aoktayd/adventofgode/utils"
)

func main() {
	input := ParseInput(utils.PuzzleInput(2017, 1))

	fmt.Printf("Part 1: %d\n", Solution(input, 1))
	fmt.Printf("Part 2: %d\n", Solution(input, 2))
}
