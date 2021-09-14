package piscine

func IterativeFactorial(nb int) int {
	result := 0
	for i := nb; i > 0; i-- {
		result += i * nb
	}

	return result
}
