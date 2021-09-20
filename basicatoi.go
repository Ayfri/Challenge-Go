package piscine

func BasicAtoi(s string) int {
	arr := []rune(s)
	y := 0
	final := 0
	for range s {
		y++
	}
	for x := 0; x <= y-1; x++ {
		final = final*10 + int(arr[x]-'0')
	}
	return final
}
