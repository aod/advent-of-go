package day02

// difference compares 2 strings (which must be of equal length) and returns
// the index if the strings differ exactly 1 character.
// Otherwise it will return -1.
func difference(a, b string) int {
	index := -1
	differs := 0
	for k := range b {
		if a[k] != b[k] {
			differs++
			index = k
		}
		if differs > 1 {
			return -1
		}
	}
	return index
}
