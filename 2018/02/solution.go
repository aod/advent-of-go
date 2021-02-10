package day02

func part1(input []string) int {
	checksums := make(map[int]int)

	for _, str := range input {
		word := make(map[rune]int)
		for _, letter := range str {
			word[letter]++
		}

		seen := make(map[int]interface{})
		seen[2] = nil
		seen[3] = nil

		for _, letterAmount := range word {
			if _, exists := seen[letterAmount]; exists {
				checksums[letterAmount]++
				delete(seen, letterAmount)
			}
			if len(seen) == 0 {
				break
			}
		}
	}

	return checksums[2] * checksums[3]
}

func part2(input []string) string {
	checked := make(map[int]bool)

	for k, str := range input {
		checked[k] = true

		for k2, str2 := range input {
			if checked[k2] {
				continue
			} else if diff := difference(str, str2); diff != -1 {
				return string(str[:diff] + str[diff+1:])
			}
		}
	}

	return ""
}
