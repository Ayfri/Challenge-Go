package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	isSorted := true
	for index, elem := range a {
		if index+1 < len(a) && f(elem, a[index+1]) < 0 {
			isSorted = false
		}
	}
	return isSorted
}
