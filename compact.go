package piscine

func Compact(ptr *[]string) int {
	var result []string
	var count int
	for _, element := range *ptr {
		if element != "" {
			result = append(result, element)
			count++
		}
	}
	*ptr = result[:count]
	return len(result)
}
