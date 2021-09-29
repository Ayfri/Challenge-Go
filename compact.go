package piscine

func Compact(ptr *[]string) int {
	var result []string
	var empty string
	for _, element := range *ptr {
		if element != empty {
			result = append(result, element)
		}
	}
	ptr = &result
	return len(result)
}
