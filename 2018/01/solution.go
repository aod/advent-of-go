package day01

func part1(input []int) (frequency int) {
	for _, v := range input {
		frequency += v
	}

	return
}

func part2(input []int) (frequency int) {
	seen := make(map[int]bool, len(input))

	for i := 0; i < len(input)-1; i++ {
		seen[frequency] = true
		frequency += input[i]
	}

	for i, ok := len(input)-1, false; !ok; i, ok = i+1, seen[frequency] {
		frequency += input[i%len(input)]
	}

	return
}
