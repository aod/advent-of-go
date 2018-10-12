package solution

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

// Part1 AOC year 2017 day 1 part 1
func Part1(input []byte) uint16 {
	return solution(input, 1)
}

// Part2 AOC year 2017 day 1 part 2
func Part2(input []byte) uint16 {
	return solution(input, 2)
}
