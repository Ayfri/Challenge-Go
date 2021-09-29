package piscine

func Abort(a, b, c, d, e int) int {
	arr := []int{a, b, c, d, e}

	medianNumber := len(arr) / 2

	if len(arr)%2 == 0 {
		return arr[medianNumber]
	}

	return (arr[medianNumber-1] + arr[medianNumber]) / 2
}
