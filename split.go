package piscine

func Split(s, sep string) []string {
	var result []string
	var cache string

	for index, char := range s {
		if index+len(sep) >= len(s) {
			continue
		}
		if sep != s[index:index+len(sep)] {
			cache += string(char)
		} else if len(cache) > 0 {
			result = append(result, cache)
			cache = ""
		}
	}
	if len(cache) > 0 {
		result = append(result, cache)
	}

	return result
}
