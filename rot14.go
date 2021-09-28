package piscine

func Rot14(s string) string {
	var result string
	for _, char := range s {
		if IsUpper(string(char)) {
			if char+14 > 'Z' {
				difference := 'Z' - char
				char = 'A' + 13 - difference
				result += string(char)
			} else {
				result += string(char + 14)
			}
		} else if IsLower(string(char)) {
			if char+14 > 'z' {
				difference := 'z' - char
				char = 'a' + 13 - difference
				result += string(char)
			} else {
				result += string(char + 14)
			}
		} else {
			result += string(char)
		}
	}
	return result
}
