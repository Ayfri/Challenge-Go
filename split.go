package piscine

func Split(s, sep string) []string {
	var result []string
	var cache string
	var isValid bool

	for index, char := range s {
		isValid = true
		for i := 0; i < len(sep); i++ {
			if index-len(sep) < 0 || index+i >= len(s) {
				continue
			}

			if sep != s[index-len(sep)+i:index+i] {
				isValid = false
			}
		}

		if isValid {
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
