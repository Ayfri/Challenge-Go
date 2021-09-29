package piscine

func ActiveBits(n int) int {
	if 0 <= n && n < 2 {
		return int(uint(n))
	}
	return int((uint(n) % 2) + uint(ActiveBits(n/2)))
}
