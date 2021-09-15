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
