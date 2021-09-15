package piscine


func FindNextPrime(nb int) int {
	if nb > 1 {
		if IsPrimeOptimized(nb) {
			return nb
		}
		for i := nb + 1; i < (1 << 31); i += 2 {
			if IsPrimeOptimized(i) {
				return i
			}
		}
	}
	return 2
}

func IsPrimeOptimized(nb int) bool {
	for i := 3; i <= nb/2; i += 2 {
		if nb%i == 0 {
			return false
		}
	}
	return true
}

