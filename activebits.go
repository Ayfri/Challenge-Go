package piscine

func ActiveBits(n int) int {
	var count int
	for i := 0; n > 0; i++ {
		count++
		n >>= 1
	}
	return count
}
