package piscine

func Join(elems []string, sep string) string {
	for _, str := range elems {
		sep += str
	}
	return sep
}
