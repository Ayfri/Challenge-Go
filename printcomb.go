package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	const newLine = rune('\n')

	for i := rune('0'); i <= rune('7'); i++ {
		for j := rune('1'); j <= rune('8'); j++ {
			for k := rune('1'); k <= rune('9'); k++ {
				if i >= j || j >= k {
					continue
				}
				z01.PrintRune(i)
				z01.PrintRune(j)
				z01.PrintRune(k)
				if string([]rune{i, j, k}) == "789" {
					z01.PrintRune(newLine)
					break
				}
				z01.PrintRune(rune(','))
				z01.PrintRune(rune(' '))
			}
		}
	}
}
