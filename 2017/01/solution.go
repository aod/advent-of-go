package main

// ParseInput Converts the input to a numbers byte array
func ParseInput(input []byte) []byte {
	var parsedInput []byte
	for _, v := range input {
		parsedInput = append(parsedInput, v-48)
	}
	return parsedInput
}

// Solution AOC year 2017 day 1
func Solution(input []byte, distance int) uint16 {
	var result uint16
	inputLength := len(input)

	for i, v := range input {
		if v == input[(i+distance)%inputLength] {
			result += uint16(v)
		}
	}

	return result
}
