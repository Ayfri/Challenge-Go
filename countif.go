package piscine

func CountIf(f func(string) bool, tab []string) int {
	count := 0
	for _, elem := range tab {
		if f(elem) {
			count++
		}
	}

	return count
}
