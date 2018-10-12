package main

import (
	"fmt"

	"github.com/aoktayd/adventofgode/2017/01/solution"
	"github.com/aoktayd/adventofgode/utils"
)

func main() {
	input, err := utils.ReadRelFile("input.txt")
	if err != nil {
		panic(err)
	}

	var parsedInput []byte
	for _, v := range input {
		parsedInput = append(parsedInput, v-48)
	}

	fmt.Printf("Part 1: %d\n", solution.Part1(parsedInput))
	fmt.Printf("Part 2: %d\n", solution.Part2(parsedInput))
}
