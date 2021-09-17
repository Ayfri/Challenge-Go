package piscine

func Join(elems []string, sep string) string {
	var result string
	for index, str := range elems {
		result += str
		if index < len(elems)-1 {
			result += sep
		}
	}
	return result
}
