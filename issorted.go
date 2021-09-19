package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	isSorted := true
	if len(a) < 2 {
		return true
	}

	isReversed := f(a[0], a[1]) > 0
	for index, elem := range a {
		if index+1 < len(a) {
			elemAfter := a[index+1]
			if f(elem, elemAfter) <= 0 {
				if isReversed {
					isSorted = false
				}
			} else if !isReversed {
				isSorted = false
			}
		}
	}
	return isSorted
}
