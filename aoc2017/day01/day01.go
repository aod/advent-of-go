package main

import (
	"fmt"

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

	fmt.Printf("Part 1: %d\n", solution(parsedInput, 1))
	fmt.Printf("Part 2: %d\n", solution(parsedInput, len(parsedInput)/2))
}

func solution(input []byte, distance int) uint16 {
	var result uint16
	inputLength := len(input)

	for i, v := range input {
		if v == input[(i+distance)%inputLength] {
			result += uint16(v)
		}
	}

	return result
}
