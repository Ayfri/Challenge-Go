package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	if len(a) < 2 {
		return true
	}

	isReversed := f(a[0], a[1]) > 0
	for i := 0; i < len(a)-1; i++ {
		if f(a[i], a[i]) <= 0 {
			if isReversed {
				return false
			}
		} else if !isReversed {
			return false
		}
	}
	return true
}
