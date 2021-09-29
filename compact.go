package piscine

func Compact(ptr *[]string) int {
	var counter int
	for index, element := range *ptr {
		if len(element) == 0 {
			counter++
			*ptr = append((*ptr)[:index], (*ptr)[index+1:]...)
		}
	}
	return counter
}
