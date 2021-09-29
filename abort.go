package piscine

func Abort(a, b, c, d, e int) int {
	arr := []int{a, b, c, d, e}
	SortNumbers(arr)

	medianNumber := len(arr) / 2

	if len(arr)%2 == 1 {
		return arr[medianNumber]
	}

	return (arr[medianNumber-1] + arr[medianNumber]) / 2
}

func SortNumbers(a []int) {
	length := len(a)
	for i := 0; i < length; i++ {
		for j := 0; j < length-i-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
}
