package piscine

func FindNextPrime(nb int) int {
	if nb > 1 {
		if IsPrime(nb) {
			return nb
		}
		for i := nb; i < (i << 31); i++ {
			if IsPrime(i) {
				return i
			}
		}
	}
	return 0
}

func IsPrime(nb int) bool {
	result := true
	for i := 2; i <= nb/2; i++ {
		if nb%i == 0 {
			result = false
		}
	}
	return result
}
