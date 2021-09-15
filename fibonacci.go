package piscine

func Fibonacci(index int) int {
	result := 0

	if index < 0 {
		return -1
	}

	if index > 1 {
		result = Fibonacci(index-1) + Fibonacci(index-2)
	} else {
		return index
	}

	return result
}
