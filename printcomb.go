package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	const newLine = rune('\n')
	const zero = rune('0')
	const nine = rune('9')

	for i := zero; i <= nine; i++ {
		for j := zero; j <= nine; j++ {
			for k := zero; k <= nine; k++ {
				if i == j || j == k || i == k {
					continue
				}

				z01.PrintRune(i)
				z01.PrintRune(j)
				z01.PrintRune(k)
				z01.PrintRune(rune(','))
				z01.PrintRune(rune(' '))
			}
		}
	}
}
