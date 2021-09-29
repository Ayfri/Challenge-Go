package piscine

func Compact(ptr *[]string) int {
	var result []string
	for _, element := range *ptr {
		if element != "" {
			result = append(result, element)
		}
	}
	ptr = &result
	return len(result)
}
