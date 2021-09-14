package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	const zero = rune('0')
	str := ""

	for i := '0'; i <= '9'; i++ {
		for j := '0'; j <= '9'; j++ {
			for k := '0'; k <= '9'; k++ {
				for l := '0'; l <= '9'; l++ {
					first := string(i) + string(j)
					second := string(k) + string(l)
					if i > k || (i == k && j >= l) || first == "99" {
						continue
					}
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
