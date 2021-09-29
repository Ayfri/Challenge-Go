package piscine

// divide by two when even
// multiply by three when odd & add one

func CollatzCountdown(start int) int {
	for i := 0; i <= start; i++ {
		if start == 1 {
			return i
		}

		if start%2 == 0 {
			start /= 2
		} else {
			start = (start * 3) + 1
		}
	}
	return -1
}
