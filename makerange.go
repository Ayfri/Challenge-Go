package piscine

func MakeRange(min, max int) []int {
	var arr []int

	for i := min; i < max; i++ {
		arr = append(arr, i)
	}

	return arr
}
