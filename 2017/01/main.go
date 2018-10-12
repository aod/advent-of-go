package main

import (
	"fmt"

	"github.com/aoktayd/adventofgode/utils"
)

func main() {
	input := ParseInput(utils.RelFile("input.txt"))

	fmt.Printf("Part 1: %d\n", Solution(input, 1))
	fmt.Printf("Part 2: %d\n", Solution(input, 2))
}
