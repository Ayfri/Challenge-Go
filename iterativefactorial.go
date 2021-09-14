package piscine

func IterativeFactorial(nb int) int {
	result := 0
	for i := nb; i > 0; i-- {
		if nb > 1000000002 {
			return 0
		}
		result += i * nb
	}

	return result
}
