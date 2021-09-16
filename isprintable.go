package piscine

func IsPrintable(s string) bool {
	nonPrintable := []rune("\b\r\t\n\f")

	if len(s) < 1 {
		return false
	}
	for _, char := range s {
		for _, nonPrintableChar := range nonPrintable {
			if char == nonPrintableChar {
				return false
			}
		}
	}
	return true
}
