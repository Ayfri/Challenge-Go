package piscine

func IsPrime(nb int) bool {
	result := true
	if nb <= 1 {
		return false
	}
	for i := 2; i <= nb / 2; i++ {
		if nb % i == 0 {
			result = false
		}
	}
	return result
}
