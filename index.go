package piscine

// Index : find index
func Index(s string, toFind string) int {
	for index, char := range s {
		if string(char) == toFind {
			return index
		}
	}
	return -1
}
