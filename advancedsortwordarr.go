package piscine

func AdvancedSortWordArr(a []string, compare func(a, b string) int) {
	length := len(a)
	for i := 0; i < length; i++ {
		for j := 0; j < length-i-1; j++ {
			if compare(a[j], a[j+1]) != 0 {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
}
