package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	const zero = rune('0')
	str := ""

	for i := rune('0'); i <= rune('9'); i++ {
		for j := rune('0'); j <= rune('9'); j++ {
			for k := rune('0'); k <= rune('9'); k++ {
				for l := rune('0'); l <= rune('9'); l++ {
					first := string(i) + string(j)
					second := string(k) + string(l)
					if first != second {
						str += first + " " + second + ", "
					}
				}
			}
		}
	}

	str += "\n"

	for _, char := range str {
		z01.PrintRune(char)
	}
}
